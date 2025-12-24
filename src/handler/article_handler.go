package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/entity"
	"github.com/wxq/metaland-blog/src/entity/page"
	"github.com/wxq/metaland-blog/src/response"
	"github.com/wxq/metaland-blog/src/service"
)

var articleService = service.ArticleService

type ArticleRequestHandler struct{}

// Create 创建文章
//
//	@Summary		创建文章
//	@Tags			文章管理
//	@Description	创建文章
//	@Accept			json
//	@Produce		json
//	@Param			article	body		entity.Article	true	"文章信息"
//	@Success		200		{object}	response.Result	"创建成功"
//	@Router			/article [POST]
//	@Security		BearerAuth
func (handler *ArticleRequestHandler) Create(ctx *gin.Context) {
	var article entity.Article
	if err := ctx.ShouldBind(&article); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if err := articleService.Create(ctx, article); err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
}

// Update 根据ID修改文章信息
//
//	@Summary		根据ID修改文章信息
//	@Tags			文章管理
//	@Description	根据ID修改文章信息，支持修改标题和内容
//	@Accept			json
//	@Produce		json
//	@Param			article	body		entity.Article	true	"文章信息"
//	@Success		200		{object}	response.Result	"修改成功"
//	@Router			/article [PUT]
//	@Security		BearerAuth
func (handler *ArticleRequestHandler) Update(ctx *gin.Context) {
	var article entity.Article
	if err := ctx.ShouldBind(&article); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if err := articleService.Update(ctx, article); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// DeleteByID 根据文章ID删除文章信息
//
//	@Summary		根据文章ID删除文章信息
//	@Tags			文章管理
//	@Description	根据文章ID删除文章信息
//	@Accept			json
//	@Produce		json
//	@Param			articleId	path		int				true	"文章ID"
//	@Success		200			{object}	response.Result	"删除成功"
//	@Router			/article/{articleId} [DELETE]
//	@Security		BearerAuth
func (handler *ArticleRequestHandler) DeleteByID(ctx *gin.Context) {
	articleId := ctx.GetInt64("articleId")
	err := articleService.DeleteByID(ctx, articleId)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// FindByID 根据文章ID查询文章详情
//
//	@Summary		根据文章ID查询文章详情
//	@Tags			文章管理
//	@Description	根据文章ID查询文章详情
//	@Produce		json
//	@Param			articleId	path		int										true	"文章ID"
//	@Success		200			{object}	response.Result{data=entity.Article}	"文章信息"
//	@Router			/article/{articleId} [GET]
//	@Security		BearerAuth
func (handler *ArticleRequestHandler) FindByID(ctx *gin.Context) {
	articleId := ctx.GetInt64("articleId")
	articles, err := articleService.FindById(articleId)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, articles)
}

// FindByPage 分页查询
//
//	@Summary		分页查询
//	@Tags			文章管理
//	@Description	根据条件分页查询，支持查询指定用户文章列表
//	@Produce		json
//	@Success		200	{object}	response.Result{data=[]entity.Article}	"文章列表"
//	@Router			/article/findAll [GET]
//	@Security		BearerAuth
func (handler *ArticleRequestHandler) FindByPage(ctx *gin.Context) {
	var article entity.Article
	queryPage := page.Defaults()
	if err := ctx.ShouldBindQuery(&queryPage); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if err := ctx.ShouldBind(&article); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	articles, err := articleService.FindByPage(queryPage, article)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, articles)
}
