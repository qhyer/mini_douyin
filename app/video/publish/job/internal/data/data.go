package data

import (
	"douyin/app/video/publish/job/internal/conf"
	rdb "douyin/common/cache/redis"
	"douyin/common/database/orm"
	minio1 "douyin/common/minio"
	"github.com/IBM/sarama"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrm, NewRedis, NewMinio, NewVideoRepo)

// Data .
type Data struct {
	db    *gorm.DB
	redis *redis.Client
	minio *minio.Client
}

// NewData .
func NewData(c *conf.Data, orm *gorm.DB, redis *redis.Client, minio *minio.Client, kafka sarama.Consumer, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: orm, redis: redis, minio: minio}, cleanup, nil
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
