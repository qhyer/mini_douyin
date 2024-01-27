// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"douyin/app/video/publish/service/internal/biz"
	"douyin/app/video/publish/service/internal/conf"
	"douyin/app/video/publish/service/internal/data"
	"douyin/app/video/publish/service/internal/server"
	"douyin/app/video/publish/service/internal/service"
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
	minioClient := data.NewMinio(confData)
	syncProducer := data.NewKafka(confData)
	seqClient := data.NewSeqClient()
	dataData, cleanup, err := data.NewData(confData, db, client, minioClient, syncProducer, seqClient, logger)
	if err != nil {
		return nil, nil, err
	}
	videoRepo := data.NewVideoRepo(dataData, logger)
	videoUsecase := biz.NewVideoUsecase(videoRepo, logger)
	publishService := service.NewPublishService(videoUsecase)
	grpcServer := server.NewGRPCServer(confServer, publishService, logger)
	httpServer := server.NewHTTPServer(confServer, publishService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
