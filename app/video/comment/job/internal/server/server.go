package server

import (
	"douyin/app/video/comment/job/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegister)

func NewRegister(c *conf.Register) registry.Registrar {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{c.GetEtcd().GetEndpoint()},
		DialTimeout: c.Etcd.DialTimeout.AsDuration(),
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(cli)
	return r
}
