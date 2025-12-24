package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/entity/page"
	"github.com/wxq/metaland-blog/src/handler"
	"github.com/wxq/metaland-blog/src/response"
)

type UserApi struct{}

func (api *UserApi) RegisterRoute(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	userHandler := &handler.UserRequestHandler{}
	userRouter.POST("/login", userHandler.Login)
	userRouter.POST("/register", userHandler.Register)
	userRouter.GET("/findCurrentUserInfo", userHandler.FindCurrentUserInfo)
	userRouter.GET("/findUser/:username", userHandler.FindUserByUserName)
	userRouter.GET("/test", func(context *gin.Context) {
		queryPage := page.Defaults()
		if err := context.ShouldBindQuery(&queryPage); err != nil {
			response.FailWithMessage(context, err.Error())
			return
		}
		fmt.Println(queryPage)
		response.SuccessWithData(context, "test")
	})
}
