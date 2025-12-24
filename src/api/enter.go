package api

import "github.com/gin-gonic/gin"

type RouteRegistrar interface {
	RegisterRoute(r *gin.RouterGroup)
}

var apis = []RouteRegistrar{
	new(UserApi),
}

func RegisterRouter(router *gin.RouterGroup) {
	for _, api := range apis {
		api.RegisterRoute(router)
	}
}
