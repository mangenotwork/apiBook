package dao

import (
	"apiBook/common/db"
	"apiBook/common/log"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"encoding/json"
	"fmt"
)

type DirDao struct {
}

func NewDirDao() *DirDao {
	return &DirDao{}
}

func (dao *DirDao) Get(pid, dirId string) (*entity.DocumentDir, error) {
	result := &entity.DocumentDir{}
	err := db.DB.Get(fmt.Sprintf(db.DocumentDirTable, pid), dirId, &result)
	return result, err
}

func (dao *DirDao) GetAll(pid string) ([]*entity.DocumentDir, error) {
	list := make([]*entity.DocumentDir, 0)
	allData, err := db.DB.GetAll(fmt.Sprintf(db.DocumentDirTable, pid))
	if err != nil {
		return list, err
	}

	for _, v := range allData {
		data := &entity.DocumentDir{}
		err = json.Unmarshal(v, &data)

		if err != nil {
			log.Error(err)
		} else {
			list = append(list, data)
		}
	}

	return list, nil
}

func (dao *DirDao) GetDirList(pid string) ([]string, error) {
	return db.DB.AllKey(fmt.Sprintf(db.DocumentDirTable, pid))
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

func (dao *DirDao) Create(pid string, data *entity.DocumentDir) error {
	has, err := dao.HasName(pid, data.DirName)
	if err != nil {
		return err
	}

	if !has {
		err = db.DB.Set(fmt.Sprintf(db.DocumentDirTable, pid), data.DirId, data)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("目录名已存在")
}

func (dao *DirDao) Delete(pid, dirId string) error {
	// 将dir 下的doc 移动到 默认
	docIdList, err := db.DB.AllKey(fmt.Sprintf(db.DocumentDirItemTable, dirId))
	if err != nil {
		return err
	}

	err = db.DB.Delete(fmt.Sprintf(db.DocumentDirTable, pid), dirId)
	if err != nil {
		return err
	}

	for _, v := range docIdList {
		// todo 待写入data
		_ = db.DB.Set(fmt.Sprintf(db.DocumentDirItemTable, define.GetDirDefault(pid)), v, 1)
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

		err = db.DB.Set(fmt.Sprintf(db.DocumentDirTable, pid), dirId, oldData)
		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("目录名已存在")
}

func (dao *DirDao) GetDoc(dirId, docId string) (*entity.DocumentDirItem, error) {
	result := &entity.DocumentDirItem{}
	err := db.DB.Get(fmt.Sprintf(db.DocumentDirItemTable, dirId), docId, &result)
	return result, err
}

func (dao *DirDao) GetDocList(dirId string) ([]*entity.DocumentDirItem, error) {
	result := make([]*entity.DocumentDirItem, 0)
	data, err := db.DB.GetAll(fmt.Sprintf(db.DocumentDirItemTable, dirId))
	if err != nil {
		return result, err
	}

	for _, v := range data {
		item := &entity.DocumentDirItem{}
		err = json.Unmarshal(v, &item)

		if err != nil {
			log.Error(err)
		} else {
			result = append(result, item)
		}
	}

	return result, nil
}

func (dao *DirDao) AddDoc(dirId string, dirItem *entity.DocumentDirItem) error {
	return db.DB.Set(fmt.Sprintf(db.DocumentDirItemTable, dirId), dirItem.DocId, dirItem)
}
