package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/handler"
)

type CommentApi struct{}

func (api *CommentApi) RegisterRoute(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/comment")
	commentHandler := &handler.CommentRequestHandler{}
	router.POST("", commentHandler.Create)
	router.POST("/queryComments/:articleId", commentHandler.FindCommentsByArticleId)
}
