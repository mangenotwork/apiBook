package handler

import (
	"apiBook/common/fenci"
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
	"time"
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
		if err == nil {

			sort.Slice(dirDocList, func(i, j int) bool {
				if dirDocList[i].Sort < dirDocList[j].Sort {
					return true
				}
				return false
			})

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

	_, err = dao.NewProjectDao().Get(param.PId, userAcc, false)
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

	_, err = dao.NewProjectDao().Get(param.PId, userAcc, false)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(err, "你不是项目创建无删除权限")
		return
	}

	dirRecycleBin, err := dao.NewDirDao().Delete(param.PId, param.DirId)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	ctx.APIOutPut(dirRecycleBin, "删除目录成功")
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

	_, err = dao.NewProjectDao().Get(param.PId, userAcc, false)
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
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentDirSortReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	for i, v := range param.DirList {
		data, err := dao.NewDirDao().Get(param.PId, v)
		if err != nil {
			log.Error(err)
			continue
		}
		data.Sort = i
		err = dao.NewDirDao().Update(param.PId, v, data)
		if err != nil {
			log.Error(err)
			continue
		}
	}

	ctx.APIOutPut("完成", "完成")
	return
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

	_, err = dao.NewProjectDao().Get(param.PId, userAcc, false)
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

	_, err = dao.NewProjectDao().Get(param.ProjectId, userAcc, false)
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

	param.Content.UserAcc = userAcc
	param.Content.CreateTime = time.Now().Unix()

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

	if param.DocId == "" {
		ctx.APIOutPutError(fmt.Errorf("文档id不能为空"), "文档id不能为空")
		return
	}

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	_, err = dao.NewProjectDao().Get(param.PId, userAcc, false)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(err, err.Error())
		return
	}

	data, err := dao.NewDocDao().GetDocumentContent(param.PId, param.DocId)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(err, err.Error())
		return
	}

	baseInfo, err := dao.NewDocDao().GetDocument(param.PId, param.DocId)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(err, err.Error())
		return
	}

	data.Resp[0].RespTypeName = data.Resp[0].RespType.GetName()

	resp := &DocumentItemResp{
		Content: &DocumentContent{
			DocId:                     data.DocId,
			ProjectId:                 data.ProjectId,
			Name:                      data.Name,
			Url:                       data.Url,
			Method:                    data.Method,
			DescriptionHtml:           data.DescriptionHtml,
			ReqHeader:                 data.ReqHeader,
			ReqType:                   data.ReqType,
			ReqTypeName:               data.ReqType.GetName(),
			ReqBodyJson:               data.ReqBodyJson,
			ReqBodyText:               data.ReqBodyText,
			ReqBodyFormData:           data.ReqBodyFormData,
			ReqBodyXWWWFormUrlEncoded: data.ReqBodyXWWWFormUrlEncoded,
			ReqBodyXml:                data.ReqBodyXml,
			ReqBodyRaw:                data.ReqBodyRaw,
			ReqBodyBinary:             data.ReqBodyBinary,
			ReqBodyGraphQL:            data.ReqBodyGraphQL,
			ReqBodyInfo:               data.ReqBodyInfo,
			Resp:                      data.Resp,
			UserAcc:                   data.UserAcc,
			Date:                      utils.Timestamp2Date(data.CreateTime),
		},
		SnapshotList: make([]*SnapshotItem, 0),
		BaseInfo:     baseInfo,
	}

	snapshotList, err := dao.NewDocDao().GetDocumentSnapshotList(param.DocId)
	if err != nil {
		log.Error(err)
	}

	for _, v := range snapshotList {
		resp.SnapshotList = append(resp.SnapshotList, &SnapshotItem{
			SnapshotIdId:  v.SnapshotIdId,
			UserAcc:       v.UserAcc,
			Operation:     v.Operation,
			CreateTime:    v.CreateTime,
			CreateTimeStr: utils.Timestamp2Date(v.CreateTime),
		})
	}

	sort.Slice(resp.SnapshotList, func(i, j int) bool {
		return resp.SnapshotList[i].CreateTime > resp.SnapshotList[j].CreateTime
	})

	dataRaw := ""
	switch resp.Content.ReqType {
	case define.ReqTypeText:
		dataRaw = resp.Content.ReqBodyText
	//	ReqTypeFormData           = "form-data"
	//	ReqTypeXWWWFormUrlEncoded = "x-www-form-urlencoded"
	case define.ReqTypeJson:
		dataRaw = resp.Content.ReqBodyJson
	case define.ReqTypeXml:
		dataRaw = resp.Content.ReqBodyXml
	case define.ReqTypeRaw:
		dataRaw = resp.Content.ReqBodyRaw
		//	ReqTypeBinary = "binary"
		//	ReqTypeGraphQL = "GraphQL"
	}

	resp.ReqCode = GetAllReqCode(&ReqCodeArg{
		Method:      MethodType(data.Method),
		Url:         resp.Content.Url,
		ContentType: string(resp.Content.ReqType),
		Header:      data.GetReqHeaderMap(),
		DataRaw:     dataRaw,
	})

	ctx.APIOutPut(resp, "")
	return
}

