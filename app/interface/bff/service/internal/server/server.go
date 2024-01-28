package server

import (
	"douyin/app/interface/bff/service/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewEtcdCli, NewRegistrar, NewDiscovery)

func NewEtcdCli(c *conf.Registry) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{c.GetEtcd().GetEndpoint()},
		DialTimeout: c.Etcd.DialTimeout.AsDuration(),
	})
	if err != nil {
		panic(err)
	}
	return cli
}

func NewRegistrar(cli *clientv3.Client) registry.Registrar {
	r := etcd.New(cli)
	return r
}

func NewDiscovery(cli *clientv3.Client) registry.Discovery {
	r := etcd.New(cli)
	return r
}
