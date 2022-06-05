package controller

import (
	"SearchEngineV2/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var u model.RegisterUserRequest
	if err := c.ShouldBind(&u); err != nil {
		ResponseErrorWithMsg(c, "接口参数错误")
		return
	}
	isOK := srv.Register.RegisterUser(&u)
	if isOK {
		ResponseSuccess(c)
	} else {
		ResponseErrorWithMsg(c, "注册用户失败")
	}
	return
}

func IsValid(c *gin.Context) {

}
