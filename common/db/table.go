package db

import "fmt"

const (
	UserTable               = "user"        // 用户表  key是用户acc  值是 entity.User
	ProjectTable            = "project"     // 项目表  key是项目id  值是 entity.Project
	UserPrivateProjectTable = "upp:%s"      // 用户项目列表： 查用户能访问的私有项目列表, %s是用户账号acc， key是项目id  值是 int64
	ProjectPrivateUserTable = "ppu:%s"      // 项目权限表: 查该项目哪个用户能访问, %s是项目id, key是用户账户acc  值是 int64
	ProjectPublicTable      = "ppub"        // 公有项目 key是项目id  值是 int64
	DocumentDirTable        = "dir:%s"      // 项目文档目录, %s是项目id(ProjectId), key是目录id(DirId)  值是 entity.DocumentDir
	DocumentDirItemTable    = "diritem:%s"  // 文档目录列表, %s是目录id(DirId)， key是文档id(DocId)  值是 int64
	DocumentTable           = "doc:%s"      // 项目文档， %s是项目id(ProjectId), key是文档id(DocId)  值是 entity.Document
	DocumentContentTable    = "docd:%s"     // 文档数据， %s是项目id(ProjectId),  Key是文档id(DocId)  值是 entity.DocumentContent
	ProjectNameTable        = "pname"       // 项目名字表用于检查重名, Key是名称  值是 int64   key是项目id
	DocumentSnapshotTable   = "snapshot:%s" // 文档快照, %s是文档id(DocId) Key是快照id  值是 entity.DocumentSnapshot
	GlobalHeader            = "gheader"     // 全局header, Key是项目id(ProjectId)  值是 entity.ProjectGlobalHeader
	GlobalCode              = "gcode"       // 全局code, Key是项目id(ProjectId)  值是 entity.ProjectGlobalCode
	ShareProjectTable       = "sp:%s"       // 项目的分享信息, %s是项目id(ProjectId), Key是分享Key  值是 int64
	ShareDocumentTable      = "sd:%s"       // 文档分享信息, %是文档id(DocId), Key是分享Key  值是 int64
	ShareDataTable          = "share"       // 分享数据, Key是分享Key  值是 entity.Share
)

var Tables = []string{
	UserTable, ProjectTable, ProjectPublicTable, ProjectNameTable, GlobalHeader, GlobalCode, ShareDataTable,
}

// GetUserTable 用户表  key是用户acc  值是 entity.User
func GetUserTable() string {
	return UserTable
}

// GetProjectTable 项目表  key是项目id  值是 entity.Project
func GetProjectTable() string {
	return ProjectTable
}

// GetUserPrivateProjectTable 用户项目列表： 查用户能访问的私有项目列表, %s是用户账号， key是项目id  值是 int64
func GetUserPrivateProjectTable(acc string) string {
	return fmt.Sprintf(UserPrivateProjectTable, acc)
}

// GetProjectPrivateUserTable 项目权限表: 查该项目哪个用户能访问, %s是项目id, key是用户账户  值是 int64
func GetProjectPrivateUserTable(pid string) string {
	return fmt.Sprintf(ProjectPrivateUserTable, pid)
}

// GetProjectPublicTable 公有项目 key是项目id  值是 int64
func GetProjectPublicTable() string {
	return ProjectPublicTable
}

// GetDocumentDirTable 项目文档目录, %s是项目id(ProjectId), key是目录id(DirId)  值是 entity.DocumentDir
func GetDocumentDirTable(pid string) string {
	return fmt.Sprintf(DocumentDirTable, pid)
}

// GetDocumentDirItemTable 文档目录列表, %s是目录id(DirId)， key是文档id(DocId)  值是 entity.DocumentDirItem
func GetDocumentDirItemTable(dirId string) string {
	return fmt.Sprintf(DocumentDirItemTable, dirId)
}

// GetDocumentTable 项目文档， %s是项目id(ProjectId), key是文档id(DocId)  值是 entity.Document
func GetDocumentTable(pid string) string {
	return fmt.Sprintf(DocumentTable, pid)
}

// GetDocumentContentTable 文档数据， %s是项目id(ProjectId),  Key是文档id(DocId)  值是 entity.DocumentContent
func GetDocumentContentTable(pid string) string {
	return fmt.Sprintf(DocumentContentTable, pid)
}

// GetProjectNameTable  项目名字表用于检查重名, Key是名称   key是项目id  值是 int64
func GetProjectNameTable() string {
	return ProjectNameTable
}

// GetDocumentSnapshotTable 文档快照, %s是文档id(DocId) Key是快照id  值是 entity.DocumentSnapshot
func GetDocumentSnapshotTable(docId string) string {
	return fmt.Sprintf(DocumentSnapshotTable, docId)
}

// GetGlobalHeader 全局header, Key是项目id(ProjectId)  值是 entity.ProjectGlobalHeader
func GetGlobalHeader() string {
	return GlobalHeader
}

// GetGlobalCode 全局code, Key是项目id(ProjectId)  值是 entity.ProjectGlobalCode
func GetGlobalCode() string {
	return GlobalCode
}

// GetShareProjectTable 项目的分享信息, %s是项目id(ProjectId), Key是分享Key  值是 int64
func GetShareProjectTable(pid string) string {
	return fmt.Sprintf(ShareProjectTable, pid)
}

// GetShareDocumentTable 文档分享信息, %是文档id(DocId), Key是分享Key  值是 int64
func GetShareDocumentTable(docId string) string {
	return fmt.Sprintf(ShareDocumentTable, docId)
}

// GetShareDataTable 分享数据, Key是分享Key  值是 entity.Share
func GetShareDataTable() string {
	return ShareDataTable
}
