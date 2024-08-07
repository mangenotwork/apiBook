package dao

import (
	"apiBook/common/db"
	"apiBook/common/log"
	"apiBook/internal/entity"
	"fmt"
)

type DocDao struct {
}

func NewDocDao() *DocDao {
	return &DocDao{}
}

func (dao *DocDao) Create(data *entity.Document, content *entity.DocumentContent) error {

	err := db.DB.Set(fmt.Sprintf(db.DocumentTable, data.ProjectId), data.DocId, data)
	if err != nil {
		return err
	}

	err = db.DB.Set(fmt.Sprintf(db.DocumentContentTable, content.ProjectId), content.DocId, content)
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
