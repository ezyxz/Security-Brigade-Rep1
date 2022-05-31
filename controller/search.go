package controller

import (
	"github.com/gin-gonic/gin"
	"time"
)

func WordQuery(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		ResponseErrorWithMsg(c, "请输入关键字")
		return
	}
	start := time.Now()

	r := srv.Searcher.WordSearch(q)
	tc := float64(time.Since(start).Nanoseconds())
	r.Time = tc / 1e6
	ResponseSuccessWithData(c, r)
}
