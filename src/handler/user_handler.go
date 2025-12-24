package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wxq/metaland-blog/src/entity"
	"github.com/wxq/metaland-blog/src/response"
	"github.com/wxq/metaland-blog/src/service"
	"github.com/wxq/metaland-blog/src/utils/constant"
)

var userService = service.UserService

type UserRequestHandler struct{}

// Login 用户登录
//
//	@Summary		用户登录
//	@Tags			用户管理
//	@Description	用户登录
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entity.User		true	"登录用户信息"
//	@Success		200		{object}	response.Result	"登录成功"
//	@Router			/user/login [POST]
//	@Security		BearerAuth
func (handler *UserRequestHandler) Login(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBind(&user); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	token, err := userService.Login(ctx, user)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, token)
}

// Register 注册用户
//
//	@Summary		注册用户
//	@Tags			用户管理
//	@Description	注册一个新用户
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entity.User		true	"注册用户信息"
//	@Success		200		{object}	response.Result	"注册成功"
//	@Router			/user/register [POST]
func (handler *UserRequestHandler) Register(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBind(&user); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if err := userService.Register(user); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// FindCurrentUserInfo 查询当前登录用户信息
//
//	@Summary		查询当前登录用户信息
//	@Tags			用户管理
//	@Description	查询当前登录用户信息
//	@Produce		json
//	@Success		200	{object}	response.Result{data=entity.User}	"返回当前登录用户信息"
//	@Router			/user/findCurrentUserInfo [GET]
//	@Security		BearerAuth
func (handler *UserRequestHandler) FindCurrentUserInfo(ctx *gin.Context) {
	userName := ctx.GetString(constant.UserNameKey)
	user, err := service.UserService.FindUserByUserName(userName)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if user == nil {
		response.FailWithMessage(ctx, "用户不存在")
		return
	}
	user.Password = ""
	// 设置密码为空
	response.SuccessWithData(ctx, user)
}

// FindUserByUserName 根据用户名查询用户信息
//
//	@Summary		根据用户名查询用户信息
//	@Tags			用户管理
//	@Description	根据用户名查询用户信息
//	@Param			username	path	string	true	"用户名"
//	@Produce		json
//	@Success		200	{object}	response.Result{data=entity.User}	"用户信息"
//	@Router			/user/findUser/{username} [GET]
//	@Security		BearerAuth
func (handler *UserRequestHandler) FindUserByUserName(ctx *gin.Context) {
	userName := ctx.Param("username")
	user, err := service.UserService.FindUserByUserName(userName)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if user == nil {
		response.FailWithMessage(ctx, "用户不存在")
		return
	}
	response.SuccessWithData(ctx, user)
}
