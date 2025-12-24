package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/dao"
	"github.com/wxq/metaland-blog/src/entity"
	"github.com/wxq/metaland-blog/src/utils/constant"
)

var commentService = &CommentService{}
var commentDao = dao.CommentDao

type CommentService struct{}

func (service *CommentService) Create(ctx *gin.Context, comment entity.Comment) error {
	userId := sessions.Default(ctx).Get(constant.UserIdKey).(int64)
	comment.UserId = userId
	return commentDao.Create(comment)
}

func (service *CommentService) FindCommentsByArticleId(articleId int64) ([]entity.Comment, error) {
	comments, err := commentDao.FindCommentsByArticleId(articleId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
