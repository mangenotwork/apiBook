package docIE

import (
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/entity"
	"encoding/json"
	"strings"
)

type OpenApiExport struct {
	Openapi    string                 `json:"openapi"`
	Info       *OpenApiInfo           `json:"info"`
	Tags       []interface{}          `json:"tags"`
	Paths      map[string]interface{} `json:"paths"`
	Components map[string]interface{} `json:"components"`
	Servers    []interface{}          `json:"servers"`
}

type OpenApiInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DocItem struct {
	Summary     string                 `json:"summary"`
	Deprecated  bool                   `json:"deprecated"`
	Description string                 `json:"description"`
	Tags        []interface{}          `json:"tags"`
	Parameters  []*ParameterItem       `json:"parameters"`
	RequestBody map[string]interface{} `json:"requestBody"`
	Responses   map[string]interface{} `json:"responses"`
	Security    []interface{}          `json:"security"`
}

type ParameterItem struct {
	Name        string                 `json:"name"`
	In          string                 `json:"in"`
	Description string                 `json:"description"`
	Required    bool                   `json:"required"`
	Example     string                 `json:"example"`
	Schema      map[string]interface{} `json:"schema"`
}

type SchemaData struct {
	Type       string                     `json:"type"`
	Properties map[string]*PropertiesData `json:"properties"`
	Required   []interface{}              `json:"required"`
}

type PropertiesData struct {
	Type        string      `json:"type"`
	Items       *SchemaData `json:"items"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
}

func NewOpenApiExport() *OpenApiExport {
	return &OpenApiExport{}
}

func (obj *OpenApiExport) Export(pid string) interface{} {
	resp := &OpenApiExport{
		Openapi:    "3.0.1",
		Info:       &OpenApiInfo{},
		Tags:       make([]interface{}, 0),
		Paths:      make(map[string]interface{}),
		Components: make(map[string]interface{}),
		Servers:    make([]interface{}, 0),
	}

	project, err := dao.NewProjectDao().GetByProjectId(pid)
	if err != nil {
		log.Error(err)
		return resp
	}

	resp.Info.Title = project.Name
	resp.Info.Description = project.Description

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

		parameters := make([]*ParameterItem, 0)

		for _, reqHeaderItem := range docObj.ReqHeader {

			required := false
			if reqHeaderItem.IsRequired == 1 {
				required = true
			}

			parameters = append(parameters, &ParameterItem{
				Name:        reqHeaderItem.Field,
				In:          "header",
				Description: reqHeaderItem.Description,
				Required:    required,
				Example:     reqHeaderItem.Example,
				Schema: map[string]interface{}{
					"type": reqHeaderItem.VarType,
				},
			})
		}

		requestBody := make(map[string]interface{})

		jsonExampleMap := make(map[string]interface{})
		_ = json.Unmarshal([]byte(docObj.ReqBodyJson), &jsonExampleMap)

		reqProperties := make(map[string]*PropertiesData)

		for _, v := range docObj.ReqBodyInfo {
			if !strings.Contains(v.Field, ".") {
				reqProperties[v.Field] = &PropertiesData{
					Type: v.VarType,
					Items: &SchemaData{
						Properties: make(map[string]*PropertiesData),
						Required:   make([]interface{}, 0),
					},
					Title:       v.Description,
					Description: v.Example,
				}
			} else {
				fList := strings.Split(v.Field, ".")

				myCase(fList, 0, reqProperties, v)

			}
			//log.Info(v)
		}

		requestBodyContent := map[string]interface{}{
			docObj.ReqType.GetName(): map[string]interface{}{
				"schema": map[string]interface{}{
					"type":       "object",
					"properties": reqProperties,
					"required":   make([]interface{}, 0),
				},
				"example": jsonExampleMap,
			},
		}

		requestBody["content"] = requestBodyContent

		tags := make([]interface{}, 0)
		if dirData, ok := dirMap[docDirInfo.DirId]; ok {
			tags = append(tags, dirData.DirName)
		}

		respBody := &entity.RespItem{}

		if len(docObj.Resp) > 0 {
			respBody = docObj.Resp[0]
		}

		respProperties := make(map[string]*PropertiesData)
		for _, v := range respBody.RespBodyInfo {
			if !strings.Contains(v.Field, ".") {
				respProperties[v.Field] = &PropertiesData{
					Type: v.VarType,
					Items: &SchemaData{
						Properties: make(map[string]*PropertiesData),
						Required:   make([]interface{}, 0),
					},
					Title:       v.Description,
					Description: v.Example,
				}
			} else {
				fList := strings.Split(v.Field, ".")

				myCase(fList, 0, respProperties, v)

			}
		}

		respBodyContent := map[string]interface{}{
			respBody.RespType.GetName(): map[string]interface{}{
				"schema": map[string]interface{}{
					"type":       "object",
					"properties": respProperties,
					"required":   make([]interface{}, 0),
				},
				"example": respBody.RespBody,
			},
		}

		resp.Paths[docObj.Url] = map[string]interface{}{
			docObj.Method: &DocItem{
				Summary:     docObj.Name,
				Deprecated:  false,
				Description: docObj.Description,
				Tags:        tags,
				Parameters:  parameters,
				RequestBody: requestBody,
				Responses: map[string]interface{}{
					"200": map[string]interface{}{
						"description": "",
						"content":     respBodyContent,
						"headers":     map[string]interface{}{},
					},
				},
				Security: make([]interface{}, 0),
			},
		}

	}

	return resp
}

func (obj *OpenApiExport) ExportJson(pid string) string {
	resp := obj.Export(pid)
	jsonStr, _ := utils.AnyToJson(resp)
	return jsonStr
}

func myCase(fItemList []string, i int, data map[string]*PropertiesData, v *entity.BodyInfoItem) {
	//log.Info("fItemList [] = ", fItemList)
	//log.Info("i= ", i, "  | Field = ", fItemList[i])

	if i > len(fItemList) {
		return
	}

	itemObj, ok := data[fItemList[i]]

	//log.Info("itemObj = ", itemObj, ok)

	if !ok {

		data[fItemList[i]] = &PropertiesData{
			Type:        v.VarType,
			Title:       v.Description,
			Description: v.Example,
			Items: &SchemaData{
				Properties: make(map[string]*PropertiesData),
				Required:   make([]interface{}, 0),
			},
		}

	} else {

		i += 1

		if itemObj.Items == nil {
			itemObj.Items = &SchemaData{
				Properties: make(map[string]*PropertiesData),
				Required:   make([]interface{}, 0),
			}
		}

		myCase(fItemList, i, itemObj.Items.Properties, v)
	}
}
