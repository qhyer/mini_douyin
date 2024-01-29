package data

import (
	"github.com/IBM/sarama"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"douyin/app/video/favorite/service/internal/conf"
	rdb "douyin/common/cache/redis"
	"douyin/common/database/orm"
	"douyin/common/queue/kafka"
	"douyin/common/sync/fanout"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrm, NewRedis, NewKafka, NewFavoriteRepo)

// Data .
type Data struct {
	db       *gorm.DB
	redis    *redis.Client
	kafka    sarama.SyncProducer
	cacheFan *fanout.Fanout
}

// NewData .
func NewData(c *conf.Data, orm *gorm.DB, redis *redis.Client, kafka sarama.SyncProducer, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:       orm,
		redis:    redis,
		kafka:    kafka,
		cacheFan: fanout.New(fanout.Worker(10), fanout.Buffer(10240)),
	}, cleanup, nil
}

func NewOrm(c *conf.Data) *gorm.DB {
	return orm.NewMySQL(&orm.Config{
		DSN:         c.GetOrm().GetDsn(),
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

func NewKafka(c *conf.Data) sarama.SyncProducer {
	return kafka.NewKafkaSyncProducer(&kafka.Config{
		Addr: c.GetKafka().GetAddr(),
	})
}