func DocumentModify(c *gin.Context) {
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

	_, err = dao.NewProjectDao().Get(param.ProjectId, userAcc, false)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	err = dao.NewDocDao().Modify(param.Content, userAcc)
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

	_, err = dao.NewProjectDao().Get(param.PId, userAcc, false)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	err = dao.NewDocDao().Delete(param.PId, param.DocId)
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

	_, err = dao.NewProjectDao().Get(param.PId, userAcc, false)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	err = dao.NewDocDao().ChangeDir(param.PId, param.DirIdNew, param.DocId)
	if err != nil {
		ctx.APIOutPutError(err, "更改文档目录失败")
		return
	}

	ctx.APIOutPut("更改成功", "更改成功")
	return
}

func DocumentSort(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentSortReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	for i, v := range param.DocList {
		data, err := dao.NewDirDao().GetDoc(param.DirId, v)
		if err != nil {
			log.Error(err)
			continue
		}

		data.Sort = i
		err = dao.NewDirDao().UpdateDoc(param.DirId, v, data)
		if err != nil {
			log.Error(err)
			continue
		}
	}

	ctx.APIOutPut("完成", "完成")
	return
}

func DocumentDocList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentDocListReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	param.DocList = utils.SliceDeduplicate[string](param.DocList)

	resp := dao.NewDocDao().GetDocListByIds(param.PId, param.DocList)
	ctx.APIOutPut(resp, "")
	return
}

func DocumentSnapshotItem(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentSnapshotItemReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	data, err := dao.NewDocDao().GetDocumentSnapshotItem(param.DocId, param.SnapshotId)
	if err != nil {
		ctx.APIOutPutError(err, "获取镜像信息失败")
		return
	}

	resp := &DocumentSnapshotItemResp{
		Item:         data,
		SnapshotList: make([]*SnapshotItem, 0),
	}

	snapshotList, err := dao.NewDocDao().GetDocumentSnapshotList(param.DocId)
	if err != nil {
		log.Error(err)
	}

	for _, v := range snapshotList {
		resp.SnapshotList = append(resp.SnapshotList, &SnapshotItem{
			SnapshotIdId:  v.SnapshotIdId,
			UserAcc:       v.UserAcc,
			Operation:     v.Operation,
			CreateTime:    v.CreateTime,
			CreateTimeStr: utils.Timestamp2Date(v.CreateTime),
		})
	}

	sort.Slice(resp.SnapshotList, func(i, j int) bool {
		return resp.SnapshotList[i].CreateTime > resp.SnapshotList[j].CreateTime
	})

	dataRaw := ""
	switch resp.Item.DocumentContent.ReqType {
	case define.ReqTypeText:
		dataRaw = resp.Item.DocumentContent.ReqBodyText
	//	ReqTypeFormData           = "form-data"
	//	ReqTypeXWWWFormUrlEncoded = "x-www-form-urlencoded"
	case define.ReqTypeJson:
		dataRaw = resp.Item.DocumentContent.ReqBodyJson
	case define.ReqTypeXml:
		dataRaw = resp.Item.DocumentContent.ReqBodyXml
	case define.ReqTypeRaw:
		dataRaw = resp.Item.DocumentContent.ReqBodyRaw
		//	ReqTypeBinary = "binary"
		//	ReqTypeGraphQL = "GraphQL"
	}

	resp.ReqCode = GetAllReqCode(&ReqCodeArg{
		Method:      MethodType(resp.Item.DocumentContent.Method),
		Url:         resp.Item.DocumentContent.Url,
		ContentType: string(resp.Item.DocumentContent.ReqType),
		Header:      resp.Item.DocumentContent.GetReqHeaderMap(),
		DataRaw:     dataRaw,
	})

	ctx.APIOutPut(resp, "")
	return

}

