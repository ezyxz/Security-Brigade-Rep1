package router

import (
	"SearchEngineV2/controller"
	"github.com/gin-gonic/gin"
)

// InitRegisterRouter 用户验证登录注册路由
func InitRegisterRouter(Router *gin.RouterGroup) {

	registerRouter := Router.Group("user")
	{
		registerRouter.GET("/auth", controller.Authenticate)
		registerRouter.GET("/reg", controller.Register)
		registerRouter.GET("/del", controller.Delete)
	}
}
