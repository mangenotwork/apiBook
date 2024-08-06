package entity

// DocumentDir 项目文档目录
// table 是 ProjectId
// key 是 DirId
type DocumentDir struct {
	DirId   string // 目录id
	DirName string // 目录名
	Sort    int    // 目录排序值
}

// DocumentDirItem 文档目录列表
// table 是 DirId
// key 是 DocId
type DocumentDirItem struct {
	DocId string // 文档id
	Sort  int    // 文档排序值
}