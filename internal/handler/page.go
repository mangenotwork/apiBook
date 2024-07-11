package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/common/conf"
	"github.com/mangenotwork/common/ginHelper"
	"net/http"
	"time"
)

func Testcase(c *ginHelper.GinCtx) {
	c.APIOutPut("test", "test")
}

func ginH(h gin.H) gin.H {
	h["Title"] = conf.Conf.Default.App.Name
	h["TimeStamp"] = time.Now().Unix() // define.TimeStamp
	return h
}

func Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		ginH(gin.H{
			"nav": "home",
		}),
	)
}
