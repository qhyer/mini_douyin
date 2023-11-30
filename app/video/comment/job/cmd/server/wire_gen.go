// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"douyin/app/video/comment/job/internal/biz"
	"douyin/app/video/comment/job/internal/conf"
	"douyin/app/video/comment/job/internal/data"
	"douyin/app/video/comment/job/internal/server"
	"douyin/app/video/comment/job/internal/service"
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
	dataData, cleanup, err := data.NewData(confData, db, client, logger)
	if err != nil {
		return nil, nil, err
	}
	commentRepo := data.NewCommentRepo(dataData, logger)
	commentUsecase := biz.NewCommentUsecase(commentRepo, logger)
	consumer := service.NewKafka(confData)
	commentService := service.NewCommentService(commentUsecase, consumer, logger)
	grpcServer := server.NewGRPCServer(confServer, commentService, logger)
	httpServer := server.NewHTTPServer(confServer, commentService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
