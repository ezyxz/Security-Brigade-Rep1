package router

import (
	"SearchEngineV2/controller"
	"SearchEngineV2/middleware"
	"github.com/gin-gonic/gin"
)

// InitRegisterRouter 用户验证登录注册路由
func InitUserRouter(Router *gin.RouterGroup) {

	userRouter := Router.Group("user")
	{
		userRouter.POST("/register", controller.Register)
		userRouter.GET("/isValid", controller.IsValid)
		userRouter.POST("/login", controller.Login)
		userRouter.GET("/whoAmI", middleware.AuthMiddleware, controller.WhoAmI)
	}
}
