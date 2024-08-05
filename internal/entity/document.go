package entity

// Document 项目接口文档
// Table 是 ProjectId
// Key 是 DocId
type Document struct {
	DocId     string // 文档id
	DirId     string // 目录id  可变的
	ProjectId string // 项目id
}

// DocumentContent 文档内容
// Key 是 DocId
type DocumentContent struct {
	DocId       string // 文档id
	ProjectId   string // 项目id
	Name        string // 接口名
	Url         string // 接口url
	Method      string // http 请求类型
	Description string // 接口说明  md文本格式
	ReqHeader   string // 请求头
	// 请求参数类型
	// 请求参数
	// 请求参数说明
	// 请求响应
}

// DocumentResources 文档资源
// Key 是 DocId
type DocumentResources struct {
}
