package entity

import "apiBook/internal/define"

type DocumentParam struct {
	ProjectId string           `json:"projectId"` // 项目id
	DirId     string           `json:"dirId"`     // 目录id  可变的
	Content   *DocumentContent `json:"content"`   // 文档内容
}

// Document 项目接口文档
// Table 是 ProjectId
// Key 是 DocId
type Document struct {
	DocId     string `json:"docId"`     // 文档id
	DirId     string `json:"dirId"`     // 目录id  可变的
	ProjectId string `json:"projectId"` // 项目id
	Name      string `json:"name"`      // 接口名
	Url       string `json:"url"`       // 接口url
	Method    string `json:"method"`    // http 请求类型
}

// DocumentContent 文档内容
// Key 是 DocId
type DocumentContent struct {
	DocId                     string                    `json:"docId"`                     // 文档id
	ProjectId                 string                    `json:"projectId"`                 // 项目id
	Name                      string                    `json:"name"`                      // 接口名
	Url                       string                    `json:"url"`                       // 接口url
	Method                    string                    `json:"method"`                    // http 请求类型
	Description               string                    `json:"description"`               // 接口说明  md文本格式
	DescriptionHtml           string                    `json:"descriptionHtml"`           // 接口说明  html
	ReqHeader                 []*ReqHeaderItem          `json:"reqHeader"`                 // 请求头
	ReqType                   define.ReqTypeCode        `json:"reqType"`                   // 请求类型
	ReqBodyJson               string                    `json:"reqBodyJson"`               // 请求参数 - json
	ReqBodyText               string                    `json:"reqBodyText"`               // 请求参数 - text
	ReqBodyFormData           []*FormDataItem           `json:"reqBodyFormData"`           // 请求参数 - form-data
	ReqBodyXWWWFormUrlEncoded []*XWWWFormUrlEncodedItem `json:"reqBodyXWWWFormUrlEncoded"` // 请求参数 - x-www-form-urlencoded"
	ReqBodyXml                string                    `json:"reqBodyXml"`                // 请求参数 - xml
	ReqBodyRaw                string                    `json:"reqBodyRaw"`                // 请求参数 - raw
	ReqBodyBinary             string                    `json:"reqBodyBinary"`             // 请求参数 - binary
	ReqBodyGraphQL            string                    `json:"reqBodyGraphQL"`            // 请求参数 - GraphQL
	ReqBodyInfo               []*BodyInfoItem           `json:"reqBodyInfo"`               // 请求参数说明
	Resp                      []*RespItem               `json:"resp"`                      // 请求响应
	CreateTime                int64                     `json:"createTime"`                // 创建时间
	UserAcc                   string                    `json:"userAcc"`                   // 创建者
}

func (data *DocumentContent) GetReqHeaderMap() map[string]string {
	result := make(map[string]string)
	for _, v := range data.ReqHeader {
		result[v.Field] = v.Example
	}
	return result
}

type ReqHeaderItem struct {
	Field       string `json:"field"`       // 字段
	VarType     string `json:"varType"`     // 类型
	Description string `json:"description"` // 描述
	Example     string `json:"example"`     // 示例
	IsRequired  int    `json:"isRequired"`  // 是否必填 1:必填
	Sort        int    `json:"sort"`        // 排序
	IsOpen      int    `json:"isOpen"`      // 是否启用
}

type FormDataItem struct {
	Name  string `json:"name"`  // 参数名
	Value string `json:"value"` // 参数值
}

type XWWWFormUrlEncodedItem struct {
	Name  string `json:"name"`  // 参数名
	Value string `json:"value"` // 参数值
}

type BodyInfoItem struct {
	Field       string `json:"field"`       // 字段
	VarType     string `json:"varType"`     // 类型
	Description string `json:"description"` // 描述
	Example     string `json:"example"`     // 示例
	IsRequired  int    `json:"isRequired"`  // 是否必填 1:必填
	Sort        int    `json:"sort"`        // 排序
	IsOpen      int    `json:"isOpen"`      // 是否启用
}

type RespItem struct {
	Tag          string             `json:"tag"`          // 默认 成功
	RespType     define.ReqTypeCode `json:"respType"`     // 响应类型
	RespTypeName string             `json:"respTypeName"` // 响应类型名称
	RespBody     string             `json:"respBody"`     // 响应参数
	RespBodyInfo []*BodyInfoItem    `json:"respBodyInfo"` // 响应参数说明
}

// DocumentSnapshot 文档快照
// Table 是 DocId
// Key 是 快照id SnapshotId
type DocumentSnapshot struct {
	SnapshotIdId    string           `json:"snapshotId"`      // 快照id
	UserAcc         string           `json:"userAcc"`         // 操作者
	Operation       string           `json:"operation"`       // 操作日志，文本信息
	CreateTime      int64            `json:"createTime"`      // 创建时间
	DocumentContent *DocumentContent `json:"documentContent"` // 文档快照
}
