// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"douyin/app/interface/bff/service/internal/biz"
	"douyin/app/interface/bff/service/internal/conf"
	"douyin/app/interface/bff/service/internal/data"
	"douyin/app/interface/bff/service/internal/server"
	"douyin/app/interface/bff/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client := server.NewEtcdCli(registry)
	discovery := server.NewDiscovery(client)
	publishClient := data.NewPublishClient(discovery, logger)
	accountClient := data.NewAccountClient(discovery, logger)
	passportClient := data.NewPassportClient(discovery, logger)
	feedClient := data.NewFeedClient(discovery, logger)
	favoriteClient := data.NewFavoriteClient(discovery, logger)
	commentClient := data.NewCommentClient(discovery, logger)
	relationClient := data.NewRelationClient(discovery, logger)
	dataData, cleanup, err := data.NewData(confData, publishClient, accountClient, passportClient, feedClient, favoriteClient, commentClient, relationClient, logger)
	if err != nil {
		return nil, nil, err
	}
	accountRepo := data.NewAccountRepo(dataData, logger)
	accountUsecase := biz.NewAccountUsecase(accountRepo, logger)
	feedRepo := data.NewFeedRepo(dataData, logger)
	feedUsecase := biz.NewFeedUsecase(feedRepo, logger)
	relationRepo := data.NewRelationRepo(dataData, logger)
	relationUsecase := biz.NewRelationUsecase(relationRepo, logger)
	commentRepo := data.NewCommentRepo(dataData, logger)
	commentUsecase := biz.NewCommentUsecase(commentRepo, logger)
	favoriteRepo := data.NewFavoriteRepo(dataData, logger)
	favoriteUsecase := biz.NewFavoriteUsecase(favoriteRepo, logger)
	publishRepo := data.NewPublishRepo(dataData, logger)
	publishUsecase := biz.NewPublishUsecase(publishRepo, logger)
	bffService := service.NewBFFService(accountUsecase, feedUsecase, relationUsecase, commentUsecase, favoriteUsecase, publishUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, bffService, logger)
	httpServer := server.NewHTTPServer(confServer, bffService, logger)
	registrar := server.NewRegistrar(client)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
