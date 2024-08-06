package handler

import (
	"apiBook/common/ginHelper"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DocumentDirList(c *gin.Context) {
}

func DocumentDirCreate(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &ProjectAddUserReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}
	// todo...
}

func DocumentDirDelete(c *gin.Context) {
}

func DocumentDirModify(c *gin.Context) {
}

func DocumentDirSort(c *gin.Context) {
}

func DocumentList(c *gin.Context) {
}

func DocumentCreate(c *gin.Context) {
}

func DocumentModify(c *gin.Context) {
}

func DocumentDelete(c *gin.Context) {
}

func DocumentSort(c *gin.Context) {
}
