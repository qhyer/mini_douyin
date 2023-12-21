package biz

import (
	"context"
	"douyin/app/infra/seq-server/service/common/entity"
	"douyin/app/infra/seq-server/service/internal/conf"
	"douyin/common/ecode"
	"douyin/common/network"
	os1 "douyin/common/os"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	clientv3 "go.etcd.io/etcd/client/v3"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	twepoch            = int64(1698768000000) // 起始时间戳
	workerIdBits       = 10
	maxWorkerId        = ^(-1 << workerIdBits)
	sequenceBits       = 12
	workerIdShift      = sequenceBits
	timestampLeftShift = sequenceBits + workerIdBits
	sequenceMask       = int64(^(-1 << sequenceBits))
	prefixEtcdPath     = "/leaf/"
	pathForever        = prefixEtcdPath + "/permanent/" // 保存所有数据持久的节点
	propPath           = "/prop/"
)

type SeqRepo interface {
	GetKeyWithPrefix(ctx context.Context, prefix string) (*clientv3.GetResponse, error)
	GetKey(ctx context.Context, key string) (*clientv3.GetResponse, error)
	SetKey(ctx context.Context, key, value string) bool
	SetKeyWithOptimizeLock(ctx context.Context, key, value string) bool
}

type SeqUsecase struct {
	repo     SeqRepo
	conf     *conf.Bootstrap
	registry *etcd.Registry
	log      *log.Helper

	mu                  sync.RWMutex
	lastTimestamp       int64
	sequence            int64
	workerId            int
	twepoch             int64
	snowflakeEtcdHolder entity.SnowFlakeEtcdHolder
}

func NewSeqUsecase(repo SeqRepo, etcdCli *clientv3.Client, conf *conf.Bootstrap, logger log.Logger) *SeqUsecase {
	s := &SeqUsecase{
		repo:     repo,
		registry: etcd.New(etcdCli),
		conf:     conf,
		log:      log.NewHelper(logger),
	}
	if !s.init() {
		panic("init seq usecase failed")
	}
	return s
}

func (u *SeqUsecase) GetID() (int64, error) {
	u.mu.Lock()
	defer u.mu.Unlock()
	ts := u.timeGen()
	if ts < u.lastTimestamp {
		offset := u.lastTimestamp - ts
		if offset <= 5 {
			time.Sleep(time.Duration(offset * 2))
			ts = u.timeGen()
			if ts < u.lastTimestamp {
				log.Errorf("clock is moving backwards. Rejecting requests until %d.", u.lastTimestamp)
				return 0, ecode.ServiceErr
			}
		} else {
			log.Errorf("clock is moving backwards. Rejecting requests until %d.", u.lastTimestamp)
			return 0, ecode.ServiceErr
		}
	}

	if ts == u.lastTimestamp {
		u.sequence = (u.sequence + 1) & sequenceMask
		if u.sequence == 0 {
			// seq 为0的时候表示是下一毫秒时间开始对seq做随机
			u.sequence = int64(rand.Intn(100))
			ts = u.tilNextMillis(ts)
		}
	} else {
		u.sequence = int64(rand.Intn(100))
	}
	u.lastTimestamp = ts
	return (ts-u.twepoch)<<int64(timestampLeftShift) | int64(u.workerId)<<int64(workerIdShift) | u.sequence, nil
}

func (u *SeqUsecase) tilNextMillis(lastTimestamp int64) int64 {
	var ts = u.timeGen()
	for ts <= lastTimestamp {
		ts = u.timeGen()
	}
	return ts
}

func (u *SeqUsecase) timeGen() int64 {
	return time.Now().UnixMilli()
}

func (u *SeqUsecase) init() bool {
	u.twepoch = twepoch
	if twepoch <= u.timeGen() {
		panic("twepoch must be > currentTimeMillis")
	}

	address, err := network.GetHostAddress("")
	if err != nil {
		panic(err)
	}

	u.snowflakeEtcdHolder.IP = address[0]
	u.snowflakeEtcdHolder.Port = strings.Split(u.conf.GetServer().GetHttp().GetAddr(), ":")[1]
	u.snowflakeEtcdHolder.EtcdAddressNode = fmt.Sprintf("%s/%s-0", pathForever, u.snowflakeEtcdHolder.ListenAddress)
	prefixEtcdPath = "/snowflake/" + u.conf.Server.GetServerName()
	propPath = filepath.Join(os1.GetCurrentAbPath(), u.conf.Server.GetServerName()) +
		"/leafconf/" + u.snowflakeEtcdHolder.Port + "/workerID.toml"
	pathForever = prefixEtcdPath + "/forever"
	u.log.Infof("workerID local cache file path : %s", propPath)

	u.snowflakeEtcdHolder.ListenAddress = u.snowflakeEtcdHolder.IP + ":" + u.snowflakeEtcdHolder.Port

	go u.cronUploadData()
	return u.initWorkerId()
}

