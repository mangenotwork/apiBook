package dao

import (
	"apiBook/common/db"
	"apiBook/common/log"
	"apiBook/internal/define"
	"apiBook/internal/entity"
)

type DocDao struct {
}

func NewDocDao() *DocDao {
	return &DocDao{}
}

func (dao *DocDao) Create(data *entity.Document, content *entity.DocumentContent) error {
	log.Info(db.GetDocumentTable(data.ProjectId), data.DocId)
	err := db.DB.Set(db.GetDocumentTable(data.ProjectId), data.DocId, data)
	if err != nil {
		return err
	}

	log.Info(db.GetDocumentContentTable(content.ProjectId), content.DocId)
	err = db.DB.Set(db.GetDocumentContentTable(content.ProjectId), content.DocId, content)
	if err != nil {
		return err
	}

	return nil
}

func (dao *DocDao) GetDocument(pid, docId string) (*entity.Document, error) {
	result := &entity.Document{}
	err := db.DB.Get(db.GetDocumentTable(pid), docId, &result)
	return result, err
}

func (dao *DocDao) GetDocumentContent(pid, docId string) (*entity.DocumentContent, error) {
	result := &entity.DocumentContent{}
	err := db.DB.Get(db.GetDocumentContentTable(pid), docId, &result)
	return result, err
}

func (dao *DocDao) GetDocList(pid string, list []*entity.DocumentDirItem) []*entity.Document {
	result := make([]*entity.Document, 0)

	for _, v := range list {

		data, err := dao.GetDocument(pid, v.DocId)
		if err != nil {
			log.Error(err)
			continue
		}

		result = append(result, data)
	}

	return result
}

func (dao *DocDao) Modify(content *entity.DocumentContent) error {
	oldDoc, err := dao.GetDocument(content.ProjectId, content.DocId)
	if err != nil {
		return err
	}

	oldDoc.Url = content.Url
	oldDoc.Name = content.Name
	oldDoc.Method = content.Method

	err = db.DB.Set(db.GetDocumentTable(content.ProjectId), content.DocId, oldDoc)
	if err != nil {
		return err
	}

	oldDocContent, err := dao.GetDocumentContent(content.ProjectId, content.DocId)
	if err != nil {
		return err
	}

	oldDocContent.Name = content.Name
	oldDocContent.Url = content.Url
	oldDocContent.Method = content.Method
	oldDocContent.Description = content.Description
	oldDocContent.ReqHeader = content.ReqHeader
	oldDocContent.ReqType = content.ReqType
	oldDocContent.ReqBodyJson = content.ReqBodyJson
	oldDocContent.ReqBodyText = content.ReqBodyText
	oldDocContent.ReqBodyFormData = content.ReqBodyFormData
	oldDocContent.ReqBodyXWWWFormUrlEncoded = content.ReqBodyXWWWFormUrlEncoded
	oldDocContent.ReqBodyXml = content.ReqBodyXml
	oldDocContent.ReqBodyRaw = content.ReqBodyRaw
	oldDocContent.ReqBodyBinary = content.ReqBodyBinary
	oldDocContent.ReqBodyGraphQL = content.ReqBodyGraphQL
	oldDocContent.ReqBodyInfo = content.ReqBodyInfo
	oldDocContent.Resp = content.Resp

	err = db.DB.Set(db.GetDocumentContentTable(content.ProjectId), content.DocId, oldDocContent)
	if err != nil {
		return err
	}

	return nil
}

func (dao *DocDao) Delete(pid, dirId, docId string) error {
	oldDoc, err := dao.GetDocument(pid, docId)
	if err != nil {
		return err
	}

	oldDoc.DirId = db.GetDocumentDirItemTable(define.GetDirRecycleBinKey(pid))
	err = db.DB.Set(db.GetDocumentTable(pid), docId, oldDoc)
	if err != nil {
		return err
	}

	err = db.DB.Set(db.GetDocumentDirTable(define.GetDirRecycleBinKey(pid)), dirId, 1)
	if err != nil {
		return err
	}

	err = db.DB.Delete(db.GetDocumentDirTable(pid), dirId)
	if err != nil {
		return err
	}

	return nil
}

func (dao *DocDao) ChangeDir(pid, dirId, dirIdNew, docId string) error {
	oldDoc, err := dao.GetDocument(pid, docId)
	if err != nil {
		return err
	}

	oldDoc.DirId = dirIdNew
	err = db.DB.Set(db.GetDocumentTable(pid), docId, oldDoc)
	if err != nil {
		return err
	}

	err = db.DB.Set(db.GetDocumentDirTable(pid), dirIdNew, 1)
	if err != nil {
		return err
	}

	err = db.DB.Delete(db.GetDocumentDirTable(pid), dirId)
	if err != nil {
		return err
	}

	return nil
}
