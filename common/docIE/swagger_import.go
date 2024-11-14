package docIE

import (
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

/*

swagger 2.0 导入

*/

type SwaggerImport struct {
	Info                map[string]interface{} `json:"info"`
	Tags                []interface{}          `json:"tags"`
	Paths               map[string]interface{} `json:"paths"`
	Swagger             string                 `json:"swagger"`
	Definitions         map[string]interface{} `json:"definitions"`
	SecurityDefinitions map[string]interface{} `json:"securityDefinitions"`
	XComponents         map[string]interface{} `json:"x-components"`
}

func NewSwaggerImport() *SwaggerImport {
	return &SwaggerImport{}
}

func (obj *SwaggerImport) Whole(text, userAcc string, private define.ProjectPrivateCode) error {
	err := obj.analysis(text)
	if err != nil {
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

	obj.analysisDoc(project, userAcc,
		func(doc *entity.DocumentContent) {

			dirId := define.GetDirDefault(project.ProjectId)

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
		},
	)
	return nil
}

func (obj *SwaggerImport) Increment(text, pid, userAcc, dirId string) error {
	project, err := dao.NewProjectDao().Get(pid, userAcc, false)
	if err != nil {
		log.Error("获取项目失败, err = ", err)
		return err
	}

	err = obj.analysis(text)
	if err != nil {
		return err
	}

	obj.analysisDoc(project, userAcc,
		func(doc *entity.DocumentContent) {

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
		},
	)
	return nil
}

func (obj *SwaggerImport) analysis(text string) error {
	err := json.Unmarshal([]byte(text), &obj)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (obj *SwaggerImport) analysisProject(userAcc string, private define.ProjectPrivateCode) *entity.Project {
	return &entity.Project{
		ProjectId:     utils.IDMd5(),
		Name:          utils.AnyToString(obj.Info["title"]),
		Description:   fmt.Sprintf("version: %s; %s", utils.AnyToString(obj.Info["version"]), utils.AnyToString(obj.Info["description"])),
		CreateUserAcc: userAcc,
		CreateDate:    utils.NowDate(),
		Private:       private,
	}
}

func (obj *SwaggerImport) analysisDoc(project *entity.Project, userAcc string, f func(doc *entity.DocumentContent)) {
	now := time.Now().Unix()

	for apiUrl, apiData := range obj.Paths {

		apiDataMap := utils.AnyToMap(apiData)

		for method, apiInfo := range apiDataMap {
			//log.Info("now = ", now)
			//log.Info("apiUrl = ", apiUrl)
			//log.Info("method = ", method)
			//log.Info("apiInfo = ", apiInfo)

			apiInfoMap := utils.AnyToMap(apiInfo)
			apiName := utils.AnyToString(apiInfoMap["summary"])
			description := utils.AnyToString(apiInfoMap["description"])

			doc := &entity.DocumentContent{
				DocId:           utils.IDMd5(),
				ProjectId:       project.ProjectId,
				Name:            apiName,
				Url:             apiUrl,
				Method:          method,
				Description:     description,
				DescriptionHtml: description,
				ReqHeader:       make([]*entity.ReqHeaderItem, 0),
				ReqBodyInfo:     make([]*entity.BodyInfoItem, 0),
				Resp:            make([]*entity.RespItem, 0),
				CreateTime:      now,
				UserAcc:         userAcc,
			}

			reqTypeStr := ""
			consumes := utils.AnyToArr(apiInfoMap["consumes"])
			if len(consumes) > 0 {
				reqTypeStr = utils.AnyToString(consumes[0])
			}

			for _, v := range define.ReqTypeArray {
				if strings.Contains(reqTypeStr, v) {
					doc.ReqType = define.ReqTypeCode(v)
				}
			}

			respTypeStr := ""
			produces := utils.AnyToArr(apiInfoMap["produces"])
			if len(produces) > 0 {
				respTypeStr = utils.AnyToString(produces[0])
			}

			// 解析参数
			for i, parameterInfo := range utils.AnyToArr(apiInfoMap["parameters"]) {
				//log.Info("parameterInfo = ", parameterInfo)

				parameterInfoMap := utils.AnyToMap(parameterInfo)

				in := utils.AnyToString(parameterInfoMap["in"])

				name := utils.AnyToString(parameterInfoMap["name"])

				schema := utils.AnyToMap(parameterInfoMap["schema"])

				schemaType := utils.AnyToString(schema["type"])

				parameterDescription := utils.AnyToString(parameterInfoMap["description"])

				example := utils.AnyToString(parameterInfoMap["example"])

				required := utils.AnyToString(parameterInfoMap["required"])
				requiredInt := 0
				if required == "true" {
					requiredInt = 1
				}

				switch in {

				case "path":
					doc.ReqBodyInfo = append(doc.ReqBodyInfo, &entity.BodyInfoItem{
						Field:       name,                 // 字段
						VarType:     schemaType,           // 类型
						Description: parameterDescription, // 描述
						Example:     example,              // 示例
						IsRequired:  requiredInt,          // 是否必填 1:必填
						Sort:        i,                    // 排序
						IsOpen:      1,                    // 是否启用
					})

				case "header":
					doc.ReqHeader = append(doc.ReqHeader, &entity.ReqHeaderItem{
						Field:       name,                 // 字段
						VarType:     schemaType,           // 类型
						Description: parameterDescription, // 描述
						Example:     example,              // 示例
						IsRequired:  requiredInt,          // 是否必填 1:必填
						Sort:        i,
						IsOpen:      1,
					})

				case "query":
					doc.ReqBodyInfo = append(doc.ReqBodyInfo, &entity.BodyInfoItem{
						Field:       name,                 // 字段
						VarType:     schemaType,           // 类型
						Description: parameterDescription, // 描述
						Example:     example,              // 示例
						IsRequired:  requiredInt,          // 是否必填 1:必填
						Sort:        i,                    // 排序
						IsOpen:      1,                    // 是否启用
					})

				case "cookie":
					// todo....

				case "body":
					properties := utils.AnyToMap(schema["properties"])
					//log.Info("properties = ", properties)

					analysisProperties("", properties, &doc.ReqBodyInfo)
				}

			}

			// todo responses 多个只识别200  一个只识别第一个
			responsesMap := utils.AnyToMap(apiInfoMap["responses"])
			responsesFlag := false
			if len(responsesMap) > 0 {
				responsesFlag = true
			}

			for key, responsesData := range responsesMap {
				if responsesFlag {
					if key != "200" {
						continue
					}
				}

				responsesDataMap := utils.AnyToMap(responsesData)

				resp := &entity.RespItem{}

				schema := utils.AnyToMap(responsesDataMap["schema"])

				for _, v := range define.ReqTypeArray {
					if strings.Contains(respTypeStr, v) {
						resp.RespType = define.ReqTypeCode(v)
					}
				}

				properties := utils.AnyToMap(schema["properties"])

				analysisProperties("", properties, &resp.RespBodyInfo)

				doc.Resp = append(doc.Resp, resp)

			}

			log.Info("doc = ", doc)

			//for _, v := range doc.ReqBodyInfo {
			//	log.Info("ReqBodyInfo = ", v)
			//}

			//if len(doc.Resp) > 0 {
			//	for _, v := range doc.Resp[0].RespBodyInfo {
			//		log.Info("RespBodyInfo = ", v)
			//	}
			//} else {
			//	log.Error("not RespBodyInfo.")
			//}

			f(doc)
		}
	}

}
