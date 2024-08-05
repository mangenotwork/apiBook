package handler

type ProjectUsersResp struct {
	List []*UserInfo `json:"list"`
}

type UserInfo struct {
	Name    string `json:"name"`
	Account string `json:"account"`
}
