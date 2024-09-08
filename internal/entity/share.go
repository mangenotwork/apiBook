package entity

type Share struct {
	Key          string `json:"key"`          // 分享key
	ProjectId    string `json:"projectId"`    // 项目id
	ShareType    int64  `json:"shareType"`    // 1:项目  2:文档
	ShareId      string `json:"shareId"`      // 分享类型对应的值
	Expiration   int64  `json:"expiration"`   // -1:永久  单位:天
	IsPassword   int64  `json:"isPassword"`   // 0:否  1:是
	PasswordCode string `json:"passwordCode"` // 密码 - 自动生成
	CreateTime   int64  `json:"createTime"`   // 创建分享的时间
}
