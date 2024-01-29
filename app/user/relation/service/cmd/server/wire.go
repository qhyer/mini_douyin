//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"douyin/app/user/relation/service/internal/biz"
	"douyin/app/user/relation/service/internal/conf"
	"douyin/app/user/relation/service/internal/data"
	"douyin/app/user/relation/service/internal/server"
	"douyin/app/user/relation/service/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
