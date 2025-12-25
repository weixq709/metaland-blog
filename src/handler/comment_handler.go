package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/entity"
	"github.com/wxq/metaland-blog/src/response"
	"github.com/wxq/metaland-blog/src/service"
	"github.com/wxq/metaland-blog/src/utils/param"
)

var commentService = service.CommentService{}

type CommentRequestHandler struct{}

// Create 新增评论
//
//	@Summary		新增评论
//	@Tags			评论管理
//	@Description	新增评论
//	@Param			comment	body	entity.Comment	true	"评论信息"
//	@Produce		json
//	@Success		200	{object}	response.Result	"新增成功"
//	@Router			/comment [POST]
//	@Security		BearerAuth
func (handler *CommentRequestHandler) Create(ctx *gin.Context) {
	var comment entity.Comment
	if err := ctx.ShouldBind(&comment); err != nil {
		_ = ctx.Error(err)
		return
	}
	if err := commentService.Create(ctx, comment); err != nil {
		_ = ctx.Error(err)
		return
	}
	response.Success(ctx)
}

// FindCommentsByArticleId 根据文章ID查询评论列表
//
//	@Summary		根据文章ID查询评论列表
//	@Tags			评论管理
//	@Description	根据文章ID查询评论列表
//	@Param			articleId	path	string	true	"文章ID"
//	@Produce		json
//	@Success		200	{object}	response.Result{data=[]entity.Comment}	"评论列表"
//	@Router			/comment/queryComments/{articleId} [GET]
//	@Security		BearerAuth
func (handler *CommentRequestHandler) FindCommentsByArticleId(ctx *gin.Context) {
	articleId := param.Path(ctx).Name("articleId").Value().GetInt64()
	comments, err := commentService.FindCommentsByArticleId(articleId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	response.SuccessWithData(ctx, comments)
}
