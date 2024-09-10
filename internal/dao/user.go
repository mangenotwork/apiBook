package dao

import (
	"apiBook/common/db"
	"apiBook/common/log"
	"apiBook/internal/entity"
	"encoding/json"
	"sort"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (dao *UserDao) GetUserNum() int {
	stats, err := db.DB.Stats(db.GetUserTable())
	if err != nil {
		log.Error(err)
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

func (dao *UserDao) Modify(account, name string) error {
	user, err := dao.Get(account)
	if err != nil {
		return err
	}

	user.Name = name

	return db.DB.Set(db.UserTable, account, user)
}

func (dao *UserDao) ResetPassword(account, password string) error {
	user, err := dao.Get(account)
	if err != nil {
		return err
	}

	user.Password = password

	return db.DB.Set(db.UserTable, account, user)
}

func (dao *UserDao) GetAllUser() []*entity.User {
	list := make([]*entity.User, 0)

	err := db.DB.GetAll(db.UserTable, func(k, v []byte) {

		item := &entity.User{}
		err := json.Unmarshal(v, item)
		if err != nil {
			return
		}
		list = append(list, item)
	})

	if err != nil {
		log.Error(err)
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].CreateTime > list[j].CreateTime {
			return false
		}
		return true
	})

	return list
}

func (dao *UserDao) IsAdmin(acc string) bool {
	user, err := dao.Get(acc)
	if err == nil && user.IsAdmin == 1 {
		return true
	}

	return false
}

func (dao *UserDao) DelUser(acc string) error {
	return db.DB.Delete(db.UserTable, acc)
}

func (dao *UserDao) DisableUser(acc string, isDisable int) error {
	user, err := dao.Get(acc)
	if err != nil {
		return err
	}

	user.IsDisable = isDisable

	return db.DB.Set(db.UserTable, acc, user)
}

func (dao *UserDao) HasUserAccount(acc string) bool {
	userInfo, _ := dao.Get(acc)
	if userInfo.Account != "" {
		return true
	}
	return false
}

func (dao *UserDao) HasUserName(name string) (bool, error) {
	has := false
	err := db.DB.GetAll(db.UserTable, func(k, v []byte) {

		item := &entity.User{}
		err := json.Unmarshal(v, item)
		if err != nil {
			return
		}

		if name == item.Name {
			has = true
		}
	})

	return has, err
}
