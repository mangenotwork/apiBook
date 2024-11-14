package docIE

import (
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"encoding/json"
	"strings"
	"time"
)

/*

apizza 导入,仅支持json

https://www.apizza.net/

*/

type ApiZZAImport struct {
	Data map[string]interface{}
}

func NewApiZZAImport() *ApiZZAImport {
	return &ApiZZAImport{
		Data: make(map[string]interface{}),
	}
}

func (obj *ApiZZAImport) Whole(text, userAcc string, private define.ProjectPrivateCode) error {

	err := obj.analysis(text)
	if err != nil {
		log.Error(err)
		return err
	}

	project := obj.analysisProject(userAcc, private)

	if dao.NewProjectDao().HasName(project.Name) {
		project.Name += utils.NowDateNotLine()
	}

	err = dao.NewProjectDao().Create(project, userAcc)
	if err != nil {
		log.Error("创建项目失败")
		return err
	}

	err = dao.NewDirDao().CreateInit(project.ProjectId)
	if err != nil {
		log.Error("创建项目失败")
		return err
	}

	obj.analysisDoc(project, "user", "",
		func(project *entity.Project, dirName string) (string, bool) {
			dir := &entity.DocumentDir{
				DirId:   utils.IDStr(),
				DirName: dirName,
				Sort:    dao.NewDirDao().GetDirNum(project.ProjectId) + 1,
			}

			err = dao.NewDirDao().Create(project.ProjectId, dir)
			if err != nil {
				log.Error("创建项目失败")
				return "", false
			}
			return dir.DirId, true
		},
		func(project *entity.Project, doc *entity.DocumentContent, dirId string) {

			documentData := &entity.Document{
				DocId:     doc.DocId,
				DirId:     dirId,
				ProjectId: doc.ProjectId,
				Name:      doc.Name,
				Url:       doc.Url,
				Method:    doc.Method,
			}

			err = dao.NewDocDao().Create(documentData, doc)
			if err != nil {
				log.Error("接口文档创建失败， err: ", err)
				return
			}

			dirItem := &entity.DocumentDirItem{
				DocId: doc.DocId,
				Sort:  0,
			}

			err = dao.NewDirDao().AddDoc(dirId, dirItem)
			if err != nil {
				log.Error("接口文档加入目录失败， err: ", err)
				return
			}
			log.Info("导入成功")

		})
	return nil
}

func (obj *ApiZZAImport) Increment(text, pid, userAcc, dirId string) error {
	project, err := dao.NewProjectDao().Get(pid, userAcc, false)
	if err != nil {
		log.Error("获取项目失败, err = ", err)
		return err
	}

	err = obj.analysis(text)
	if err != nil {
		return err
	}

	obj.analysisDoc(project, "user", dirId,
		func(project *entity.Project, dirName string) (string, bool) {
			return "", false
		},
		func(project *entity.Project, doc *entity.DocumentContent, dirId string) {

			documentData := &entity.Document{
				DocId:     doc.DocId,
				DirId:     dirId,
				ProjectId: doc.ProjectId,
				Name:      doc.Name,
				Url:       doc.Url,
				Method:    doc.Method,
			}

			err = dao.NewDocDao().Create(documentData, doc)
			if err != nil {
				log.Error("接口文档创建失败， err: ", err)
				return
			}

			dirItem := &entity.DocumentDirItem{
				DocId: doc.DocId,
				Sort:  0,
			}

			err = dao.NewDirDao().AddDoc(dirId, dirItem)
			if err != nil {
				log.Error("接口文档加入目录失败， err: ", err)
				return
			}
			log.Info("导入成功")

		})

	return nil
}

