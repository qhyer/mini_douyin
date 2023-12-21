package data

import (
	"context"
	account "douyin/api/user/account/service/v1"
	comment "douyin/api/video/comment/service/v1"
	favorite "douyin/api/video/favorite/service/v1"
	publish "douyin/api/video/publish/service/v1"
	"douyin/app/video/feed/service/internal/conf"
	rdb "douyin/common/cache/redis"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewFavoriteClient, NewPublishClient,
	NewAccountClient, NewCommentClient, NewRedis, NewMemcached, NewFeedRepo)

// Data .
type Data struct {
	publishCli  publish.PublishClient
	favoriteCli favorite.FavoriteClient
	accountCli  account.AccountClient
	commentCli  comment.CommentClient
	redis       *redis.Client
	memcached   *memcache.Client
}

// NewData .
func NewData(c *conf.Data, pc publish.PublishClient, fc favorite.FavoriteClient,
	ac account.AccountClient, cc comment.CommentClient, r *redis.Client, m *memcache.Client,
	logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		publishCli:  pc,
		favoriteCli: fc,
		accountCli:  ac,
		commentCli:  cc,
		redis:       r,
		memcached:   m,
	}, cleanup, nil
}

func NewPublishClient() publish.PublishClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return publish.NewPublishClient(conn)
}

func NewFavoriteClient() favorite.FavoriteClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return favorite.NewFavoriteClient(conn)
}

func NewAccountClient() account.AccountClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return account.NewAccountClient(conn)
}

func NewCommentClient() comment.CommentClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return comment.NewCommentClient(conn)
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
