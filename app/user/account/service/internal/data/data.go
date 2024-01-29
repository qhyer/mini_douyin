package data

import (
	"context"

	"github.com/bluele/gcache"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"

	passport "douyin/api/user/passport/service/v1"
	relation "douyin/api/user/relation/service/v1"
	favorite "douyin/api/video/favorite/service/v1"
	publish "douyin/api/video/publish/service/v1"
	"douyin/app/user/account/service/internal/conf"
	rdb "douyin/common/cache/redis"
	"douyin/common/sync/fanout"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPassportClient, NewRelationClient, NewFavoriteClient,
	NewPublishClient, NewRedis, NewMemcached, NewAccountRepo)

// Data .
type Data struct {
	relationCli relation.RelationClient
	passportCli passport.PassportClient
	favoriteCli favorite.FavoriteClient
	publishCli  publish.PublishClient
	redis       *redis.Client
	memcached   *memcache.Client
	cacheFan    *fanout.Fanout
	localCache  gcache.Cache
}

// NewData .
func NewData(c *conf.Data, rc relation.RelationClient, pc passport.PassportClient, fc favorite.FavoriteClient,
	pubc publish.PublishClient, rds *redis.Client, mem *memcache.Client, logger log.Logger,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		relationCli: rc,
		passportCli: pc,
		favoriteCli: fc,
		publishCli:  pubc,
		redis:       rds,
		memcached:   mem,
		cacheFan:    fanout.New(fanout.Worker(10), fanout.Buffer(10240)),
		localCache:  gcache.New(10240).LFU().Build(),
	}, cleanup, nil
}

func NewPassportClient(r registry.Discovery, logger log.Logger) passport.PassportClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///douyin.passport.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}
	return passport.NewPassportClient(conn)
}

func NewRelationClient(r registry.Discovery, logger log.Logger) relation.RelationClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///douyin.relation.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}
	return relation.NewRelationClient(conn)
}

func NewFavoriteClient(r registry.Discovery, logger log.Logger) favorite.FavoriteClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///douyin.favorite.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}
	return favorite.NewFavoriteClient(conn)
}

func NewPublishClient(r registry.Discovery, logger log.Logger) publish.PublishClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///douyin.publish.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}
	return publish.NewPublishClient(conn)
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

func NewMemcached(c *conf.Data) *memcache.Client {
	return memcache.New(c.GetMemcached().GetAddr())
}
