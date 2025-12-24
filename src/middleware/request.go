package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/xzap/logger"
	"go.uber.org/zap"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request

		logger.Info(fmt.Sprintf("[RequestLogger]"),
			zap.String("method", req.Method),
			zap.String("url", req.URL.String()))

		c.Next()
	}
}
