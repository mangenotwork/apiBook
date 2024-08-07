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
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentDirDeleteReq{}

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	_, err := dao.NewProjectDao().Get(param.PId, userAcc)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	err = dao.NewDirDao().Delete(param.PId, param.DirId)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	// 目录重新排序

	ctx.APIOutPut("删除目录成功", "删除目录成功")
	return
}

func DocumentDirModify(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentDirModifyReq{}
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

	err = dao.NewDirDao().Modify(param.PId, param.DirId, param.DirName)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut("修改目录成功", "修改目录成功")
	return
}

func DocumentDirSort(c *gin.Context) {
	// todo 方案未定
}

func DocumentList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentListReq{}
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

	dirDocList, err := dao.NewDirDao().GetDocList(param.DirId)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	resp := dao.NewDocDao().GetDocList(param.PId, dirDocList)

	ctx.APIOutPut(resp, "")
	return
}

func DocumentCreate(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := entity.DocumentParam{}
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

	_, err = dao.NewProjectDao().Get(param.ProjectId, userAcc)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	docId := utils.IDMd5()

	doc := &entity.Document{
		DocId:     docId,
		DirId:     param.DirId,
		ProjectId: param.ProjectId,
		Name:      param.Content.Name,
		Url:       param.Content.Url,
		Method:    param.Content.Method,
	}

	err = dao.NewDocDao().Create(doc, param.Content)
	if err != nil {
		ctx.APIOutPutError(err, "接口文档创建失败")
		return
	}

	dirItem := &entity.DocumentDirItem{
		DocId: docId,
		Sort:  0,
	}

	err = dao.NewDirDao().AddDoc(param.DirId, dirItem)
	if err != nil {
		ctx.APIOutPutError(err, "接口文档创建失败")
		return
	}

	ctx.APIOutPut("创建文档成功", "创建文档成功")
	return
}

func DocumentItem(c *gin.Context) {

}

func DocumentModify(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := entity.DocumentParam{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}
}

func DocumentDelete(c *gin.Context) {
}

func DocumentSort(c *gin.Context) {
	// todo...
}
