package handler

import (
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ShareCreate(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &entity.Share{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	param.Key = utils.NewShortCode()

	err = dao.NewShareDao().Create(param)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(fmt.Errorf("创建分享失败"), "创建分享失败")
		return
	}

	ctx.APIOutPut(param.Key, "")
	return
}

func GetShareInfoProject(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	pid := ctx.Query("pid")
	data, err := dao.NewShareDao().GetShareProjectList(pid)
	if err != nil {
		log.Error(err)
	}
	ctx.APIOutPut(data, "")
	return
}

func GetShareInfoDocument(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	docId := ctx.Query("docId")
	data, err := dao.NewShareDao().GetShareDocumentList(docId)
	if err != nil {
		log.Error(err)
	}
	ctx.APIOutPut(data, "")
	return
}
