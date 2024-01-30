package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	account "douyin/api/user/account/service/v1"
	passport "douyin/api/user/passport/service/v1"
	relation "douyin/api/user/relation/service/v1"
	comment "douyin/api/video/comment/service/v1"
	favorite "douyin/api/video/favorite/service/v1"
	feed "douyin/api/video/feed/service/v1"
	publish "douyin/api/video/publish/service/v1"
	"douyin/app/interface/bff/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPublishClient, NewAccountClient, NewPassportClient,
	NewFeedClient, NewCommentClient, NewFavoriteClient, NewRelationClient, NewPublishRepo,
	NewAccountRepo, NewCommentRepo, NewFavoriteRepo, NewFeedRepo, NewRelationRepo)

// Data .
type Data struct {
	PublishRPC  publish.PublishClient
	AccountRPC  account.AccountClient
	PassportRPC passport.PassportClient
	VideoRPC    feed.FeedClient
	CommentRPC  comment.CommentClient
	FavoriteRPC favorite.FavoriteClient
	RelationRPC relation.RelationClient
}

// NewData .
func NewData(c *conf.Data, pc publish.PublishClient, ac account.AccountClient, pasc passport.PassportClient,
	vc feed.FeedClient, favc favorite.FavoriteClient, cc comment.CommentClient, rc relation.RelationClient,
	logger log.Logger,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		PublishRPC:  pc,
		AccountRPC:  ac,
		PassportRPC: pasc,
		VideoRPC:    vc,
		CommentRPC:  cc,
		FavoriteRPC: favc,
		RelationRPC: rc,
	}, cleanup, nil
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

func NewAccountClient(r registry.Discovery, logger log.Logger) account.AccountClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///douyin.account.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}
	return account.NewAccountClient(conn)
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

func NewFeedClient(r registry.Discovery, logger log.Logger) feed.FeedClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///douyin.feed.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}
	return feed.NewFeedClient(conn)
}

func NewCommentClient(r registry.Discovery, logger log.Logger) comment.CommentClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///douyin.comment.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}
	return comment.NewCommentClient(conn)
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
