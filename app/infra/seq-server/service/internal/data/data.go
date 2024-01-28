package data

import (
	"douyin/app/infra/seq-server/service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEtcdCli, NewSeqRepo)

// Data .
type Data struct {
	etcdCli *clientv3.Client
}

// NewData .
func NewData(c *conf.Data, etcdCli *clientv3.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{etcdCli: etcdCli}, cleanup, nil
}

func NewEtcdCli(c *conf.Data, logger log.Logger) *clientv3.Client {
	l := log.NewHelper(log.With(logger))
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{c.GetEtcd().GetEndpoint()},
		DialTimeout: c.Etcd.DialTimeout.AsDuration(),
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		l.Error(err)
		panic(err)
	}
	return cli
}
