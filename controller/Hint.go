package controller

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Suffix struct {
	Time  float64  `json:"time(ms),omitempty"` //查询用时
	Total int      `json:"total"`              //总数
	Hints []string `json:"hints,omitempty"`    //搜索关键词
}

func GetHint(c *gin.Context) {
	q := c.Query("h")
	if q == "" {
		ResponseErrorWithMsg(c, "输入关键字H")
		return
	}
	start := time.Now()

	hints := srv.Hint.WordFindEnd(q)
	tc := float64(time.Since(start).Nanoseconds())

	ResponseSuccessWithData(c, &Suffix{
		Total: len(hints),
		Time:  tc / 1e6,
		Hints: hints,
	})
}
