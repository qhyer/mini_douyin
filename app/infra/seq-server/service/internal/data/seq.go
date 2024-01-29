package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"

	"douyin/app/infra/seq-server/service/internal/biz"
)

type SeqRepo struct {
	data *Data
	log  *log.Helper
}

func NewSeqRepo(data *Data, logger log.Logger) biz.SeqRepo {
	return &SeqRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *SeqRepo) GetKeyWithPrefix(ctx context.Context, prefix string) (*clientv3.GetResponse, error) {
	return r.data.etcdCli.Get(ctx, prefix, clientv3.WithPrefix())
}

func (r *SeqRepo) GetKey(ctx context.Context, key string) (*clientv3.GetResponse, error) {
	return r.data.etcdCli.Get(ctx, key)
}

func (r *SeqRepo) SetKey(ctx context.Context, key, value string) bool {
	_, err := r.data.etcdCli.Put(ctx, key, value)
	if err != nil {
		return false
	}
	return true
}

func (r *SeqRepo) SetKeyWithOptimizeLock(ctx context.Context, key, value string) bool {
	txnResponse, err := r.data.etcdCli.Txn(ctx).
		If(clientv3.Compare(clientv3.CreateRevision(key), "=", 0)).
		Then(clientv3.OpPut(key, value)).Commit()
	if err != nil {
		return false
	}
	return txnResponse.Succeeded
}
