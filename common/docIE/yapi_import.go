package docIE

import (
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"encoding/json"
	"fmt"
	"time"
)

/*

yapi 导入,仅支持json

https://hellosean1025.github.io/yapi/

*/

type YApiImport struct {
	Data []interface{}
}

func NewYApiImport() *YApiImport {
	return &YApiImport{
		Data: make([]interface{}, 0),
	}
}

func (obj *YApiImport) Whole(text, userAcc string, private define.ProjectPrivateCode) error {

	err := obj.analysis(text)
	if err != nil {
		log.Error(err)
		return err
	}

	project := obj.analysisProject(userAcc, private)

	if dao.NewProjectDao().HasName(project.Name) {
		project.Name = fmt.Sprintf("%s-%s", project.Name, utils.NowDateNotLine())
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

	obj.analysisDoc(project, userAcc, "",
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

func (obj *YApiImport) Increment(text, pid, userAcc, dirId string) error {

	project, err := dao.NewProjectDao().Get(pid, userAcc, false)
	if err != nil {
		log.Error("获取项目失败, err = ", err)
		return err
	}

	err = obj.analysis(text)
	if err != nil {
		return err
	}

	if dirId == "" {
		dirId = define.GetDirDefault(pid)
	}

	obj.analysisDoc(project, userAcc, dirId,
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

func (obj *YApiImport) analysis(text string) error {
	err := json.Unmarshal([]byte(text), &obj.Data)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (obj *YApiImport) analysisProject(userAcc string, private define.ProjectPrivateCode) *entity.Project {
	day := utils.NowDateYMDStr()
	return &entity.Project{
		ProjectId:     utils.IDMd5(),
		Name:          "yapi导入" + day,
		Description:   "",
		CreateUserAcc: userAcc,
		CreateDate:    day,
		Private:       private,
	}
}

func (obj *YApiImport) analysisDoc(project *entity.Project, userAcc, dirId string,
	createDir func(project *entity.Project, dirName string) (string, bool),
	f func(project *entity.Project, doc *entity.DocumentContent, dirId string)) {

	now := time.Now().Unix()

	for _, v := range obj.Data {
		vMap := utils.AnyToMap(v)
		dirName := utils.AnyToString(vMap["name"])
		log.Info(dirName)

		list := utils.AnyToArr(vMap["list"])
		for _, item := range list {
			itemMap := utils.AnyToMap(item)

			path := utils.AnyToString(itemMap["path"])
			method := utils.AnyToString(itemMap["method"])
			title := utils.AnyToString(itemMap["title"])
			desc := utils.AnyToString(itemMap["desc"])
			reqParams := utils.AnyToArr(itemMap["req_params"])
			reqHeaders := utils.AnyToArr(itemMap["req_headers"])
			reqBodyType := utils.AnyToString(itemMap["req_body_type"])
			resBodyType := utils.AnyToString(itemMap["res_body_type"])
			//reqBodyOther := utils.AnyToString(itemMap["req_body_other"])
			//resBody := utils.AnyToString(itemMap["res_body"])

			//log.Info("path : ", path)
			//log.Info("method : ", method)
			//log.Info("title : ", title)
			//log.Info("desc : ", desc)
			//log.Info("reqParams : ", reqParams)
			//log.Info("reqHeaders : ", reqHeaders)
			//log.Info("reqBodyType : ", reqBodyType)
			//log.Info("resBodyType : ", resBodyType)
			//log.Info("reqBodyOther : ", reqBodyOther)
			//log.Info("resBody : ", resBody)

			doc := &entity.DocumentContent{
				DocId:           utils.IDMd5(),
				ProjectId:       project.ProjectId,
				Name:            title,
				Url:             path,
				Method:          method,
				Description:     desc,
				DescriptionHtml: desc,
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

			for i, reqHeaderItem := range reqHeaders {
				reqHeaderItemMap := utils.AnyToMap(reqHeaderItem)
				reqHeaderName := utils.AnyToString(reqHeaderItemMap["name"])
				reqHeaderRequired := utils.AnyToInt(reqHeaderItemMap["required"])
				reqHeaderValue := ""
				if value, ok := reqHeaderItemMap["value"]; ok {
					reqHeaderValue = utils.AnyToString(value)
				}
				reqHeaderDesc := ""
				if reqHeaderDescVal, ok := reqHeaderItemMap["desc"]; ok {
					reqHeaderDesc = utils.AnyToString(reqHeaderDescVal)
				}

				doc.ReqHeader = append(doc.ReqHeader, &entity.ReqHeaderItem{
					Field:       reqHeaderName,     // 字段
					VarType:     "string",          // 类型
					Description: reqHeaderDesc,     // 描述
					Example:     reqHeaderValue,    // 示例
					IsRequired:  reqHeaderRequired, // 是否必填 1:必填
					Sort:        i,
					IsOpen:      1,
				})

			}

			for i, reqParamsItem := range reqParams {
				reqParamsItemMap := utils.AnyToMap(reqParamsItem)
				reqParamsName := utils.AnyToString(reqParamsItemMap["name"])
				reqParamsRequired := 1
				if reqParamsRequiredVal, ok := reqParamsItemMap["required"]; ok {
					reqParamsRequired = utils.AnyToInt(reqParamsRequiredVal)
				}
				reqParamsValue := ""
				if value, ok := reqParamsItemMap["value"]; ok {
					reqParamsValue = utils.AnyToString(value)
				}
				reqParamsDesc := ""
				if reqHeaderDescVal, ok := reqParamsItemMap["desc"]; ok {
					reqParamsDesc = utils.AnyToString(reqHeaderDescVal)
				}

				doc.ReqBodyInfo = append(doc.ReqBodyInfo, &entity.BodyInfoItem{
					Field:       reqParamsName,     // 字段
					VarType:     "string",          // 类型
					Description: reqParamsDesc,     // 描述
					Example:     reqParamsValue,    // 示例
					IsRequired:  reqParamsRequired, // 是否必填 1:必填
					Sort:        i,                 // 排序
					IsOpen:      1,                 // 是否启用
				})
			}

			switch reqBodyType {
			case "json":
				doc.ReqType = define.ReqTypeJson
				//doc.ReqBodyJson = reqBodyOther
			case "form":
				doc.ReqType = define.ReqTypeFormData
			case "raw":
				doc.ReqType = define.ReqTypeRaw
				//doc.ReqBodyRaw = reqBodyOther
			}

			respData := &entity.RespItem{}

			switch resBodyType {
			case "json":
				respData.RespType = define.ReqTypeJson
			case "form":
				respData.RespType = define.ReqTypeFormData
			case "raw":
				respData.RespType = define.ReqTypeRaw
			}

			//respData.RespBody = resBody

			doc.Resp = append(doc.Resp, respData)

			log.Info("doc : ", doc)

			if newDirId, ok := createDir(project, dirName); ok {
				dirId = newDirId
			}

			f(project, doc, dirId)
		}
	}

}
