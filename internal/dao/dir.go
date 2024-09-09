package dao

import (
	"apiBook/common/db"
	"apiBook/common/log"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"encoding/json"
	"errors"
)

type DirDao struct {
}

func NewDirDao() *DirDao {
	return &DirDao{}
}

func (dao *DirDao) Get(pid, dirId string) (*entity.DocumentDir, error) {
	result := &entity.DocumentDir{}
	err := db.DB.Get(db.GetDocumentDirTable(pid), dirId, &result)
	return result, err
}

func (dao *DirDao) GetAll(pid string) ([]*entity.DocumentDir, error) {
	list := make([]*entity.DocumentDir, 0)

	err := db.DB.GetAll(db.GetDocumentDirTable(pid), func(k, v []byte) {
		data := &entity.DocumentDir{}
		err := json.Unmarshal(v, &data)

		if err != nil {
			log.Error(err)
		} else {
			list = append(list, data)
		}
	})

	if err != nil {
		return list, err
	}

	return list, nil
}

func (dao *DirDao) GetDirList(pid string) ([]string, error) {
	return db.DB.AllKey(db.GetDocumentDirTable(pid))
}

func (dao *DirDao) GetDirNum(pid string) int {
	list, err := dao.GetDirList(pid)
	if err != nil {
		return 0
	}

	return len(list)
}

func (dao *DirDao) HasName(pid, name string) (bool, error) {
	allData, err := dao.GetAll(pid)
	if err != nil {
		log.Error(err)
		return true, err
	}

	for _, v := range allData {
		if v.DirName == name {
			return true, nil
		}
	}

	return false, nil
}

func (dao *DirDao) CreateInit(pid string) error {
	dirDef := &entity.DocumentDir{
		DirId:   define.GetDirDefault(pid),
		DirName: define.DirNameDefault,
		Sort:    1,
	}

	dirRecycleBin := &entity.DocumentDir{
		DirId:   define.GetDirRecycleBinKey(pid),
		DirName: define.DirNameRecycleBin,
		Sort:    2,
	}

	err := db.DB.Set(db.GetDocumentDirTable(pid), dirDef.DirId, dirDef)
	if err != nil {
		return err
	}

	err = db.DB.Set(db.GetDocumentDirTable(pid), dirRecycleBin.DirId, dirRecycleBin)
	if err != nil {
		return err
	}

	return nil
}

func (dao *DirDao) Create(pid string, data *entity.DocumentDir) error {
	has, err := dao.HasName(pid, data.DirName)
	if err != nil && !errors.Is(err, db.TableNotFound) {
		return err
	}

	if !has {
		err = db.DB.Set(db.GetDocumentDirTable(pid), data.DirId, data)
		if err != nil {
			return err
		}
		return nil
	}

	return define.DirNameExistErr
}

func (dao *DirDao) Delete(pid, dirId string) error {
	docIdList, err := db.DB.AllKey(db.GetDocumentDirItemTable(dirId))
	if err != nil {
		return err
	}

	err = db.DB.Delete(db.GetDocumentDirTable(pid), dirId)
	if err != nil {
		return err
	}

	for _, v := range docIdList {
		_ = db.DB.Set(db.GetDocumentDirItemTable(define.GetDirDefault(pid)), v, 1)
	}

	return nil
}

func (dao *DirDao) Modify(pid, dirId, dirName string) error {
	has, err := dao.HasName(pid, dirName)
	if err != nil {
		return err
	}

	if !has {
		oldData, err := dao.Get(pid, dirId)
		if err != nil {
			return err
		}

		oldData.DirName = dirName

		err = db.DB.Set(db.GetDocumentDirTable(pid), dirId, oldData)
		if err != nil {
			return err
		}

		return nil
	}

	return define.DirNameExistErr
}

func (dao *DirDao) GetDoc(dirId, docId string) (*entity.DocumentDirItem, error) {
	result := &entity.DocumentDirItem{}
	err := db.DB.Get(db.GetDocumentDirItemTable(dirId), docId, &result)
	return result, err
}

func (dao *DirDao) GetDocList(dirId string) ([]*entity.DocumentDirItem, error) {
	result := make([]*entity.DocumentDirItem, 0)

	err := db.DB.GetAll(db.GetDocumentDirItemTable(dirId), func(k, v []byte) {
		item := &entity.DocumentDirItem{}
		err := json.Unmarshal(v, &item)

		if err != nil {
			log.Error(err)
		} else {
			result = append(result, item)
		}

	})

	if err != nil {
		return result, err
	}

	return result, nil
}

func (dao *DirDao) AddDoc(dirId string, dirItem *entity.DocumentDirItem) error {
	return db.DB.Set(db.GetDocumentDirItemTable(dirId), dirItem.DocId, dirItem)
}
