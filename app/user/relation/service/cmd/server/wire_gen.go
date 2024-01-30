// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"

	"douyin/app/user/relation/service/internal/biz"
	"douyin/app/user/relation/service/internal/conf"
	"douyin/app/user/relation/service/internal/data"
	"douyin/app/user/relation/service/internal/server"
	"douyin/app/user/relation/service/internal/service"

	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewOrm(confData)
	client := data.NewRedis(confData)
	syncProducer := data.NewKafka(confData)
	dataData, cleanup, err := data.NewData(confData, db, client, syncProducer, logger)
	if err != nil {
		return nil, nil, err
	}
	relationRepo := data.NewRelationRepo(dataData, logger)
	relationUsecase := biz.NewRelationUsecase(relationRepo, logger)
	relationService := service.NewRelationService(relationUsecase)
	grpcServer := server.NewGRPCServer(confServer, relationService, logger)
	httpServer := server.NewHTTPServer(confServer, logger)
	clientv3Client := server.NewEtcdCli(registry)
	registrar := server.NewRegistrar(clientv3Client)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
