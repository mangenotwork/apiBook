package docIE

import (
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
)

type ApiBookExport struct {
	Project map[string]interface{}   `json:"project"`
	Dir     []map[string]interface{} `json:"dir"`
	Doc     []map[string]interface{} `json:"doc"`
}

func NewApiBookExport() *ApiBookExport {
	return &ApiBookExport{}
}

func (obj *ApiBookExport) Export(pid string) interface{} {
	resp := &ApiBookExport{
		Project: make(map[string]interface{}),
		Dir:     make([]map[string]interface{}, 0),
		Doc:     make([]map[string]interface{}, 0),
	}

	project, err := dao.NewProjectDao().GetByProjectId(pid)
	if err != nil {
		log.Error(err)
		return resp
	}

	resp.Project["name"] = project.Name
	resp.Project["description"] = project.Description
	resp.Project["createDate"] = project.CreateDate

	dirList, err := dao.NewDirDao().GetAll(pid)
	if err != nil {
		log.Error(err)
		return resp
	}

	for _, dirItem := range dirList {
		resp.Dir = append(resp.Dir, map[string]interface{}{
			"dirName": dirItem.DirName,
			"dirId":   dirItem.DirId,
		})
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

		resp.Doc = append(resp.Doc, map[string]interface{}{
			"docId":       docObj.DocId,
			"dirId":       docDirInfo.DirId,
			"name":        docObj.Name,
			"url":         docObj.Url,
			"method":      docObj.Method,
			"description": docObj.Description,
			"reqHeader":   docObj.ReqHeader,
			"reqType":     docObj.ReqType,
			"reqBodyJson": docObj.ReqBodyJson,
			"reqBodyInfo": docObj.ReqBodyInfo,
			"resp":        docObj.Resp,
		})
	}

	return resp
}

func (obj *ApiBookExport) ExportJson(pid string) string {
	resp := obj.Export(pid)
	jsonStr, _ := utils.AnyToJson(resp)
	return jsonStr
}
