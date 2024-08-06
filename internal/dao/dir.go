package dao

import (
	"apiBook/common/db"
	"apiBook/common/log"
	"apiBook/internal/entity"
	"encoding/json"
	"fmt"
)

type DirDao struct {
}

func NewDirDao() *DirDao {
	return &DirDao{}
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
	}

	return nil
}

func (dao *DirDao) Delete(pid, dirId string) {
	// todo...
}
