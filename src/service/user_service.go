package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/wxq/metaland-blog/src/dao"
	"github.com/wxq/metaland-blog/src/entity"
	"github.com/wxq/metaland-blog/src/utils/constant"
	"github.com/wxq/metaland-blog/src/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

var UserService = &userService{}
var userDao = dao.UserDao

type userService struct{}

func (userService *userService) Login(ctx *gin.Context, loginUser entity.User) (*string, error) {
	if loginUser.UserName == "" {
		return nil, errors.New("用户名不能为空")
	}
	if loginUser.Password == "" {
		return nil, errors.New("密码不能为空")
	}
	user, err := userDao.FindUserByUserName(loginUser.UserName)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		return nil, errors.New("密码错误")
	}
	token, err := jwt.Generate(loginUser.UserName)
	if err != nil {
		return nil, err
	}
	session := sessions.Default(ctx)
	session.Set(constant.UserIdKey, user.ID)
	session.Set(constant.UserNameKey, user.UserName)
	if err := session.Save(); err != nil {
		return nil, err
	}
	return &token, nil
}

func (userService *userService) Register(user entity.User) error {
	if user.UserName == "" {
		return errors.New("用户名不能为空")
	}
	if user.Password == "" {
		return errors.New("密码不能为空")
	}
	// 根据用户名查找用户
	storeUser, err := userService.FindUserByUserName(user.UserName)
	if err != nil {
		return err
	}
	if storeUser != nil {
		return errors.New("用户已存在")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := userDao.Create(&user); err != nil {
		return errors.Wrap(err, "创建用户失败")
	}
	return nil
}

func (userService *userService) FindUserByUserName(userName string) (*entity.User, error) {
	if userName == "" {
		return nil, errors.New("用户名不能为空")
	}
	user, err := userDao.FindUserByUserName(userName)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
