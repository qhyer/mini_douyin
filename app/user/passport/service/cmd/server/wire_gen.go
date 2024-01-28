// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"douyin/app/user/passport/service/internal/biz"
	"douyin/app/user/passport/service/internal/conf"
	"douyin/app/user/passport/service/internal/data"
	"douyin/app/user/passport/service/internal/server"
	"douyin/app/user/passport/service/internal/service"
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
	seqClient := data.NewSeqClient(discovery, logger)
	db := data.NewOrm(confData)
	redisClient := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, seqClient, db, redisClient, logger)
	if err != nil {
		return nil, nil, err
	}
	passportRepo := data.NewPassportRepo(dataData, logger)
	passportUsecase := biz.NewPassportUseCase(passportRepo, logger)
	passportService := service.NewPassportService(passportUsecase)
	grpcServer := server.NewGRPCServer(confServer, passportService, logger)
	httpServer := server.NewHTTPServer(confServer, passportService, logger)
	registrar := server.NewRegistrar(client)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
