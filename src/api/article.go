package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/handler"
)

type ArticleApi struct{}

func (api *ArticleApi) RegisterRoute(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/article")
	articleHandler := &handler.ArticleRequestHandler{}
	router.POST("/", articleHandler.Create)
	router.PUT("/", articleHandler.Update)
	router.DELETE("/:articleId", articleHandler.DeleteByID)
	router.GET("/:articleId", articleHandler.FindByID)
	router.GET("/FindByPage", articleHandler.FindByPage)
}
