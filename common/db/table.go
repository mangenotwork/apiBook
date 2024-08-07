package db

import "fmt"

const (
	UserTable               = "user"       // 用户表  key是用户acc
	ProjectTable            = "project"    // 项目表  key是项目id
	UserPrivateProjectTable = "upp:%s"     // 用户项目列表： 查用户能访问的私有项目列表, %s是用户账号acc， key是项目id
	ProjectPrivateUserTable = "ppu:%s"     // 项目权限表: 查该项目哪个用户能访问, %s是项目id, key是用户账户acc
	ProjectPublicTable      = "ppub"       // 公有项目 key是项目id
	DocumentDirTable        = "dir:%s"     // 项目文档目录, %s是项目id(ProjectId), key是目录id(DirId)
	DocumentDirItemTable    = "diritem:%s" // 文档目录列表, %s是目录id(DirId)， key是文档id(DocId)
	DocumentTable           = "doc:%s"     // 项目文档， %s是项目id(ProjectId), key是文档id(DocId)
	DocumentContentTable    = "docd:%s"    // 文档数据， %s是项目id(ProjectId),  Key是文档id(DocId)
	ProjectNameTable        = "pname"      // 项目名字表用于检查重名, Key是名称   key是项目id
)

var Tables = []string{
	UserTable, ProjectTable, ProjectPublicTable, ProjectNameTable,
}

// GetUserTable 用户表  key是用户acc
func GetUserTable() string {
	return UserTable
}

// GetProjectTable 项目表  key是项目id
func GetProjectTable() string {
	return ProjectTable
}

// GetUserPrivateProjectTable 用户项目列表： 查用户能访问的私有项目列表, %s是用户账号， key是项目id
func GetUserPrivateProjectTable(acc string) string {
	return fmt.Sprintf(UserPrivateProjectTable, acc)
}

// GetProjectPrivateUserTable 项目权限表: 查该项目哪个用户能访问, %s是项目id, key是用户账户
func GetProjectPrivateUserTable(pid string) string {
	return fmt.Sprintf(ProjectPrivateUserTable, pid)
}

// GetProjectPublicTable 公有项目 key是项目id
func GetProjectPublicTable() string {
	return ProjectPublicTable
}

// GetDocumentDirTable 项目文档目录, %s是项目id(ProjectId), key是目录id(DirId)
func GetDocumentDirTable(pid string) string {
	return fmt.Sprintf(DocumentDirTable, pid)
}

// GetDocumentDirItemTable 文档目录列表, %s是目录id(DirId)， key是文档id(DocId)
func GetDocumentDirItemTable(dirId string) string {
	return fmt.Sprintf(DocumentDirItemTable, dirId)
}

// GetDocumentTable 项目文档， %s是项目id(ProjectId), key是文档id(DocId)
func GetDocumentTable(pid string) string {
	return fmt.Sprintf(DocumentTable, pid)
}

// GetDocumentContentTable 文档数据， %s是项目id(ProjectId),  Key是文档id(DocId)
func GetDocumentContentTable(pid string) string {
	return fmt.Sprintf(DocumentContentTable, pid)
}

// GetProjectNameTable  项目名字表用于检查重名, Key是名称   key是项目id
func GetProjectNameTable() string {
	return ProjectNameTable
}
