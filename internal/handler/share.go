package handler

import (
	"apiBook/common/conf"
	"apiBook/common/fenci"
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
)

func ShareCreate(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &entity.Share{}
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

	log.SendOperationLog(userAcc, fmt.Sprintf("创建分享: %s", utils.AnyToJsonNotErr(param)))

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

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	info, err := dao.NewShareDao().Del(key)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(fmt.Errorf("删除分享失败"), "删除分享失败")
		return
	}

	log.SendOperationLog(userAcc, fmt.Sprintf("删除分享: %s", key))

	ctx.APIOutPut(info, "删除成功")
	return
}

func ShareDocumentDirList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	hashKey := ctx.Query("hashKey")

	shareInfo, err := dao.NewShareDao().GetInfo(hashKey)
	if err != nil {
		log.Error(err)
		ginHelper.AuthErrorOut(c)
		return
	}

	if shareInfo.ShareType != 1 {
		ginHelper.AuthErrorOut(c)
		return
	}

	pid := shareInfo.ShareId

	resp := make([]*DocumentDirListItem, 0)

	resp = getDirCache(pid)
	if len(resp) > 0 {
		log.Info("读取的缓存")
		ctx.APIOutPut(resp, "")
		return
	}

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
			docList := dao.NewDocDao().GetDocList(pid, dirDocList)

			for _, docItem := range docList {
				item.Doc = append(item.Doc, &DocRespItem{
					DocId:  docItem.DocId,
					Method: strings.ToUpper(docItem.Method),
					Title:  docItem.Name,
				})
			}

		}

		resp = append(resp, item)
	}

	setDirCache(pid, resp)

	ctx.APIOutPut(resp, "")
	return
}

func ShareDocumentDocList(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	hashKey := ctx.Query("hashKey")

	shareInfo, err := dao.NewShareDao().GetInfo(hashKey)
	if err != nil {
		log.Error(err)
		ginHelper.AuthErrorOut(c)
		return
	}

	if shareInfo.ShareType != 1 {
		ginHelper.AuthErrorOut(c)
		return
	}

	param := &DocumentDocListReq{}
	err = ctx.GetPostArgs(&param)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	param.DocList = utils.SliceDeduplicate[string](param.DocList)

	for i, v := range param.DocList {
		if v == "" {
			param.DocList = append(param.DocList[:i], param.DocList[i+1:]...)
		}
	}

	resp := dao.NewDocDao().GetDocListByIds(param.PId, param.DocList)
	ctx.APIOutPut(resp, "")
	return
}

func ShareDocumentItem(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	hashKey := ctx.Query("hashKey")

	_, err := dao.NewShareDao().GetInfo(hashKey)
	if err != nil {
		log.Error(err)
		ginHelper.AuthErrorOut(c)
		return
	}

	param := &DocumentItemParam{}
	err = ctx.GetPostArgs(&param)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	_, err = dao.NewProjectDao().Get(param.PId, "", true)
	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(err, err.Error())
		return
	}

	resp, has := getDocCache(param.PId, param.DocId)
	if has {
		ctx.APIOutPut(resp, "")
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

	resp = &DocumentItemResp{
		Content: &DocumentContent{
			DocId:                     data.DocId,
			ProjectId:                 data.ProjectId,
			Name:                      data.Name,
			Url:                       data.Url,
			Method:                    strings.ToUpper(data.Method),
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

	setDocCache(param.PId, param.DocId, resp)

	ctx.APIOutPut(resp, "")
	return
}

func ShareProjectCodeGet(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	hashKey := ctx.Query("hashKey")

	shareInfo, err := dao.NewShareDao().GetInfo(hashKey)
	if err != nil {
		log.Error(err)
		ginHelper.AuthErrorOut(c)
		return
	}

	pid := shareInfo.ProjectId

	codeList, err := dao.NewProjectDao().GetGlobalCode(pid)
	if err != nil {
		log.Error(err)
		codeList = make([]*entity.GlobalCodeItem, 0)
	}

	ctx.APIOutPut(codeList, "")
	return
}

func ShareProjectHeaderGet(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	hashKey := ctx.Query("hashKey")

	shareInfo, err := dao.NewShareDao().GetInfo(hashKey)
	if err != nil {
		log.Error(err)
		ginHelper.AuthErrorOut(c)
		return
	}

	pid := shareInfo.ProjectId

	headerList, err := dao.NewProjectDao().GetGlobalHeader(pid)
	if err != nil {
		log.Error(err)
		headerList = make([]*entity.ReqHeaderItem, 0)
	}

	ctx.APIOutPut(headerList, "")
	return
}

func ShareDocumentSnapshotItem(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	hashKey := ctx.Query("hashKey")

	_, err := dao.NewShareDao().GetInfo(hashKey)
	if err != nil {
		log.Error(err)
		ginHelper.AuthErrorOut(c)
		return
	}

	param := &DocumentSnapshotItemReq{}
	err = ctx.GetPostArgs(&param)
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

func ShareVerify(ctx *gin.Context) {
	hashKey := ctx.Param("hashKey")
	password := ctx.PostForm("password")

	shareInfo, err := dao.NewShareDao().GetInfo(hashKey)
	if err != nil {
		ctx.HTML(200, "err.html", gin.H{
			"Title":     conf.Conf.Default.App.Name,
			"err":       "未知页面",
			"returnUrl": "/",
		})
		return
	}

	if shareInfo.PasswordCode != password {
		ctx.HTML(200, "err.html", gin.H{
			"Title":     conf.Conf.Default.App.Name,
			"err":       "阅读码错误",
			"returnUrl": "/browse/" + shareInfo.Key,
		})
		return
	}

	browseSignKey := utils.IDMd5()
	ctx.SetCookie("browseKey_"+shareInfo.Key, browseSignKey, 60*60*24*30, "/", "", false, true)

	browseSign := utils.GetMD5Encode(shareInfo.Key + shareInfo.PasswordCode + browseSignKey)
	ctx.SetCookie("browseSign_"+shareInfo.Key, browseSign, 60*60*24*30, "/", "", false, true)

	ctxHelper := ginHelper.NewGinCtx(ctx)
	ip := ""
	if val, ok := ctxHelper.Get(ginHelper.ReqIP); ok {
		ip = utils.AnyToString(val)
	}
	log.SendOperationLog(ip, fmt.Sprintf("分享验证: %s", hashKey))

	ctx.Redirect(http.StatusFound, "/browse/"+shareInfo.Key)
	return
}

func ShareDocumentSearch(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	hashKey := ctx.Query("hashKey")

	_, err := dao.NewShareDao().GetInfo(hashKey)
	if err != nil {
		log.Error(err)
		ginHelper.AuthErrorOut(c)
		return
	}

	param := &DocumentSearchReq{}
	err = ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
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
