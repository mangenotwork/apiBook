package db

const (
	UserTable               = "user"
	ProjectTable            = "project"
	UserPrivateProjectTable = "upp:%s"     // 用户项目列表： 查用户能访问的私有项目列表, %s是用户账号， key是项目id
	ProjectPrivateUserTable = "ppu:%s"     // 项目权限表: 查该项目哪个用户能访问, %s是项目, key是用户账户
	ProjectPublicTable      = "ppub"       // 公有项目 key是项目id
	DocumentDirTable        = "dir:%s"     // 项目文档目录, %s是项目id(ProjectId), key是目录id(DirId)
	DocumentDirItemTable    = "diritem:%s" // 文档目录列表, %s是目录id(DirId)， key是文档id(DocId)
)

var Tables = []string{
	UserTable, ProjectTable, ProjectPublicTable,
}
