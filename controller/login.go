package controller

import (
	"SearchEngineV2/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.LoginRequest
	if err := c.ShouldBind(&u); err != nil {
		ResponseErrorWithMsg(c, "接口参数错误")
		return
	}
	isOK, token := srv.Login.LoginUser(&u)
	if isOK {
		ResponseSuccessWithData(c, token)
	} else {
		ResponseErrorWithMsg(c, token)
	}
	return
}

func WhoAmI(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		ResponseErrorWithMsg(c, "请先登录")
	} else {
		ResponseSuccessWithData(c, gin.H{"uid": uid})
	}
	return
}
