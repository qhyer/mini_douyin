//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"douyin/app/video/publish/service/internal/biz"
	"douyin/app/video/publish/service/internal/conf"
	"douyin/app/video/publish/service/internal/data"
	"douyin/app/video/publish/service/internal/server"
	"douyin/app/video/publish/service/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
