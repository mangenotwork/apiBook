package handler

import (
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
)

func DocumentDirList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	pid := ctx.Query("pid")

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	resp := make([]*DocumentDirListItem, 0)

	data, err := dao.NewDirDao().GetAll(pid)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	sort.Slice(data, func(i, j int) bool {
		if data[i].Sort < data[j].Sort {
			return true
		}
		return false
	})

	for _, v := range data {
		item := &DocumentDirListItem{
			Dir: &DirRespItem{
				DirId: v.DirId,
				Name:  v.DirName,
			},
			Doc: make([]*DocRespItem, 0),
		}
		dirDocList, err := dao.NewDirDao().GetDocList(v.DirId)
		if err != nil {
			log.Error(err)
		} else {
			docList := dao.NewDocDao().GetDocList(pid, dirDocList)
			for _, docItem := range docList {
				item.Doc = append(item.Doc, &DocRespItem{
					DocId:  docItem.DocId,
					Method: docItem.Method,
					Title:  docItem.Name,
				})
			}
		}
		resp = append(resp, item)
	}

	ctx.APIOutPut(resp, "")
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
	param := &entity.DocumentParam{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		log.Error(err)
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
		log.Error("获取项目权限失败， err: ", err)
		ctx.APIOutPutError(err, err.Error())
		return
	}

	param.Content.DocId = utils.IDMd5()

	doc := &entity.Document{
		DocId:     param.Content.DocId,
		DirId:     param.DirId,
		ProjectId: param.ProjectId,
		Name:      param.Content.Name,
		Url:       param.Content.Url,
		Method:    param.Content.Method,
	}

	err = dao.NewDocDao().Create(doc, param.Content)
	if err != nil {
		log.Error("接口文档创建失败， err: ", err)
		ctx.APIOutPutError(err, "接口文档创建失败")
		return
	}

	dirItem := &entity.DocumentDirItem{
		DocId: param.Content.DocId,
		Sort:  0,
	}

	err = dao.NewDirDao().AddDoc(param.DirId, dirItem)
	if err != nil {
		log.Error("接口文档加入目录失败， err: ", err)
		ctx.APIOutPutError(err, "接口文档创建失败")
		return
	}

	ctx.APIOutPut("创建文档成功", "创建文档成功")
	return
}

func DocumentItem(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentItemParam{}
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

	data, err := dao.NewDocDao().GetDocumentContent(param.PId, param.DocId)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	resp := &DocumentItemResp{
		Content:      data,
		SnapshotList: make([]*SnapshotItem, 0),
	}

	snapshotList, err := dao.NewDocDao().GetDocumentSnapshotList(param.DocId)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	for _, v := range snapshotList {
		resp.SnapshotList = append(resp.SnapshotList, &SnapshotItem{
			SnapshotIdId: v.SnapshotIdId,
			UserAcc:      v.UserAcc,
			Operation:    v.Operation,
			CreateTime:   v.CreateTime,
		})
	}

	ctx.APIOutPut(resp, "")
	return
}

func DocumentModify(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &entity.DocumentParam{}
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

	err = dao.NewDocDao().Modify(param.Content)
	if err != nil {
		ctx.APIOutPutError(err, "修改文档失败")
		return
	}

	ctx.APIOutPut("修改文档成功", "修改文档成功")
	return
}

func DocumentDelete(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentDeleteReq{}
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

	err = dao.NewDocDao().Delete(param.PId, param.DirId, param.DocId)
	if err != nil {
		ctx.APIOutPutError(err, "删除文档失败")
		return
	}

	ctx.APIOutPut("删除文档成功", "删除文档成功")
	return
}

func DocumentChangeDir(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentChangeDirReq{}
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

	err = dao.NewDocDao().ChangeDir(param.PId, param.DirId, param.DirIdNew, param.DocId)
	if err != nil {
		ctx.APIOutPutError(err, "更改文档目录失败")
		return
	}

	ctx.APIOutPut("更改成功", "更改成功")
	return
}

func DocumentSort(c *gin.Context) {
	// todo...
}

func DocumentDocList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentDocListReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	resp := dao.NewDocDao().GetDocListByIds(param.PId, param.DocList)
	ctx.APIOutPut(resp, "")
	return
}
