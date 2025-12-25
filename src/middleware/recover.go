package middleware

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/response"
	"github.com/wxq/metaland-blog/src/xzap/logger"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if cause := recover(); cause != nil {
				logger.Errorf("[Recovery] panic recovered, err: %s\n%s", cause, string(debug.Stack()))
				response.FailWithMessage(c, "系统错误")
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			logger.Errorf("%s\n%s", err.Error(), string(debug.Stack()))
			response.FailWithMessage(c, err.Error())
		}
	}
}
