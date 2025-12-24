package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/response"
	"github.com/wxq/metaland-blog/src/xzap/logger"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if cause := recover(); cause != nil {
				logger.Errorf("[Recovery] panic recovered, %v", cause)
				response.FailWithMessage(c, "系统错误")
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				logger.Errorf("[Recovery] %+v", err)
			}
			response.FailWithMessage(c, "系统错误")
			return
		}
	}
}
