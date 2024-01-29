package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

// Config redis config.
type Config struct {
	Name         string // redis name, for trace
	Network      string // tcp/udp ..
	Addr         string // address like localhost:6379
	Password     string // auth
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewRedis(c *Config) (rdb *redis.Client) {
	if c.DialTimeout <= 0 || c.ReadTimeout <= 0 || c.WriteTimeout <= 0 {
		panic("must config redis timeout")
	}
	rdb = redis.NewClient(&redis.Options{
		Network:      c.Network,
		Addr:         c.Addr,
		ClientName:   c.Name,
		Password:     c.Password,
		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
	})

	return
}
