package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	"douyin/common/constants"
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
						if tokenStr == "" {
							// 游客
							ctx = context.WithValue(ctx, "uid", constants.GuestUserID)
							return handler(ctx, req)
						}
					}

					// 解析JWT
					id, err := jwt.ParseTokenToID(tokenStr)
					if err != nil {
						return nil, ecode.AuthorizeErr
					}
					ctx = context.WithValue(ctx, "uid", id)
				}
			}
			return handler(ctx, req)
		}
	}
}
