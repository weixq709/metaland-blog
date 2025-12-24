package router

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wxq/metaland-blog/docs"
	"github.com/wxq/metaland-blog/src/api"
)

func Start() {
	r := gin.New()

	r.Use(gin.Recovery())
	//store := cookie.NewStore([]byte("userInfo"))
	//// 配置 Session
	//store.Options(sessions.Options{
	//	Path:     "/",                  // Cookie 路径
	//	Domain:   "",                   // 域名，空表示当前域名
	//	MaxAge:   30 * 60,              // 过期时间（秒）：30分钟
	//	Secure:   false,                // true 表示仅 HTTPS
	//	HttpOnly: true,                 // 防止 JavaScript 访问
	//	SameSite: http.SameSiteLaxMode, // 防止 CSRF
	//})
	//r.Use(sessions.Sessions("userSession", store))
	//r.Use(middleware.Authentication())

	contextRouter := r.Group("/")
	contextRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.RegisterRouter(contextRouter)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	log.Printf("Listening and serving HTTP on %s\n", srv.Addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
