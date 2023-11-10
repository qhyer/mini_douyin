package memcache

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

type Server struct {
	Addr string
	Port int
}

func NewMemcached(srvs []Server) *memcache.Client {
	var configStrings []string
	for _, srv := range srvs {
		configStrings = append(configStrings, fmt.Sprintf("%s:%d", srv.Addr, srv.Port))
	}
	return memcache.New(configStrings...)
}
