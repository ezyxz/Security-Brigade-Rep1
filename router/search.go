package router

import (
	"SearchEngineV2/controller"
	"github.com/gin-gonic/gin"
)

func InitSearchRouter(Router *gin.RouterGroup) {

	searchRouter := Router.Group("query")
	{
		searchRouter.GET("/", controller.WordQuery)
	}
}
