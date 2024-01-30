package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	"douyin/common/ecode"
	"douyin/common/jwt"
)

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if hr, ok := tr.(*http.Transport); ok {
					var tokenStr string
					tokenStr = hr.Request().PostFormValue("token")
					if tokenStr == "" {
						tokenStr = hr.Request().URL.Query().Get("token")
					}

					log.Debugf("tokenStr: %s", tokenStr)
					// 解析JWT
					id, err := jwt.ParseTokenToID(tokenStr)
					if isRouteRequireLogin(hr.Request().URL.Path) && err != nil {
						log.Debugf("ParseTokenToID err: %v", err)
						return nil, ecode.AuthorizeErr
					}
					ctx = context.WithValue(ctx, "userId", id)
				}
			}
			return handler(ctx, req)
		}
	}
}

var RequireLogin = []string{
	"/douyin/publish/action/",
	"/douyin/favorite/action/",
	"/douyin/comment/action/",
	"/douyin/relation/action/",
	"/relation/friend/list/",
	"/message/action/",
	"/message/chat/",

	"/douyin/publish/action",
	"/douyin/favorite/action",
	"/douyin/comment/action",
	"/douyin/relation/action",
	"/relation/friend/list",
	"/message/action",
	"/message/chat",
}

func isRouteRequireLogin(path string) bool {
	for _, v := range RequireLogin {
		if v == path {
			return true
		}
	}
	return false
}