func (u *SeqUsecase) initWorkerId() bool {
	prefixKeyResp, err := u.repo.GetKeyWithPrefix(context.Background(), pathForever)
	if err == nil {
		if prefixKeyResp.Count == 0 {
			// 未实例化
			u.snowflakeEtcdHolder.EtcdAddressNode = fmt.Sprintf("%s/%s-0", pathForever, u.snowflakeEtcdHolder.ListenAddress)
			ok := u.repo.SetKeyWithOptimizeLock(context.Background(), u.snowflakeEtcdHolder.EtcdAddressNode, u.encodeEndPoint())
			if !ok {
				// 已经有其他实例
				u.log.Infof("etcd node already exist snowflake %+v", u.snowflakeEtcdHolder.EtcdAddressNode)
				return false
			}
		} else {
			// 已经存在，寻找自己
			u.log.Infof("etcd node already exist snowflake %+v", u.snowflakeEtcdHolder.EtcdAddressNode)
			isSelfFound := false
			for _, kv := range prefixKeyResp.Kvs {
				nodeKey := strings.Split(filepath.Base(string(kv.Key)), "-")
				listenAddr := nodeKey[0]
				workerId := cast.ToInt(nodeKey[1])
				if listenAddr == u.snowflakeEtcdHolder.ListenAddress {
					if !u.checkInitTimeStamp(u.snowflakeEtcdHolder.EtcdAddressNode) {
						return false
					}
					u.workerId = workerId
					isSelfFound = true
					break
				}
			}
			if !isSelfFound {
				// 没找到自己，说明是新节点，创建一个，不用check时间
				workerId := 0
				// 找到当前最大的id
				for _, kv := range prefixKeyResp.Kvs {
					nodeKey := strings.Split(filepath.Base(string(kv.Key)), "-")
					wid := cast.ToInt(nodeKey[1])
					workerId = max(workerId, wid)
				}
				// 自增
				workerId++
				u.snowflakeEtcdHolder.WorkerId = workerId
				u.snowflakeEtcdHolder.EtcdAddressNode = fmt.Sprintf("%s/%s-%d", pathForever, u.snowflakeEtcdHolder.ListenAddress, workerId)
				ok := u.repo.SetKeyWithOptimizeLock(context.Background(), u.snowflakeEtcdHolder.EtcdAddressNode, u.encodeEndPoint())
				if !ok {
					// 已经有其他实例
					u.log.Infof("etcd node already exist snowflake %+v", u.snowflakeEtcdHolder.EtcdAddressNode)
					return false
				}
			}
		}
	} else {
		u.log.Errorf("connect to etcd failed, read workerID from local file")
		// 读不到etcd就从本地尝试读取
		if _, err := os.Stat(propPath); err == nil {
			readFile, err := ioutil.ReadFile(propPath)
			if err != nil {
				u.log.Errorf("%+v", err)
				return false
			}
			u.workerId = cast.ToInt(strings.Split(string(readFile), "=")[1])
		} else {
			u.log.Error("prop not exist")
			return false
		}
	}

	u.writeWorkerIdToProp(u.workerId)
	return true
}

func (u *SeqUsecase) writeWorkerIdToProp(workerId int) {
	if _, err := os.Stat(propPath); err != nil {
		err := os.MkdirAll(filepath.Dir(propPath), os.ModePerm)
		if err != nil {
			u.log.Errorf("%+v", err)
			return
		}
	}

	err := ioutil.WriteFile(propPath, []byte(fmt.Sprintf("workerID=%d", workerId)), os.ModePerm)
	if err != nil {
		u.log.Errorf("%+v", err)
		return
	}
	return
}

func (u *SeqUsecase) encodeEndPoint() string {
	end := entity.Endpoint{
		IP:        u.snowflakeEtcdHolder.IP,
		Port:      u.snowflakeEtcdHolder.Port,
		Timestamp: time.Now().UnixMilli(),
	}
	jsonString, err := end.EncodeToJSONString()
	if err != nil {
		u.log.Errorf("%+v", err)
		return ""
	}
	return jsonString
}

func (u *SeqUsecase) checkInitTimeStamp(addrNode string) bool {
	key, err := u.repo.GetKey(context.Background(), addrNode)
	if err != nil {
		return false
	}
	end := entity.Endpoint{}
	err = end.DecodeFromJSONString(string(key.Kvs[0].Value))
	if err != nil {
		return false
	}
	return !(end.Timestamp > u.timeGen())
}

func (u *SeqUsecase) cronUploadData() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			u.uploadData()
		}
	}
}

func (u *SeqUsecase) uploadData() {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.snowflakeEtcdHolder.LastUpdateTime = time.Now().UnixMilli()
	u.snowflakeEtcdHolder.WorkerId = u.workerId
	u.repo.SetKey(context.Background(), u.snowflakeEtcdHolder.EtcdAddressNode, u.encodeEndPoint())
}
