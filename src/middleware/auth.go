package middleware

import (
	"runtime/debug"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/glob"
	"github.com/pkg/errors"
	"github.com/wxq/metaland-blog/src/config"
	"github.com/wxq/metaland-blog/src/response"
	"github.com/wxq/metaland-blog/src/utils/constant"
	xjwt "github.com/wxq/metaland-blog/src/utils/jwt"
	"github.com/wxq/metaland-blog/src/xzap/logger"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestPath := ctx.FullPath()

		for _, pattern := range config.SysConfig.ExcludeLoginPaths {
			matcher, err := glob.Compile(pattern)
			if err != nil {
				logger.Errorf("invalid path pattern: %s", pattern)
				continue
			}
			if matcher.Match(requestPath) {
				ctx.Next()
				return
			}
		}

		token := strings.TrimPrefix(ctx.GetHeader(constant.Authorization), constant.TokenPrefix)
		if token == "" {
			logger.Warn("无效token")
			response.FailWithMessage(ctx, "无效token")
			ctx.Abort()
			return
		}

		_, err := xjwt.Parse(token)
		if err != nil {
			var tokenExpiredErr *jwt.TokenExpiredError
			if errors.As(err, &tokenExpiredErr) {
				logger.Warn("token过期")
				response.FailWithMessage(ctx, "token过期")
				ctx.Abort()
				return
			}
			logger.Errorf("%s\n%s", err.Error(), string(debug.Stack()))
			response.FailWithMessage(ctx, err.Error())
			ctx.Abort()
			return
		}

		session := sessions.Default(ctx)
		userName := session.Get(constant.UserNameKey)
		if userName == nil || userName == "" {
			logger.Warnf("用户未登录")
			response.FailWithMessage(ctx, "用户未登录")
			ctx.Abort()
			return
		}

		ctx.Set(constant.UserIdKey, session.Get(constant.UserIdKey))
		ctx.Set(constant.UserNameKey, userName)

		ctx.Next()
	}
}
