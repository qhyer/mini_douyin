package server

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"

	"douyin/app/infra/seq-server/service/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar, NewDiscovery)

func NewRegistrar(c *conf.Data) registry.Registrar {
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

func NewDiscovery(c *conf.Data) registry.Discovery {
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
