package cache

import "time"

// 只对指定业务做缓存
var (
	DocumentDirTableCache        = NewCache(24*7*time.Hour, 24*1*time.Hour) // 项目文档目录, %s是项目id(ProjectId), key是目录id(DirId)  值是 entity.DocumentDir
	DocumentDirItemTableCache    = NewCache(24*7*time.Hour, 24*1*time.Hour) // 文档目录列表, %s是目录id(DirId)， key是文档id(DocId)  值是 int64
	DocumentTableCache           = NewCache(24*7*time.Hour, 24*1*time.Hour) // 项目文档， %s是项目id(ProjectId), key是文档id(DocId)  值是 entity.Document
	DocumentContentTableCache    = NewCache(24*7*time.Hour, 24*1*time.Hour) // 文档数据， %s是项目id(ProjectId),  Key是文档id(DocId)  值是 entity.DocumentContent
	ProjectNameTableCache        = NewCache(24*7*time.Hour, 24*1*time.Hour) // 项目名字表用于检查重名, Key是名称  值是 int64   key是项目id
	DocumentSnapshotTableCache   = NewCache(24*7*time.Hour, 24*1*time.Hour) // 文档快照, %s是文档id(DocId) Key是快照id  值是 entity.DocumentSnapshot
	InvertIndexTableCache        = NewCache(24*7*time.Hour, 24*1*time.Hour) // 倒排索引数据, %s是项目id(ProjectId), Key是词 值是 []*entity.InvertIndex
	DocInvertIndexListTableCache = NewCache(24*7*time.Hour, 24*1*time.Hour) // 文档有哪些倒排词, %s是文档id(DocId),  Key是倒排词  值是 int64
)
