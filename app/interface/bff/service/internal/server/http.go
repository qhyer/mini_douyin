package server

import (
	"github.com/go-kratos/kratos/v2/middleware/validate"

	v1 "douyin/api/bff"
	"douyin/app/interface/bff/service/internal/conf"
	"douyin/app/interface/bff/service/internal/middleware"
	"douyin/app/interface/bff/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, bff *service.BFFService, logger log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.ErrorEncoder(middleware.ErrorEncoder),
		http.RequestDecoder(PublishActionDecoder),
		http.Middleware(
			validate.Validator(),
			recovery.Recovery(),
			middleware.Auth(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterBFFHTTPServer(srv, bff)
	return srv
}
