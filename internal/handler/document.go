package handler

import (
	"apiBook/common/ginHelper"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DocumentDirList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	pid := ctx.Query("pid")

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	data, err := dao.NewDirDao().GetAll(pid)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut(data, "")
	return
}

func DocumentDirCreate(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentDirCreateReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	_, err = dao.NewProjectDao().Get(param.PId, userAcc)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	dir := &entity.DocumentDir{
		DirId:   utils.IDStr(),
		DirName: param.DirName,
		Sort:    dao.NewDirDao().GetDirNum(param.PId) + 1,
	}

	err = dao.NewDirDao().Create(param.PId, dir)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut("创建目录成功", "创建目录成功")
	return
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
