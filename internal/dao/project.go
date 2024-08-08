package dao

import (
	"apiBook/common/db"
	"apiBook/common/log"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"fmt"
	"sort"
)

type ProjectDao struct {
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{}
}

func (dao *ProjectDao) HasName(name string) bool {
	var result int = 0
	_ = db.DB.Get(db.GetProjectNameTable(), name, &result)
	if result > 0 {
		return true
	}
	return false
}

func (dao *ProjectDao) Create(data *entity.Project, userAcc string) error {

	if dao.HasName(data.Name) {
		return fmt.Errorf("项目名已存在")
	}

	if data.Private == define.ProjectPrivate {
		err := db.DB.Set(db.GetUserPrivateProjectTable(userAcc), data.ProjectId, 1)
		if err != nil {
			return err
		}

		err = db.DB.Set(db.GetProjectPrivateUserTable(data.ProjectId), userAcc, 1)
		if err != nil {
			return err
		}

	} else {
		err := db.DB.Set(db.GetProjectPublicTable(), data.ProjectId, 1)
		if err != nil {
			return err
		}
	}

	err := db.DB.Set(db.GetProjectTable(), data.ProjectId, data)
	if err != nil {
		return err
	}

	return db.DB.Set(db.GetProjectNameTable(), data.Name, 1)
}

func (dao *ProjectDao) GetList(userAcc string) []*entity.Project {
	resp := make([]*entity.Project, 0)

	pubKeyList, err := db.DB.AllKey(db.ProjectPublicTable)
	if err != nil {
		log.Info(err)
	}

	uKeyList, err := db.DB.AllKey(db.GetUserPrivateProjectTable(userAcc))
	if err != nil {
		log.Info(err)
	}

	pubKeyList = append(pubKeyList, uKeyList...)

	for _, v := range pubKeyList {
		projectData := &entity.Project{}
		_ = db.DB.Get(db.ProjectTable, v, &projectData)
		resp = append(resp, projectData)
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].CreateDate < resp[j].CreateDate
	})

	return resp
}

func (dao *ProjectDao) Get(pid, userAcc string) (*entity.Project, error) {
	projectData := &entity.Project{}
	err := db.DB.Get(db.ProjectTable, pid, &projectData)
	if err != nil {
		return projectData, err
	}

	if projectData.Private == define.ProjectPrivate {
		var has int = 0
		_ = db.DB.Get(db.GetUserPrivateProjectTable(userAcc), pid, &has)
		if has == 0 {
			return projectData, fmt.Errorf("没有权限")
		}
	}

	return projectData, nil
}

func (dao *ProjectDao) Modify(newData *entity.Project, userAcc string) error {
	oldData, err := dao.Get(newData.ProjectId, userAcc)
	if err != nil {
		log.Error(err)
		return err
	}

	if oldData.CreateUserAcc != userAcc {
		return fmt.Errorf("您不是项目创建者，无权修改")
	}

	oldData.Name = newData.Name
	oldData.Description = newData.Description

	if oldData.Private != newData.Private {
		if newData.Private == define.ProjectPrivate {
			err = db.DB.Set(db.GetUserPrivateProjectTable(userAcc), oldData.ProjectId, 1)
			if err != nil {
				return err
			}

			err = db.DB.Set(db.GetProjectPrivateUserTable(oldData.ProjectId), userAcc, 1)
			if err != nil {
				return err
			}

		} else {
			err = db.DB.Set(db.ProjectPublicTable, oldData.ProjectId, 1)
			if err != nil {
				return err
			}
		}
		oldData.Private = newData.Private
	}

	return db.DB.Set(db.ProjectTable, oldData.ProjectId, oldData)
}

func (dao *ProjectDao) Delete(pid, userAcc string) error {
	data, err := dao.Get(pid, userAcc)
	if err != nil {
		log.Error(err)
		return err
	}

	if data.CreateUserAcc != userAcc {
		return fmt.Errorf("您不是项目创建者，无权修改")
	}

	if data.Private == define.ProjectPrivate {
		err = db.DB.Delete(db.GetUserPrivateProjectTable(userAcc), data.ProjectId)
		if err != nil {
			log.Error(err)
		}

		err = db.DB.Delete(db.GetProjectPrivateUserTable(data.ProjectId), userAcc)
		if err != nil {
			log.Error(err)
		}

	} else {
		err = db.DB.Delete(db.ProjectPublicTable, data.ProjectId)
		if err != nil {
			log.Error(err)
		}
	}

	return db.DB.Delete(db.ProjectTable, data.ProjectId)
}

func (dao *ProjectDao) GetUserList(pid, userAcc string) ([]string, error) {
	resp := make([]string, 0)

	data, err := dao.Get(pid, userAcc)
	if err != nil {
		log.Error(err)
		return resp, err
	}

	resp, err = db.DB.AllKey(db.GetProjectPrivateUserTable(data.ProjectId))
	if err != nil {
		log.Error(err)
		return resp, err
	}

	return resp, nil
}

func (dao *ProjectDao) AddUser(pid, userAcc, addAcc string) error {
	data, err := dao.Get(pid, userAcc)
	if err != nil {
		log.Error(err)
		return err
	}

	if data.Private == define.ProjectPublic {
		return fmt.Errorf("公有项目无需添加协助者")
	}

	err = db.DB.Set(db.GetUserPrivateProjectTable(addAcc), data.ProjectId, 1)
	if err != nil {
		return err
	}

	err = db.DB.Set(db.GetProjectPrivateUserTable(data.ProjectId), addAcc, 1)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ProjectDao) DelUser(pid, userAcc, delAcc string) error {
	data, err := dao.Get(pid, userAcc)
	if err != nil {
		log.Error(err)
		return err
	}

	if data.CreateUserAcc == delAcc {
		return fmt.Errorf("不能移除项目创建者")
	}

	err = db.DB.Delete(db.GetUserPrivateProjectTable(delAcc), data.ProjectId)
	if err != nil {
		log.Error(err)
	}

	err = db.DB.Delete(db.GetProjectPrivateUserTable(data.ProjectId), delAcc)
	if err != nil {
		log.Error(err)
	}

	return nil
}
