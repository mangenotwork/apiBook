package handler

type ProjectUsersResp struct {
	List []*UserInfo `json:"list"`
}

type UserInfo struct {
	Name       string `json:"name"`
	Account    string `json:"account"`
	CreateTime int64  `json:"create_time"` // 创建时间
	IsDisable  int    `json:"disable"`     // 是否被禁用 1:是
}

type ProjectAddUserReq struct {
	PId     string `json:"pid"`
	Account string `json:"account"`
}

type UserModifyReq struct {
	Name string `json:"name"`
}

type UserResetPasswordReq struct {
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

type AdminCreateUserReq struct {
	Name      string `json:"name"`
	Account   string `json:"account"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

type AdminDeleteUserReq struct {
	Account string `json:"account"`
}

type AdminDisableUserReq struct {
	Account   string `json:"account"`
	IsDisable int    `json:"is_disable"`
}

type DocumentDirCreateReq struct {
	PId     string `json:"pid"`
	DirName string `json:"dir_name"`
}
