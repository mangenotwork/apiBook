package dao

import (
	"apiBook/common/db"
	"apiBook/common/log"
	"apiBook/internal/entity"
	"time"
)

type ShareDao struct {
}

func NewShareDao() *ShareDao {
	return &ShareDao{}
}

func (dao *ShareDao) Create(data *entity.Share) error {

	data.CreateTime = time.Now().Unix()

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

func (dao *ShareDao) GetShareProjectList(pid string) ([]*entity.Share, error) {
	result := make([]*entity.Share, 0)

	allKey, err := db.DB.AllKey(db.GetShareProjectTable(pid))
	if err != nil {
		return result, err
	}

	for _, v := range allKey {
		item, err := dao.GetInfo(v)
		if err != nil {
			log.Error(err)
		} else {
			result = append(result, item)
		}

	}

	return result, err
}

func (dao *ShareDao) GetShareDocumentList(docId string) ([]*entity.Share, error) {
	result := make([]*entity.Share, 0)
	allKey, err := db.DB.AllKey(db.GetShareDocumentTable(docId))
	if err != nil {
		return result, err
	}

	for _, v := range allKey {
		item, err := dao.GetInfo(v)
		if err != nil {
			log.Error(err)
		} else {
			result = append(result, item)
		}
	}

	return result, err
}

func (dao *ShareDao) GetInfo(key string) (*entity.Share, error) {
	result := &entity.Share{}
	err := db.DB.Get(db.GetShareDataTable(), key, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (dao *ShareDao) Del(key string) (*entity.Share, error) {
	info, err := dao.GetInfo(key)
	if err != nil {
		return nil, err
	}

	switch info.ShareType {
	case 1:
		_ = db.DB.Delete(db.GetShareProjectTable(info.ShareId), key)
	case 2:
		_ = db.DB.Delete(db.GetShareDocumentTable(info.ShareId), key)
	}

	err = db.DB.Delete(db.GetShareDataTable(), key)
	if err != nil {
		return info, err
	}

	return info, nil
}
