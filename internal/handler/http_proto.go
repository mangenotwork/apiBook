package handler

type ProjectUsersResp struct {
	List []*UserInfo `json:"list"`
}

type UserInfo struct {
	Name       string `json:"name"`
	Account    string `json:"account"`
	CreateTime int64  `json:"createTime"` // 创建时间
	IsDisable  int    `json:"disable"`    // 是否被禁用 1:是
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
	IsAdmin   int    `json:"isAdmin"` // 1是
}

type AdminDeleteUserReq struct {
	Account string `json:"account"`
}

type AdminDisableUserReq struct {
	Account   string `json:"account"`
	IsDisable int    `json:"isDisable"`
}

type DocumentDirCreateReq struct {
	PId     string `json:"pid"`
	DirName string `json:"dirName"`
}

type DocumentDirDeleteReq struct {
	PId   string `json:"pid"`
	DirId string `json:"dirId"`
}

type DocumentDirModifyReq struct {
	PId     string `json:"pid"`
	DirId   string `json:"dirId"`
	DirName string `json:"dirName"`
}

type DocumentListReq struct {
	PId   string `json:"pid"`
	DirId string `json:"dirId"`
}

type DocumentItemParam struct {
	PId   string `json:"pid"`
	DocId string `json:"docId"`
}

type DocumentDeleteReq struct {
	PId   string `json:"pid"`
	DirId string `json:"dirId"`
	DocId string `json:"docId"`
}

type DocumentChangeDirReq struct {
	PId      string `json:"pid"`
	DirId    string `json:"dirId"`
	DirIdNew string `json:"dirIdNew"`
	DocId    string `json:"docId"`
}
