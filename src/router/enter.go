package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wxq/metaland-blog/docs"
	"github.com/wxq/metaland-blog/src/api"
	"github.com/wxq/metaland-blog/src/config"
	"github.com/wxq/metaland-blog/src/middleware"
	logger "github.com/wxq/metaland-blog/src/xzap/logger"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(middleware.Recover())
	store := cookie.NewStore([]byte("userInfo"))
	// 配置 Session
	store.Options(sessions.Options{
		Path:     "/",                  // Cookie 路径
		Domain:   "",                   // 域名，空表示当前域名
		MaxAge:   30 * 60,              // 过期时间（秒）：30分钟
		Secure:   false,                // true 表示仅 HTTPS
		HttpOnly: true,                 // 防止 JavaScript 访问
		SameSite: http.SameSiteLaxMode, // 防止 CSRF
	})
	r.Use(sessions.Sessions("userSession", store))
	r.Use(middleware.RequestLogger())
	r.Use(middleware.Authentication())

	contextRouter := r.Group(config.SysConfig.ContextPath)
	contextRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.RegisterRouter(contextRouter)

	for _, route := range r.Routes() {
		logger.Infof("%-6s %-40s --> %s", route.Method, route.Path, route.Handler)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.SysConfig.Port),
		Handler: r.Handler(),
	}

	logger.Infof("Listening and serving HTTP on: %d (http) with context path '%s'", config.SysConfig.Port, config.SysConfig.ContextPath)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("Error starting HTTP server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("Server Shutdown Failed:%+v", err)
	}
	logger.Infof("Server exiting")
}