func DocumentGetDirAll(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	pid := ctx.Query("pid")

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

	ctx.APIOutPut(data, "")
	return
}

func DocumentGetDocAll(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentGetDocAllReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	list := make([]*DocRespItem, 0)

	dirDocList, err := dao.NewDirDao().GetDocList(param.DirId)
	if err == nil {

		sort.Slice(dirDocList, func(i, j int) bool {
			if dirDocList[i].Sort < dirDocList[j].Sort {
				return true
			}
			return false
		})

		docList := dao.NewDocDao().GetDocList(param.PId, dirDocList)
		for _, docItem := range docList {
			list = append(list, &DocRespItem{
				DocId:  docItem.DocId,
				Method: docItem.Method,
				Title:  docItem.Name,
			})
		}
	}

	ctx.APIOutPut(list, "")

	return
}

func MoveToRecycleBin(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &MoveToRecycleBinReq{}
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

	_, err = dao.NewProjectDao().Get(param.PId, userAcc, false)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	allDir, err := dao.NewDirDao().GetDirList(param.PId)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(err, err.Error())
		return
	}

	recycleBinDir := ""

	for _, v := range allDir {
		if strings.Contains(v, "recycleBin_") {
			recycleBinDir = v
		}
	}

	log.Info("recycleBinDir = ", recycleBinDir)

	err = dao.NewDocDao().ChangeDir(param.PId, recycleBinDir, param.DocId)
	if err != nil {
		ctx.APIOutPutError(err, "更改文档目录失败")
		return
	}

	ctx.APIOutPut(recycleBinDir, "更改成功")
	return

}

func DocumentSearch(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &DocumentSearchReq{}
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

	_, err = dao.NewProjectDao().Get(param.PId, userAcc, false)
	if err != nil {
		ctx.APIOutPutError(err, err.Error())
		return
	}

	list := make([]*entity.InvertIndex, 0)

	strList := fenci.TermExtract(param.Word)
	for _, v := range strList {

		item, err := dao.NewInvertIndexDao().Get(param.PId, v.Text)
		if err != nil {
			log.Error(err)
			continue
		}

		list = append(list, item...)
	}

	docMap := make(map[string]*entity.InvertIndex)
	for _, v := range list {
		if docData, ok := docMap[v.DocId]; !ok {
			v.Score = 1
			docMap[v.DocId] = v
		} else {
			docData.Score += 1
		}
	}

	dirDocList := make([]*entity.DocumentDirItem, 0)
	i := 0
	for k, _ := range docMap {
		dirDocList = append(dirDocList, &entity.DocumentDirItem{
			DocId: k,
			Sort:  i,
		})
		i++
	}

	resp := &DocumentSearchResp{
		Count: len(docMap),
		List:  make([]*DocumentSearchRespItem, 0),
	}

	docList := dao.NewDocDao().GetDocList(param.PId, dirDocList)
	for _, docItem := range docList {
		item := &DocumentSearchRespItem{
			DocId:   docItem.DocId,
			Method:  docItem.Method,
			Title:   docItem.Name,
			Word:    docMap[docItem.DocId].Word,
			ModType: docMap[docItem.DocId].ModType,
			Score:   docMap[docItem.DocId].Score,
		}
		resp.List = append(resp.List, item)
	}

	sort.Slice(resp.List, func(i, j int) bool {
		if resp.List[i].Score > resp.List[j].Score {
			return true
		}
		return false
	})

	ctx.APIOutPut(resp, "更改成功")
	return
}
