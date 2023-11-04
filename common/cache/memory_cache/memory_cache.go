package memory_cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Config struct {
	DefaultExpiration time.Duration
	CleanUpInterval   time.Duration
}

func NewMemoryCache(c *Config) (mc *cache.Cache) {
	mc = cache.New(c.DefaultExpiration, c.CleanUpInterval)
	return
}
