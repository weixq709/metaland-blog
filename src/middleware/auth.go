package middleware

import (
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/response"
	"github.com/wxq/metaland-blog/src/utils/constant"
	"github.com/wxq/metaland-blog/src/utils/jwt"
)

var excludeLoginPaths = []string{"/user/login", "/swagger"}

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestPath := ctx.FullPath()

		// TODO 鉴权路径白名单
		for _, path := range excludeLoginPaths {
			if strings.HasPrefix(requestPath, path) {
				ctx.Next()
				return
			}
		}

		token := strings.TrimPrefix(ctx.GetHeader(constant.Authorization), constant.TokenPrefix)
		if token == "" {
			response.FailWithMessage(ctx, "无效token")
			ctx.Abort()
			return
		}

		_, err := jwt.Parse(token)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			ctx.Abort()
			return
		}

		session := sessions.Default(ctx)
		userName := session.Get(constant.UserNameKey)
		if userName == nil || userName == "" {
			response.FailWithMessage(ctx, "用户未登录")
			ctx.Abort()
			return
		}

		ctx.Set(constant.UserIdKey, session.Get(constant.UserIdKey))
		ctx.Set(constant.UserNameKey, userName)

		ctx.Next()
	}
}
