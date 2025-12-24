package dao

import (
	"errors"

	"github.com/wxq/metaland-blog/src/db"
	"github.com/wxq/metaland-blog/src/entity"
)

var UserDao = &userDao{}

type userDao struct{}

func (userDao *userDao) Create(user *entity.User) (err error) {
	res := db.GlobalDB.Create(user)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("create user failed")
	}
	return nil
}

func (userDao *userDao) FindUserByUserName(username string) (*entity.User, error) {
	var user *entity.User
	res := db.GlobalDB.Where("username = ?", username).Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}
