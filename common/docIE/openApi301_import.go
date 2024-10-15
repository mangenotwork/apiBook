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

open-api 3.0.* 导入

https://spec.openapis.org/oas/v3.0.0

局限:
	- 不能导出目录结构
	- 全局导入都在默认目录下
	- 增量导入都在默认目录下
	- 未使用 tag
	- 未使用 components
*/

type OpenApi301Import struct {
	Openapi    string                 `json:"openapi"`
	Info       map[string]interface{} `json:"info"`
	Tags       []interface{}          `json:"tags"`
	Paths      map[string]interface{} `json:"paths"`
	Components map[string]interface{} `json:"components"`
	Servers    []interface{}          `json:"servers"`
}

func NewOpenApi301Import() *OpenApi301Import {
	return &OpenApi301Import{}
}

func (obj *OpenApi301Import) Whole(text, userAcc string, private define.ProjectPrivateCode) error {
	err := obj.analysis(text)
	if err != nil {
		return err
	}

	project := obj.getProject(userAcc, private)

	log.Info(project)

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

func (obj *OpenApi301Import) Increment(text, pid, userAcc, dirId string) error {
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

			err := dao.NewDocDao().Create(documentData, doc)
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

func (obj *OpenApi301Import) analysis(text string) error {
	err := json.Unmarshal([]byte(text), &obj)
	if err != nil {
		log.Error(err)
		return err
	}

	if obj.Openapi != "3.0.1" {
		return fmt.Errorf("openapi is not 3.0.1")
	}

	return nil
}

func (obj *OpenApi301Import) analysisDoc(project *entity.Project, userAcc string, f func(doc *entity.DocumentContent)) {

	now := time.Now().Unix()

	for apiUrl, apiData := range obj.Paths {

		apiDataMap := utils.AnyToMap(apiData)

		for method, apiInfo := range apiDataMap {

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
				}

			}

			if requestBody, ok := apiInfoMap["requestBody"]; ok {
				requestBodyContent := utils.AnyToMap(utils.AnyToMap(requestBody)["content"])

				//log.Info("requestBodyContent = ", requestBodyContent)
				//log.Info("requestBodyContent len ", len(requestBodyContent))

				for requestBodyType, requestBodyData := range requestBodyContent {

					requestBodyDataMap := utils.AnyToMap(requestBodyData)
					schema := utils.AnyToMap(requestBodyDataMap["schema"])

					example, exampleOk := requestBodyDataMap["example"]
					if !exampleOk {
						example = ""
					}

					//log.Info("requestBodyType = ", requestBodyType)

					for _, v := range define.ReqTypeArray {
						if strings.Contains(requestBodyType, v) {
							doc.ReqType = define.ReqTypeCode(v)
						}
					}

					switch requestBodyType {
					case "application/json":
						exampleJson, err := utils.AnyToJson(example)
						if err != nil {
							log.Error("example 转json失败")
						}
						doc.ReqBodyJson = exampleJson
					}

					//log.Info("example = ", example)

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

				// todo 忽略 返回状态码，description，headers;
				// todo 目录响应参数只有一个，以后扩展多个此处需要进行适配

				responsesDataMap := utils.AnyToMap(responsesData)

				responsesContent := utils.AnyToMap(responsesDataMap["content"])

				//log.Info("responsesContent = ", responsesContent)

				resp := &entity.RespItem{}

				for responsesContentType, responsesContentData := range responsesContent {
					responsesContentDataMap := utils.AnyToMap(responsesContentData)

					schema := utils.AnyToMap(responsesContentDataMap["schema"])

					examples, examplesOk := responsesContentDataMap["examples"]
					if examplesOk {

						examplesMap := utils.AnyToMap(examples)
						for _, examplesData := range examplesMap {

							examplesDataMap := utils.AnyToMap(examplesData)
							if exampleValue, exampleValueOk := examplesDataMap["value"]; exampleValueOk {

								switch responsesContentType {
								case "application/json":
									exampleJson, err := utils.AnyToJson(exampleValue)
									if err != nil {
										log.Info("example 转json失败")
									}
									resp.RespBody = exampleJson
								}

							}

						}

					}

					//log.Info("responsesContentType = ", responsesContentType)

					for _, v := range define.ReqTypeArray {
						if strings.Contains(responsesContentType, v) {
							resp.RespType = define.ReqTypeCode(v)
						}
					}

					properties := utils.AnyToMap(schema["properties"])
					//log.Info("properties = ", properties)

					analysisProperties("", properties, &resp.RespBodyInfo)

					doc.Resp = append(doc.Resp, resp)

				}

			}

			log.Info("doc = ", doc)

			//log.Info("doc.ReqBodyInfo = ", doc.ReqBodyInfo)
			//for _, v := range doc.ReqBodyInfo {
			//	log.Info(v)
			//}

			//if len(doc.Resp) > 0 {
			//	log.Info("doc.Resp[0] = ", doc.Resp[0])
			//	for _, v := range doc.Resp[0].RespBodyInfo {
			//		log.Info(v)
			//	}
			//}

			log.Info("__________________________________")

			f(doc)

		}

	}
}

func analysisProperties(pKey string, properties map[string]interface{}, reqBodyInfo *[]*entity.BodyInfoItem) {
	if len(properties) == 0 {
		return
	}
	for key, value := range properties {
		//log.Info("===============================")
		nowKey := key
		if pKey != "" {
			nowKey = pKey + "." + key
		}
		//log.Info("key = ", nowKey)

		valueMap := utils.AnyToMap(value)
		valueType := utils.AnyToString(valueMap["type"])
		items := utils.AnyToMap(valueMap["items"])
		valueTitle := utils.AnyToString(valueMap["title"])
		valueDescription := utils.AnyToString(valueMap["description"])

		//log.Info("valueType = ", valueType)
		//log.Info("valueTitle = ", valueTitle)
		//log.Info("valueDescription = ", valueDescription)

		*reqBodyInfo = append(*reqBodyInfo, &entity.BodyInfoItem{
			Field:       nowKey,                                             // 字段
			VarType:     valueType,                                          // 类型
			Description: fmt.Sprintf("%s %s", valueTitle, valueDescription), // 描述
			Example:     "",                                                 // 示例
			IsRequired:  1,                                                  // 是否必填 1:必填
			Sort:        len(*reqBodyInfo) + 1,                              // 排序
			IsOpen:      1,                                                  // 是否启用
		})

		//log.Info("===============================\n\n")

		p := utils.AnyToMap(items["properties"])
		analysisProperties(nowKey, p, reqBodyInfo)

	}
}

func (obj *OpenApi301Import) getProject(userAcc string, private define.ProjectPrivateCode) *entity.Project {
	return &entity.Project{
		ProjectId:     utils.IDMd5(),
		Name:          utils.AnyToString(obj.Info["title"]),
		Description:   fmt.Sprintf("version: %s; %s", utils.AnyToString(obj.Info["version"]), utils.AnyToString(obj.Info["description"])),
		CreateUserAcc: userAcc,
		CreateDate:    utils.NowDate(),
		Private:       private,
	}
}
