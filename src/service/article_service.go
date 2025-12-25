package service

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/dao"
	"github.com/wxq/metaland-blog/src/entity"
	"github.com/wxq/metaland-blog/src/entity/page"
	"github.com/wxq/metaland-blog/src/utils/constant"
)

var ArticleService = &articleService{}

var articleDao = dao.ArticleDao

type articleService struct{}

func (articleService *articleService) Create(ctx *gin.Context, article entity.Article) error {
	if article.Title == "" {
		return errors.New("标题不能为空")
	}
	if article.Content == "" {
		return errors.New("内容不能为空")
	}
	userId := sessions.Default(ctx).Get(constant.UserIdKey).(int64)
	article.UserID = userId
	return articleDao.Create(&article)
}

func (articleService *articleService) Update(ctx *gin.Context, article entity.Article) error {
	if article.ID == 0 {
		return errors.New("ID不能为空")
	}
	if article.Title == "" {
		return errors.New("标题不能为空")
	}
	if article.Content == "" {
		return errors.New("内容不能为空")
	}
	userId := sessions.Default(ctx).Get(constant.UserIdKey).(int64)
	article.UserID = userId
	return articleDao.Update(article)
}

func (articleService *articleService) DeleteByID(ctx *gin.Context, id int64) error {
	if id == 0 {
		return errors.New("ID不能为空")
	}
	userId := sessions.Default(ctx).Get(constant.UserIdKey).(int64)
	return articleDao.DeleteByID(entity.Article{ID: id, UserID: userId})
}

func (articleService *articleService) FindById(id int64) (*entity.Article, error) {
	if id == 0 {
		return nil, errors.New("ID不能为空")
	}
	article, err := articleDao.FindById(id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (articleService *articleService) FindByPage(queryPage page.QueryPage, article entity.Article) ([]entity.Article, error) {
	return articleDao.FindByPage(queryPage, article)
}
