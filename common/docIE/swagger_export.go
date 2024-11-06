package docIE

import (
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/entity"
	"strings"
)

type SwaggerExport struct {
	Info                map[string]interface{} `json:"info"`
	Tags                []interface{}          `json:"tags"`
	Paths               map[string]interface{} `json:"paths"`
	Swagger             string                 `json:"swagger"`
	Definitions         map[string]interface{} `json:"definitions"`
	SecurityDefinitions map[string]interface{} `json:"securityDefinitions"`
	XComponents         map[string]interface{} `json:"x-components"`
}

type PathsItem struct {
	Summary     string                 `json:"summary"`
	Deprecated  bool                   `json:"deprecated"`
	Description string                 `json:"description"`
	Tags        []interface{}          `json:"tags"`
	Parameters  []interface{}          `json:"parameters"`
	Responses   map[string]interface{} `json:"responses"`
	Security    []interface{}          `json:"security"`
	Consumes    []interface{}          `json:"consumes"`
	Produces    []interface{}          `json:"produces"`
}

type Schema struct {
	Type       string                            `json:"type"`
	Properties map[string]*SwaggerPropertiesData `json:"properties"`
	Required   []interface{}                     `json:"required"`
}

type SwaggerPropertiesData struct {
	Type        string  `json:"type"`
	Items       *Schema `json:"items"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}

func NewSwaggerExport() *SwaggerExport {
	return &SwaggerExport{}
}

func (obj *SwaggerExport) Export(pid string) interface{} {

	resp := &SwaggerExport{
		Info:                make(map[string]interface{}),
		Tags:                make([]interface{}, 0),
		Paths:               make(map[string]interface{}),
		Swagger:             "2.0",
		Definitions:         make(map[string]interface{}),
		SecurityDefinitions: make(map[string]interface{}),
		XComponents:         make(map[string]interface{}),
	}

	project, err := dao.NewProjectDao().GetByProjectId(pid)
	if err != nil {
		log.Error(err)
		return resp
	}

	resp.Info["title"] = project.Name
	resp.Info["description"] = project.Description
	resp.Info["version"] = ""

	dirList, err := dao.NewDirDao().GetAll(pid)
	if err != nil {
		log.Error(err)
		return resp
	}

	dirMap := make(map[string]*entity.DocumentDir)

	for _, dirItem := range dirList {
		resp.Tags = append(resp.Tags, dirItem.DirName)
		dirMap[dirItem.DirId] = dirItem
	}

	docIds, err := dao.NewDocDao().GetProjectAllDocId(pid)
	if err != nil {
		log.Error(err)
		return resp
	}

	for _, docId := range docIds {

		docObj, docErr := dao.NewDocDao().GetDocumentContent(pid, docId)
		if docErr != nil {
			log.Error(docErr)
			continue
		}

		docDirInfo, dirErr := dao.NewDocDao().GetDocument(pid, docId)
		if dirErr != nil {
			log.Error(dirErr)
			continue
		}

		item := &PathsItem{
			Summary:     docObj.Name,
			Deprecated:  false,
			Description: docObj.Description,
			Tags:        make([]interface{}, 0),
			Parameters:  make([]interface{}, 0),
			Responses:   make(map[string]interface{}),
			Security:    make([]interface{}, 0),
			Consumes:    make([]interface{}, 0),
			Produces:    make([]interface{}, 0),
		}

		item.Consumes = append(item.Consumes, docObj.ReqType.GetName())

		if dirData, ok := dirMap[docDirInfo.DirId]; ok {
			item.Tags = append(item.Tags, dirData.DirName)
		}

		for _, reqHeaderItem := range docObj.ReqHeader {

			required := false
			if reqHeaderItem.IsRequired == 1 {
				required = true
			}

			item.Parameters = append(item.Parameters, map[string]interface{}{
				"name":        reqHeaderItem.Field,
				"in":          "header",
				"description": reqHeaderItem.Description,
				"required":    required,
				"type":        reqHeaderItem.VarType,
				"x-example":   reqHeaderItem.Example,
			})
		}

		reqSchema := &Schema{
			Type:       "obj",
			Properties: make(map[string]*SwaggerPropertiesData),
			Required:   make([]interface{}, 0),
		}

		for _, v := range docObj.ReqBodyInfo {
			if !strings.Contains(v.Field, ".") {
				reqSchema.Properties[v.Field] = &SwaggerPropertiesData{
					Type: v.VarType,
					Items: &Schema{
						Type:       "obj",
						Properties: make(map[string]*SwaggerPropertiesData),
						Required:   make([]interface{}, 0),
					},
					Title:       v.Description,
					Description: v.Example,
				}
			} else {
				fList := strings.Split(v.Field, ".")

				swaggerCase(fList, 0, reqSchema.Properties, v)

			}
		}

		reqBody := map[string]interface{}{
			"name":   "body",
			"in":     "body",
			"schema": reqSchema,
		}

		item.Parameters = append(item.Parameters, reqBody)

		respBody := &entity.RespItem{}

		if len(docObj.Resp) > 0 {
			respBody = docObj.Resp[0]
		}

		respProperties := make(map[string]*SwaggerPropertiesData)
		for _, v := range respBody.RespBodyInfo {
			if !strings.Contains(v.Field, ".") {
				respProperties[v.Field] = &SwaggerPropertiesData{
					Type: v.VarType,
					Items: &Schema{
						Properties: make(map[string]*SwaggerPropertiesData),
						Required:   make([]interface{}, 0),
					},
					Title:       v.Description,
					Description: v.Example,
				}
			} else {
				fList := strings.Split(v.Field, ".")

				swaggerCase(fList, 0, respProperties, v)

			}
		}

		item.Responses["200"] = map[string]interface{}{
			"description": "",
			"schema": map[string]interface{}{
				"type":       "object",
				"properties": respProperties,
				"required":   make([]interface{}, 0),
			},
			"headers": map[string]interface{}{},
		}

		resp.Paths[docObj.Url] = map[string]interface{}{
			docObj.Method: item,
		}

	}

	return resp
}

func (obj *SwaggerExport) ExportJson(pid string) string {
	resp := obj.Export(pid)
	jsonStr, _ := utils.AnyToJson(resp)
	return jsonStr
}

func swaggerCase(fItemList []string, i int, data map[string]*SwaggerPropertiesData, v *entity.BodyInfoItem) {
	if i > len(fItemList) {
		return
	}

	itemObj, ok := data[fItemList[i]]

	if !ok {

		data[fItemList[i]] = &SwaggerPropertiesData{
			Type:        v.VarType,
			Title:       v.Description,
			Description: v.Example,
			Items: &Schema{
				Properties: make(map[string]*SwaggerPropertiesData),
				Required:   make([]interface{}, 0),
			},
		}

	} else {

		i += 1

		if itemObj.Items == nil {
			itemObj.Items = &Schema{
				Properties: make(map[string]*SwaggerPropertiesData),
				Required:   make([]interface{}, 0),
			}
		}

		swaggerCase(fItemList, i, itemObj.Items.Properties, v)
	}
}
