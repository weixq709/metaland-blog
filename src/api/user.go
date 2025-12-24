package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/handler"
)

type UserApi struct{}

func (api *UserApi) RegisterRoute(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	userHandler := &handler.UserRequestHandler{}
	userRouter.POST("/login", userHandler.Login)
	userRouter.POST("/register", userHandler.Register)
	userRouter.GET("/findCurrentUserInfo", userHandler.FindCurrentUserInfo)
	userRouter.GET("/findUser/:username", userHandler.FindUserByUserName)
}