func (obj *ApiZZAImport) analysis(text string) error {
	err := json.Unmarshal([]byte(text), &obj.Data)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (obj *ApiZZAImport) analysisProject(userAcc string, private define.ProjectPrivateCode) *entity.Project {
	projectInfo := utils.AnyToMap(obj.Data["project_info"])
	name := utils.AnyToString(projectInfo["name"])
	comment := utils.AnyToString(projectInfo["comment"])
	createTime := utils.AnyToString(projectInfo["create_time"])
	return &entity.Project{
		ProjectId:     utils.IDMd5(),
		Name:          name,
		Description:   comment,
		CreateUserAcc: userAcc,
		CreateDate:    createTime,
		Private:       private,
	}
}

func (obj *ApiZZAImport) analysisDoc(project *entity.Project, userAcc, dirId string,
	createDir func(project *entity.Project, dirName string) (string, bool),
	f func(project *entity.Project, doc *entity.DocumentContent, dirId string)) {
	categorys := utils.AnyToArr(obj.Data["categorys"])
	for _, category := range categorys {

		categoryMap := utils.AnyToMap(category)
		dirName := utils.AnyToString(categoryMap["name"])
		apiList := utils.AnyToArr(categoryMap["api_list"])

		for _, v := range apiList {
			vMap := utils.AnyToMap(v)
			obj.analysisApi(project, dirName, userAcc, dirId, vMap, createDir, f)
		}

		subCategorys := utils.AnyToArr(categoryMap["sub_categorys"])
		if len(subCategorys) > 0 {
			for _, sub := range subCategorys {
				subMap := utils.AnyToMap(sub)
				subDirName := dirName + "-" + utils.AnyToString(subMap["name"])
				subApiList := utils.AnyToArr(subMap["api_list"])

				for _, v := range subApiList {
					vMap := utils.AnyToMap(v)
					obj.analysisApi(project, subDirName, userAcc, dirId, vMap, createDir, f)
				}
			}
		}

	}

}

func (obj *ApiZZAImport) analysisApi(project *entity.Project, dirName, userAcc, dirId string, apiData map[string]interface{},
	createDir func(project *entity.Project, dirName string) (string, bool),
	f func(project *entity.Project, doc *entity.DocumentContent, dirId string)) {

	log.Info("apiData = ", apiData)
	now := time.Now().Unix()
	name := utils.AnyToString(apiData["name"])
	method := utils.AnyToString(apiData["method"])
	apiUrl := utils.AnyToString(apiData["url"])
	// apiType := utils.AnyToString(apiData["type"])
	headerParamsList := utils.AnyToArr(apiData["header_params"])
	// queryParamsList := utils.AnyToArr(apiData["query_params"]) // todo
	bodyParamsList := utils.AnyToArr(apiData["body_params"])
	//bodyRaw := utils.AnyToString(apiData["body_raw"])
	bodyRawExample := utils.AnyToString(apiData["body_raw_example"])
	rawContentType := utils.AnyToString(apiData["raw_content_type"])
	// cookieParams := utils.AnyToArr(apiData["cookie_params"]) // todo
	responseDoc := utils.AnyToString(apiData["response_doc"])
	responseExample := utils.AnyToString(apiData["response_example"])
	responseExampleParams := utils.AnyToArr(apiData["response_example_params"])
	//markdownContent := utils.AnyToString(apiData["markdown_content"])

	doc := &entity.DocumentContent{
		DocId:           utils.IDMd5(),
		ProjectId:       project.ProjectId,
		Name:            name,
		Url:             apiUrl,
		Method:          method,
		Description:     responseDoc,
		DescriptionHtml: responseDoc,
		ReqHeader:       make([]*entity.ReqHeaderItem, 0),
		//ReqType                   define.ReqTypeCode        `json:"reqType"`                   // 请求类型
		//ReqBodyJson               string                    `json:"reqBodyJson"`               // 请求参数 - json
		//ReqBodyText               string                    `json:"reqBodyText"`               // 请求参数 - text
		//ReqBodyFormData           []*FormDataItem           `json:"reqBodyFormData"`           // 请求参数 - form-data
		//ReqBodyXWWWFormUrlEncoded []*XWWWFormUrlEncodedItem `json:"reqBodyXWWWFormUrlEncoded"` // 请求参数 - x-www-form-urlencoded"
		//ReqBodyXml                string                    `json:"reqBodyXml"`                // 请求参数 - xml
		//ReqBodyRaw                string                    `json:"reqBodyRaw"`                // 请求参数 - raw
		//ReqBodyBinary             string                    `json:"reqBodyBinary"`             // 请求参数 - binary
		//ReqBodyGraphQL            string                    `json:"reqBodyGraphQL"`            // 请求参数 - GraphQL
		ReqBodyInfo: make([]*entity.BodyInfoItem, 0),
		Resp:        make([]*entity.RespItem, 0),
		CreateTime:  now,
		UserAcc:     userAcc,
	}

	for i, headerParam := range headerParamsList {

		headerParamMap := utils.AnyToMap(headerParam)
		key := utils.AnyToString(headerParamMap["key"])
		//value := utils.AnyToString(headerParamMap["value"])
		desc := utils.AnyToString(headerParamMap["desc"])
		headerType := utils.AnyToString(headerParamMap["type"])
		eg := utils.AnyToString(headerParamMap["eg"])
		require := utils.AnyToInt(headerParamMap["require"])

		doc.ReqHeader = append(doc.ReqHeader, &entity.ReqHeaderItem{
			Field:       key,        // 字段
			VarType:     headerType, // 类型
			Description: desc,       // 描述
			Example:     eg,         // 示例
			IsRequired:  require,    // 是否必填 1:必填
			Sort:        i,
			IsOpen:      1,
		})

	}

	for i, bodyParam := range bodyParamsList {
		bodyParamMap := utils.AnyToMap(bodyParam)
		key := utils.AnyToString(bodyParamMap["key"])
		desc := utils.AnyToString(bodyParamMap["desc"])
		paramType := utils.AnyToString(bodyParamMap["type"])
		eg := utils.AnyToString(bodyParamMap["eg"])
		require := utils.AnyToInt(bodyParamMap["require"])

		doc.ReqBodyInfo = append(doc.ReqBodyInfo, &entity.BodyInfoItem{
			Field:       key,       // 字段
			VarType:     paramType, // 类型
			Description: desc,      // 描述
			Example:     eg,        // 示例
			IsRequired:  require,   // 是否必填 1:必填
			Sort:        i,         // 排序
			IsOpen:      1,         // 是否启用
		})

	}

	rawContentType = strings.Split(rawContentType, "(")[0]

	switch rawContentType {
	case "Form-Data":
		doc.ReqType = define.ReqTypeFormData
	case "Form-UrlEncoded":
		doc.ReqType = define.ReqTypeXWWWFormUrlEncoded
	case "JSON":
		doc.ReqType = define.ReqTypeJson
		doc.ReqBodyJson = bodyRawExample
	case "XML":
		doc.ReqType = define.ReqTypeXml
	case "Text":
		doc.ReqType = define.ReqTypeText
	case "HTML":
		doc.ReqType = define.ReqTypeText
	case "JavaScript":
		doc.ReqType = define.ReqTypeRaw
	}

	resp := &entity.RespItem{
		RespType:     define.ReqTypeJson, // 默认json
		RespBody:     responseExample,
		RespBodyInfo: make([]*entity.BodyInfoItem, 0),
	}

	for i, responseParam := range responseExampleParams {

		responseParamMap := utils.AnyToMap(responseParam)
		key := utils.AnyToString(responseParamMap["key"])
		desc := utils.AnyToString(responseParamMap["desc"])
		paramType := utils.AnyToString(responseParamMap["type"])
		require := utils.AnyToInt(responseParamMap["require"])

		resp.RespBodyInfo = append(resp.RespBodyInfo, &entity.BodyInfoItem{
			Field:       key,       // 字段
			VarType:     paramType, // 类型
			Description: desc,      // 描述
			Example:     "",        // 示例
			IsRequired:  require,   // 是否必填 1:必填
			Sort:        i,         // 排序
			IsOpen:      1,         // 是否启用
		})

	}

	if newDirId, ok := createDir(project, dirName); ok {
		dirId = newDirId
	}

	f(project, doc, dirId)

}
