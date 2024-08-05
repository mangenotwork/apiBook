package dao

import (
	"apiBook/common/db"
	"apiBook/common/log"
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
	return db.DB.Set(db.UserTable, user.Account, user)
}

func (dao *UserDao) Get(account string) (*entity.User, error) {
	user := &entity.User{}
	err := db.DB.Get(db.UserTable, account, &user)
	return user, err
}

func (dao *UserDao) GetUsers(list []string) ([]*entity.User, error) {
	data := make([]*entity.User, 0)
	for _, v := range list {
		u, e := dao.Get(v)
		if e != nil {
			log.Error(e)
			continue
		}
		data = append(data, u)
	}
	return data, nil
}
