package middleware

import (
	"SearchEngineV2/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	//获取Authorization，header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusOK, &model.ResponseData{
			State:   false,
			Message: "请先登录",
			Data:    nil,
		})
		//抛弃请求
		c.Abort()
		return
	}
	claims, err := ParserToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, &model.ResponseData{
			State:   false,
			Message: err.Error(),
			Data:    nil,
		})
		c.Abort()
		return
	}
	userID := claims.ID
	c.Set("uid", int32(userID))
	c.Next()
}
