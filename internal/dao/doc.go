package dao

import (
	"apiBook/common/db"
	"apiBook/common/log"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"encoding/json"
	"time"
)

type DocDao struct {
}

func NewDocDao() *DocDao {
	return &DocDao{}
}

func (dao *DocDao) Create(data *entity.Document, content *entity.DocumentContent) error {
	err := db.DB.Set(db.GetDocumentTable(data.ProjectId), data.DocId, data)
	if err != nil {
		return err
	}

	err = db.DB.Set(db.GetDocumentContentTable(content.ProjectId), content.DocId, content)
	if err != nil {
		return err
	}

	err = dao.AddDocumentSnapshot(content, content.UserAcc, define.OperationLogCreateDoc)
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

func (dao *DirDao) UpdateDocument(pid, docId string, data *entity.Document) error {
	return db.DB.Set(db.GetDocumentTable(pid), docId, data)
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

func (dao *DocDao) GetDocListByIds(pid string, list []string) []*entity.Document {
	result := make([]*entity.Document, 0)

	for _, v := range list {
		data, err := dao.GetDocument(pid, v)
		if err != nil {
			log.Error(err)
			continue
		}

		result = append(result, data)
	}

	return result
}

func (dao *DocDao) Modify(content *entity.DocumentContent, userAcc string) error {
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
	oldDocContent.DescriptionHtml = content.DescriptionHtml
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

	content.UserAcc = userAcc
	err = dao.AddDocumentSnapshot(content, userAcc, define.OperationLogCreateDoc)
	if err != nil {
		return err
	}

	return nil
}

func (dao *DocDao) Delete(pid, docId string) error {
	oldDoc, err := dao.GetDocument(pid, docId)
	if err != nil {
		return err
	}

	err = db.DB.Delete(db.GetDocumentTable(pid), docId)
	if err != nil {
		return err
	}

	err = db.DB.Delete(db.GetDocumentDirItemTable(oldDoc.DirId), docId)
	if err != nil {
		return err
	}

	err = db.DB.Delete(db.GetDocumentContentTable(pid), docId)
	if err != nil {
		return err
	}

	return nil
}

func (dao *DocDao) ChangeDir(pid, dirIdNew, docId string) error {
	oldDoc, err := dao.GetDocument(pid, docId)
	if err != nil {
		return err
	}

	oldDir := oldDoc.DirId

	if oldDir == dirIdNew {
		return nil
	}

	oldDoc.DirId = dirIdNew
	err = db.DB.Set(db.GetDocumentTable(pid), docId, oldDoc)
	if err != nil {
		return err
	}

	err = db.DB.Set(db.GetDocumentDirItemTable(dirIdNew), docId, &entity.DocumentDirItem{
		DocId: docId,
		Sort:  0,
	})
	if err != nil {
		return err
	}

	err = db.DB.Delete(db.GetDocumentDirItemTable(oldDir), docId)
	if err != nil {
		return err
	}

	return nil
}

func (dao *DocDao) AddDocumentSnapshot(content *entity.DocumentContent, userAcc, logStr string) error {
	snapshotId := define.GetSnapshotId()

	data := &entity.DocumentSnapshot{
		SnapshotIdId:    snapshotId,
		UserAcc:         userAcc,
		Operation:       logStr,
		CreateTime:      time.Now().Unix(),
		DocumentContent: content,
	}

	err := db.DB.Set(db.GetDocumentSnapshotTable(content.DocId), snapshotId, data)
	if err != nil {
		return err
	}

	return nil
}

func (dao *DocDao) GetDocumentSnapshotList(docId string) ([]*entity.DocumentSnapshot, error) {

	result := make([]*entity.DocumentSnapshot, 0)

	err := db.DB.GetAll(db.GetDocumentSnapshotTable(docId), func(k, v []byte) {
		item := &entity.DocumentSnapshot{}
		err := json.Unmarshal(v, &item)
		if err != nil {
			log.Error(err)
		} else {
			result = append(result, item)
		}
	})

	return result, err
}

func (dao *DocDao) GetDocumentSnapshotItem(docId, snapshotId string) (*entity.DocumentSnapshot, error) {
	result := &entity.DocumentSnapshot{}
	err := db.DB.Get(db.GetDocumentSnapshotTable(docId), snapshotId, &result)
	return result, err
}
