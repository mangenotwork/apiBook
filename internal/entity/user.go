package entity

type User struct {
	Account    string `json:"account"` // 唯一不变 key
	Name       string `json:"name"`
	Password   string `json:"password"`
	IsAdmin    int    `json:"isAdmin"`    // 是否是超级管理员  1:是
	CreateTime int64  `json:"createTime"` // 创建时间
	IsDisable  int    `json:"disable"`    // 是否被禁用 1:是
}
