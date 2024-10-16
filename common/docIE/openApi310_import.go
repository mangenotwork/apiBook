package docIE

import (
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"encoding/json"
	"fmt"
)

/*

open-api 3.1.* 导入

https://spec.openapis.org/oas/v3.1.0


*/

type OpenApi310Import struct {
	Openapi    string                 `json:"openapi"`
	Info       map[string]interface{} `json:"info"`
	Tags       []interface{}          `json:"tags"`
	Paths      map[string]interface{} `json:"paths"`
	Components map[string]interface{} `json:"components"`
	Servers    []interface{}          `json:"servers"`
}

func NewOpenApi310Import() *OpenApi310Import {
	return &OpenApi310Import{}
}

func (obj *OpenApi310Import) Whole(text, userAcc string, private define.ProjectPrivateCode) error {
	return nil
}

func (obj *OpenApi310Import) Increment(text, pid, userAcc, dirId string) error {
	return nil
}

func (obj *OpenApi310Import) analysis(text string) error {
	err := json.Unmarshal([]byte(text), &obj)
	if err != nil {
		log.Error(err)
		return err
	}

	if obj.Openapi != "3.1.0" {
		return fmt.Errorf("openapi is not 3.1.0")
	}

	return nil
}

func (obj *OpenApi310Import) analysisProject(userAcc string, private define.ProjectPrivateCode) *entity.Project {
	return &entity.Project{
		ProjectId:     utils.IDMd5(),
		Name:          utils.AnyToString(obj.Info["title"]),
		Description:   fmt.Sprintf("version: %s; %s", utils.AnyToString(obj.Info["version"]), utils.AnyToString(obj.Info["description"])),
		CreateUserAcc: userAcc,
		CreateDate:    utils.NowDate(),
		Private:       private,
	}
}

func (obj *OpenApi310Import) analysisDoc(project *entity.Project, userAcc string, f func(doc *entity.DocumentContent)) {
}
