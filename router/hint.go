package router

import (
	"SearchEngineV2/controller"
	"github.com/gin-gonic/gin"
)

func InitHintRouter(Router *gin.RouterGroup) {

	hintRouter := Router.Group("findEnd")
	{
		hintRouter.GET("/", controller.GetHint)
	}
}
