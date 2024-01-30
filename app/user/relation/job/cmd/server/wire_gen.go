// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"

	"douyin/app/user/relation/job/internal/biz"
	"douyin/app/user/relation/job/internal/conf"
	"douyin/app/user/relation/job/internal/data"
	"douyin/app/user/relation/job/internal/server"
	"douyin/app/user/relation/job/internal/service"

	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client := server.NewEtcdCli(registry)
	discovery := server.NewDiscovery(client)
	seqClient := data.NewSeqClient(discovery, logger)
	db := data.NewOrm(confData)
	redisClient := data.NewRedis(confData)
	consumer := data.NewKafkaConsumer(confData)
	syncProducer := data.NewKafkaProducer(confData)
	dataData, cleanup, err := data.NewData(confData, seqClient, db, redisClient, consumer, syncProducer, logger)
	if err != nil {
		return nil, nil, err
	}
	relationRepo := data.NewRelationRepo(dataData, logger)
	relationUsecase := biz.NewRelationUsecase(relationRepo, logger)
	relationService := service.NewRelationService(relationUsecase, consumer, logger)
	grpcServer := server.NewGRPCServer(confServer, relationService, logger)
	httpServer := server.NewHTTPServer(confServer, relationService, logger)
	registrar := server.NewRegistrar(client)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
