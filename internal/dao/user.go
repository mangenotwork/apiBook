package dao

import (
	"apiBook/common/db"
	"apiBook/internal/entity"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (dao *UserDao) GetUserNum() int {
	stats, err := db.DB.Stats(db.UserTable)
	if err != nil {
		panic(err)
	}
	return stats.KeyN
}

func (dao *UserDao) Create(user *entity.User) error {
	return db.DB.Set(db.UserTable, user.Name, user)
}

func (dao *UserDao) Get(name string) (*entity.User, error) {
	user := &entity.User{}
	err := db.DB.Get(db.UserTable, name, &user)
	return user, err
}
