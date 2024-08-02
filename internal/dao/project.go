package dao

import (
	"apiBook/common/db"
	"apiBook/internal/entity"
	"fmt"
)

type ProjectDao struct {
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{}
}

func (dao *ProjectDao) Create(data *entity.Project, userAcc string) error {

	if data.Private == 1 {
		err := db.DB.Set(fmt.Sprintf(db.UserPrivateProjectTable, userAcc), data.ProjectId, 1)
		if err != nil {
			return err
		}

		err = db.DB.Set(fmt.Sprintf(db.ProjectPrivateUserTable, data.ProjectId), userAcc, 1)
		if err != nil {
			return err
		}

	}

	return db.DB.Set(db.ProjectTable, data.ProjectId, data)
}
