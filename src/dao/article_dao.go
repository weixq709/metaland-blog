package dao

import (
	"github.com/pkg/errors"
	"github.com/wxq/metaland-blog/src/db"
	"github.com/wxq/metaland-blog/src/entity"
	"github.com/wxq/metaland-blog/src/entity/page"
)

var ArticleDao = &articleDao{}

type articleDao struct{}

func (articleDao *articleDao) Create(article *entity.Article) error {
	res := db.GlobalDB.Create(&article)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("create failed")
	}
	return nil
}

func (articleDao *articleDao) Update(article entity.Article) error {
	res := db.GlobalDB.
		Model(&article).
		Where("user_id = ?", article.UserID).
		Select("Title", "Content", "UpdateTime").
		Updates(article)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (articleDao *articleDao) DeleteByID(article entity.Article) error {
	res := db.GlobalDB.
		Where("id = ?", article.ID).
		Where("user_id = ?", article.UserID).
		Delete(&entity.Article{})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (articleDao *articleDao) FindById(id int64) (*entity.Article, error) {
	var article entity.Article
	res := db.GlobalDB.Where("id = ?", id).First(&article)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &article, nil
}

func (articleDao *articleDao) FindByPage(queryPage page.QueryPage, article entity.Article) ([]entity.Article, error) {
	var articles []entity.Article
	tx := db.GlobalDB.Where("1=1")
	if article.UserID != 0 {
		tx = tx.Where("user_id = ?", article.UserID)
	}
	res := tx.Offset((queryPage.PageNum - 1) * queryPage.PageSize).
		Limit(queryPage.PageNum * queryPage.PageSize).
		Find(&articles)
	if res.Error != nil {
		return nil, res.Error
	}
	return articles, nil
}
