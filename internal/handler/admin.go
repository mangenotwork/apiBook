package handler

import (
	"apiBook/common/ginHelper"
	"github.com/gin-gonic/gin"
)

func AdminCreateUser(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	ctx.APIOutPut("ok", "ok")
}

func AdminDeleteUser(c *gin.Context) {
}

func AdminDisableUser(c *gin.Context) {
}
