package entity

type User struct {
	UserId   string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	IsAdmin  int    `json:"isAdmin"` // 是否是管理员  1:是
}
