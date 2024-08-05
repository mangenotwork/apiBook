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

func (dao *ProjectDao) Create(data *entity.Project, userAcc string) error {

	if data.Private == define.ProjectPrivate {
		err := db.DB.Set(fmt.Sprintf(db.UserPrivateProjectTable, userAcc), data.ProjectId, 1)
		if err != nil {
			return err
		}

		err = db.DB.Set(fmt.Sprintf(db.ProjectPrivateUserTable, data.ProjectId), userAcc, 1)
		if err != nil {
			return err
		}

	} else {
		err := db.DB.Set(db.ProjectPublicTable, data.ProjectId, 1)
		if err != nil {
			return err
		}
	}

	return db.DB.Set(db.ProjectTable, data.ProjectId, data)
}

func (dao *ProjectDao) GetList(userAcc string) []*entity.Project {
	resp := make([]*entity.Project, 0)

	pubKeyList, err := db.DB.AllKey(db.ProjectPublicTable)
	if err != nil {
		log.Info(err)
	}

	uKeyList, err := db.DB.AllKey(fmt.Sprintf(db.UserPrivateProjectTable, userAcc))
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
		_ = db.DB.Get(fmt.Sprintf(db.UserPrivateProjectTable, userAcc), pid, &has)
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
			err = db.DB.Set(fmt.Sprintf(db.UserPrivateProjectTable, userAcc), oldData.ProjectId, 1)
			if err != nil {
				return err
			}

			err = db.DB.Set(fmt.Sprintf(db.ProjectPrivateUserTable, oldData.ProjectId), userAcc, 1)
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
		err = db.DB.Delete(fmt.Sprintf(db.UserPrivateProjectTable, userAcc), data.ProjectId)
		if err != nil {
			log.Error(err)
		}

		err = db.DB.Delete(fmt.Sprintf(db.ProjectPrivateUserTable, data.ProjectId), userAcc)
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

	resp, err = db.DB.AllKey(fmt.Sprintf(db.ProjectPrivateUserTable, data.ProjectId))
	if err != nil {
		log.Error(err)
		return resp, err
	}

	return resp, nil
}
