package dao

import "apiBook/internal/entity"

type DirDao struct {
}

func NewDirDao() *DirDao {
	return &DirDao{}
}

func (dao *DirDao) Create(data *entity.DocumentDir, userAcc string) error {
	return nil
}
