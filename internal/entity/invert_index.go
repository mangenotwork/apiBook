package entity

import "apiBook/common/fenci"

type InvertIndex struct {
	DocId      string // 文档id
	Word       string // 词
	Sentence   string // 原句
	ModType    string // 原句类型   title:标题  description:文档说明  header:请求header  req:请求参数  resp:响应参数
	CreateTime int64  // 生成时间
	Term       *fenci.Term
	Score      int
}
