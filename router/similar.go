package router

import (
	"SearchEngineV2/controller"

	"github.com/gin-gonic/gin"
)

// InitSimilar 相关搜索路由
func InitSimilar(Router *gin.RouterGroup) {

	wordRouter := Router.Group("similar")
	{
		wordRouter.GET("similar", controller.Similar)
	}
}
