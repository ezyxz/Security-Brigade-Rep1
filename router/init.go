package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	group := router.Group("/api")
	{
		InitWordRouter(group)   // 分词管理
		InitSearchRouter(group) //搜索
		InitUserRouter(group)
	}

	return router

}
