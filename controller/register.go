package controller

import (
	"SearchEngineV2/register"
	"fmt"
	"github.com/gin-gonic/gin"
)

type QueryResult struct {
	uid int32 `json:"uid"`
}

func Authenticate(c *gin.Context) {
	u := c.Query("u")
	p := c.Query("p")
	if u == "" || p == "" {
		ResponseErrorWithMsg(c, "请输入用户名和密码")
		return
	}
	uid := srv.Register.CheckUserExisted(&register.LogUser{
		UserName: u,
		UserPwd:  p,
	})
	if uid == -1 {
		ResponseErrorWithMsg(c, "用户名或密码错误")
	} else {
		fmt.Println(uid)
		ResponseSuccessWithData(c, &QueryResult{uid: uid})
	}

}

func Register(c *gin.Context) {
	u := c.Query("u")
	p := c.Query("p")
	if u == "" || p == "" {
		ResponseErrorWithMsg(c, "请输入用户名和密码")
		return
	}
	res := srv.Register.RegisterNewUser(&register.User{
		UserName: u,
		UserPwd:  p,
	})
	if res > 0 {
		ResponseSuccessWithData(c, "注册重成功")
	} else {
		ResponseErrorWithMsg(c, "重复用户名")
	}
}

func Delete(c *gin.Context) {
	ResponseErrorWithMsg(c, "功能开发中")
}
