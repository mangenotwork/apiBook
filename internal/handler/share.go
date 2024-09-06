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

	if param.Expiration == 0 {
		param.Expiration = -1
	}

	err = dao.NewShareDao().Create(param)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(fmt.Errorf("创建分享失败"), "创建分享失败")
		return
	}

	ctx.APIOutPut("创建成功", "创建成功")
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

func DeleteShare(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	key := ctx.Query("key")

	info, err := dao.NewShareDao().Del(key)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(fmt.Errorf("删除分享失败"), "删除分享失败")
		return
	}

	ctx.APIOutPut(info, "删除成功")
	return
}
