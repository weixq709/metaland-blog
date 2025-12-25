package response

import (
	"github.com/gin-gonic/gin"
)

const (
	fail    = 0
	success = 1
)

type Result struct {
	Message *string     `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func Success(ctx *gin.Context) {
	ctx.JSON(200, &Result{
		Code: success,
	})
}

func SuccessWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, &Result{Code: success, Data: data})
}

func SuccessWithMessage(ctx *gin.Context, message string) {
	ctx.JSON(200, &Result{Code: success, Message: &message})
}

func SuccessWithDetail(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(200, &Result{Code: success, Message: &message, Data: data})
}

func Fail(ctx *gin.Context) {
	message := "系统错误"
	ctx.JSON(200, &Result{Code: fail, Message: &message})
}

func FailWithMessage(ctx *gin.Context, message string) {
	ctx.JSON(200, &Result{Code: fail, Message: &message})
}

func FailWithDetail(ctx *gin.Context, code int, message string) {
	ctx.JSON(200, &Result{Code: code, Message: &message})
}
