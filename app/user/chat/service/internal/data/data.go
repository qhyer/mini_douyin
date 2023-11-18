package data

import (
	"douyin/app/user/chat/service/internal/conf"
	rdb "douyin/common/cache/redis"
	"douyin/common/queue/kafka"
	"github.com/IBM/sarama"
	"github.com/redis/go-redis/v9"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRedis, NewKafka, NewChatRepo)

// Data .
type Data struct {
	redis *redis.Client
	kafka sarama.SyncProducer
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
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
