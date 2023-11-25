package data

import (
	"context"
	passport "douyin/api/user/passport/service/v1"
	relation "douyin/api/user/relation/service/v1"
	favorite "douyin/api/video/favorite/service/v1"
	publish "douyin/api/video/publish/service/v1"
	"douyin/app/user/account/service/internal/conf"
	rdb "douyin/common/cache/redis"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPassportClient, NewRelationClient, NewRedis, NewAccountRepo)

// Data .
type Data struct {
	relationCli relation.RelationClient
	passportCli passport.PassportClient
	favoriteCli favorite.FavoriteClient
	publishCli  publish.PublishClient
	redis       *redis.Client
}

// NewData .
func NewData(c *conf.Data, rc relation.RelationClient, pc passport.PassportClient, fc favorite.FavoriteClient, pubc publish.PublishClient, rds *redis.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		relationCli: rc,
		passportCli: pc,
		favoriteCli: fc,
		publishCli:  pubc,
		redis:       rds,
	}, cleanup, nil
}

func NewPassportClient() passport.PassportClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return passport.NewPassportClient(conn)
}

func NewRelationClient() relation.RelationClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return relation.NewRelationClient(conn)
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
