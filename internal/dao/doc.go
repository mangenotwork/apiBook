package dao

import "apiBook/internal/entity"

type DocDao struct {
}

func NewDocDao() *DocDao {
	return &DocDao{}
}

func (dao *DocDao) Create(data *entity.Document, userAcc string) error {
	return nil
}
