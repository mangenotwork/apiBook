package db

import "fmt"

const (
	UserTable               = "user"
	ProjectTable            = "project"
	UserPrivateProjectTable = "upp:%s"     // 用户项目列表： 查用户能访问的私有项目列表, %s是用户账号， key是项目id
	ProjectPrivateUserTable = "ppu:%s"     // 项目权限表: 查该项目哪个用户能访问, %s是项目id, key是用户账户
	ProjectPublicTable      = "ppub"       // 公有项目 key是项目id
	DocumentDirTable        = "dir:%s"     // 项目文档目录, %s是项目id(ProjectId), key是目录id(DirId)
	DocumentDirItemTable    = "diritem:%s" // 文档目录列表, %s是目录id(DirId)， key是文档id(DocId)
	DocumentTable           = "doc:%s"     // 项目文档， %s是项目id(ProjectId), key是文档id(DocId)
	DocumentContentTable    = "docd:%s"    // 文档数据， %s是项目id(ProjectId),  Key是文档id(DocId)
)

var Tables = []string{
	UserTable, ProjectTable, ProjectPublicTable,
}

func GetUserTable() string {
	return UserTable
}

func GetProjectTable() string {
	return ProjectTable
}

func GetUserPrivateProjectTable(acc string) string {
	return fmt.Sprintf(UserPrivateProjectTable, acc)
}

func GetProjectPrivateUserTable(pid string) string {
	return fmt.Sprintf(ProjectPrivateUserTable, pid)
}

func GetProjectPublicTable() string {
	return ProjectPublicTable
}

func GetDocumentDirTable(pid string) string {
	return fmt.Sprintf(DocumentDirTable, pid)
}

func GetDocumentDirItemTable(dirId string) string {
	return fmt.Sprintf(DocumentDirItemTable, dirId)
}

func GetDocumentTable(pid string) string {
	return fmt.Sprintf(DocumentTable, pid)
}

func GetDocumentContentTable(pid string) string {
	return fmt.Sprintf(DocumentContentTable, pid)
}
