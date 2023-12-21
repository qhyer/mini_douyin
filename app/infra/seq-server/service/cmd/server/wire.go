//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"douyin/app/infra/seq-server/service/internal/biz"
	"douyin/app/infra/seq-server/service/internal/conf"
	"douyin/app/infra/seq-server/service/internal/data"
	"douyin/app/infra/seq-server/service/internal/server"
	"douyin/app/infra/seq-server/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, *conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
