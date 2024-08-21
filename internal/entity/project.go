package entity

import "apiBook/internal/define"

type Project struct {
	ProjectId     string                    `json:"projectId"`
	Name          string                    `json:"name"`        // 项目名
	Description   string                    `json:"description"` // 项目简述
	CreateUserAcc string                    `json:"createUserAcc"`
	CreateDate    string                    `json:"createDate"` // 创建时间
	Private       define.ProjectPrivateCode `json:"private"`    // 1:公有(所有人可见)  2:私有
}

type ProjectProperty struct {
	ProjectId string `json:"projectId"`
	UserId    string `json:"userId"`
	Property  string `json:"Property"` // 第1位查看权限  第2位编辑权限  第3位删除权限
}

type ProjectGlobalHeader struct {
	ProjectId string           `json:"projectId"`
	ReqHeader []*ReqHeaderItem `json:"reqHeader"` // 请求头
}

type ProjectGlobalCode struct {
	ProjectId string            `json:"projectId"`
	List      []*GlobalCodeItem `json:"list"`
}

type GlobalCodeItem struct {
	CodeValue   string `json:"field"`       // 状态值
	Description string `json:"description"` // 描述
}
