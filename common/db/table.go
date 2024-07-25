package db

const (
	UserTable     = "user_table"
	UserNameTable = "user_name_table" // key是用户名 值是用户id
)

var Tables = []string{
	UserTable,
}
