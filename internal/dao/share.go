package dao

import (
	"apiBook/common/db"
	"apiBook/internal/entity"
)

type ShareDao struct {
}

func NewShareDao() *ShareDao {
	return &ShareDao{}
}

func (dao *ShareDao) Create(data *entity.Share) error {
	err := db.DB.Set(db.GetShareDataTable(), data.Key, data)
	if err != nil {
		return err
	}

	switch data.ShareType {
	case 1:
		_ = db.DB.Set(db.GetShareProjectTable(data.ShareId), data.Key, 1)
	case 2:
		_ = db.DB.Set(db.GetShareDocumentTable(data.ShareId), data.Key, 1)
	}

	return nil
}

func (dao *ShareDao) GetShareProjectList(pid string) ([]string, error) {
	return db.DB.AllKey(db.GetShareProjectTable(pid))
}

func (dao *ShareDao) GetShareDocumentList(docId string) ([]string, error) {
	return db.DB.AllKey(db.GetShareDocumentTable(docId))
}

func (dao *ShareDao) GetInfo(key string) (*entity.Share, error) {
	result := &entity.Share{}
	err := db.DB.Get(db.GetShareDataTable(), key, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
