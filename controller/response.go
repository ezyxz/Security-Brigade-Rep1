package controller

import (
	"SearchEngineV2/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseSuccessWithData 携带数据成功返回
func ResponseSuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &model.ResponseData{
		State:   true,
		Message: "success",
		Data:    data,
	})
}

// ResponseErrorWithMsg 返回错误
func ResponseErrorWithMsg(c *gin.Context, message string) {
	c.JSON(http.StatusOK, &model.ResponseData{
		State:   false,
		Message: message,
		Data:    nil,
	})
}

// ResponseSuccess 返回成功
func ResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, &model.ResponseData{
		State:   true,
		Message: "success",
		Data:    nil,
	})
}
