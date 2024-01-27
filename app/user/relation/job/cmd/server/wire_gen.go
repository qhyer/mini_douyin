// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"douyin/app/user/relation/job/internal/biz"
	"douyin/app/user/relation/job/internal/conf"
	"douyin/app/user/relation/job/internal/data"
	"douyin/app/user/relation/job/internal/server"
	"douyin/app/user/relation/job/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewOrm(confData)
	client := data.NewRedis(confData)
	consumer := data.NewKafkaConsumer(confData)
	syncProducer := data.NewKafkaProducer(confData)
	dataData, cleanup, err := data.NewData(confData, db, client, consumer, syncProducer, logger)
	if err != nil {
		return nil, nil, err
	}
	relationRepo := data.NewRelationRepo(dataData, logger)
	relationUsecase := biz.NewRelationUsecase(relationRepo, logger)
	relationService := service.NewRelationService(relationUsecase, consumer, logger)
	grpcServer := server.NewGRPCServer(confServer, relationService, logger)
	httpServer := server.NewHTTPServer(confServer, relationService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
