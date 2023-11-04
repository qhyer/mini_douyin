package data

import (
	"douyin/app/user/passport/service/internal/conf"
	rdb "douyin/common/cache/redis"
	"douyin/common/database/orm"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPassportRepo)

// Data .
type Data struct {
	db    *gorm.DB
	redis *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: orm.NewMySQL(&orm.Config{
		DSN:         c.GetOrm().GetDSN(),
		Active:      int(c.GetOrm().GetActive()),
		Idle:        int(c.GetOrm().GetIdle()),
		IdleTimeout: c.GetOrm().GetIdleTimeout().AsDuration(),
	}), redis: rdb.NewRedis(&rdb.Config{
		Name:         c.GetRedis().GetName(),
		Network:      c.GetRedis().GetNetwork(),
		Addr:         c.GetRedis().GetAddr(),
		Password:     c.GetRedis().GetPassword(),
		DialTimeout:  c.GetRedis().GetDialTimeout().AsDuration(),
		ReadTimeout:  c.GetRedis().GetReadTimeout().AsDuration(),
		WriteTimeout: c.GetRedis().GetWriteTimeout().AsDuration(),
	})}, cleanup, nil
}
