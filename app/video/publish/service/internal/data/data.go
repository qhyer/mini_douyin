package data

import (
	"douyin/app/video/publish/service/internal/conf"
	rdb "douyin/common/cache/redis"
	"douyin/common/database/orm"
	minio1 "douyin/common/minio"
	"douyin/common/queue/kafka"
	"douyin/common/sync/fanout"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrm, NewRedis, NewMinio, NewKafka, NewVideoRepo)

// Data .
type Data struct {
	db       *gorm.DB
	redis    *redis.Client
	minio    *minio.Client
	kafka    sarama.SyncProducer
	cacheFan *fanout.Fanout
}

// NewData .
func NewData(c *conf.Data, orm *gorm.DB, redis *redis.Client, minio *minio.Client, kafka sarama.SyncProducer, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:       orm,
		redis:    redis,
		minio:    minio,
		kafka:    kafka,
		cacheFan: fanout.New(fanout.Worker(10), fanout.Buffer(10240)),
	}, cleanup, nil
}

func NewOrm(c *conf.Data) *gorm.DB {
	return orm.NewMySQL(&orm.Config{
		DSN:         c.GetOrm().GetDSN(),
		Active:      int(c.GetOrm().GetActive()),
		Idle:        int(c.GetOrm().GetIdle()),
		IdleTimeout: c.GetOrm().GetIdleTimeout().AsDuration(),
	})
}

func NewRedis(c *conf.Data) *redis.Client {
	return rdb.NewRedis(&rdb.Config{
		Name:         c.GetRedis().GetName(),
		Network:      c.GetRedis().GetNetwork(),
		Addr:         c.GetRedis().GetAddr(),
		Password:     c.GetRedis().GetPassword(),
		DialTimeout:  c.GetRedis().GetDialTimeout().AsDuration(),
		ReadTimeout:  c.GetRedis().GetReadTimeout().AsDuration(),
		WriteTimeout: c.GetRedis().GetWriteTimeout().AsDuration(),
	})
}

func NewMinio(c *conf.Data) *minio.Client {
	return minio1.NewMinio(&minio1.Config{
		EndPoint:        c.GetMinio().GetEndpoint(),
		AccessKeyID:     c.GetMinio().GetAccessKeyId(),
		SecretAccessKey: c.GetMinio().GetSecretAccessKey(),
	})
}

func NewKafka(c *conf.Data) sarama.SyncProducer {
	return kafka.NewKafkaSyncProducer(&kafka.Config{
		Addr: c.GetKafka().GetAddr(),
	})
}
