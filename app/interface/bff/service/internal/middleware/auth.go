package middleware

import (
	"context"

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
					if hr.Request().Method == "POST" && hr.Request().Header.Get("Content-Type") == "application/json" {
						tokenStr = hr.Request().Header.Get("token")
					} else {
						tokenStr = hr.Request().URL.Query().Get("token")
					}

					// 解析JWT
					id, err := jwt.ParseTokenToID(tokenStr)
					if isRouteRequireLogin(hr.Request().URL.Path) && err != nil {
						return nil, ecode.AuthorizeErr
					}
					ctx = context.WithValue(ctx, "uid", id)
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
