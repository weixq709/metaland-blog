package dao

import (
	"github.com/pkg/errors"
	"github.com/wxq/metaland-blog/src/db"
	"github.com/wxq/metaland-blog/src/entity"
)

var CommentDao = &commentDao{}

type commentDao struct{}

func (d *commentDao) Create(comment entity.Comment) error {
	res := db.GlobalDB.Create(&comment)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("create user failed")
	}
	return res.Error
}

func (d *commentDao) FindCommentsByArticleId(articleId int64) ([]entity.Comment, error) {
	var comments []entity.Comment
	res := db.GlobalDB.Where("article_id = ?", articleId).Find(&comments)
	if res.Error != nil {
		return nil, res.Error
	}
	return comments, nil
}
