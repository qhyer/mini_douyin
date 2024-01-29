package data

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	seq "douyin/api/seq-server/service/v1"
	"douyin/app/video/comment/job/internal/conf"
	rdb "douyin/common/cache/redis"
	"douyin/common/database/orm"
	"douyin/common/queue/kafka"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrm, NewRedis, NewKafkaConsumer, NewKafkaProducer, NewCommentRepo, NewSeqClient)

// Data .
type Data struct {
	db            *gorm.DB
	redis         *redis.Client
	kafkaConsumer sarama.Consumer
	kafkaProducer sarama.SyncProducer
	seqRPC        seq.SeqClient
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, rds *redis.Client, kafkaConsumer sarama.Consumer,
	kafkaProducer sarama.SyncProducer, logger log.Logger, s seq.SeqClient,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:            db,
		redis:         rds,
		seqRPC:        s,
		kafkaProducer: kafkaProducer,
		kafkaConsumer: kafkaConsumer,
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

func NewSeqClient(r registry.Discovery, logger log.Logger) seq.SeqClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///douyin.seq.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}
	return seq.NewSeqClient(conn)
}

func NewKafkaProducer(c *conf.Data) sarama.SyncProducer {
	return kafka.NewKafkaSyncProducer(&kafka.Config{
		Addr: c.GetKafka().GetAddr(),
	})
}

func NewKafkaConsumer(c *conf.Data) sarama.Consumer {
	return kafka.NewKafkaConsumer(&kafka.Config{
		Addr: c.GetKafka().GetAddr(),
	})
}
