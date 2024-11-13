package handler

import (
	"apiBook/internal/define"
	"apiBook/internal/entity"
)

type ProjectUsersResp struct {
	List      []*UserInfo `json:"list"`
	CreateAcc string      `json:"createAcc"`
}

type ProjectUser struct {
	Name       string `json:"name"`
	Account    string `json:"account"`
	IsCreate   int    `json:"isCreate"`   // 是否是创建者
	CreateTime int64  `json:"createTime"` // 创建时间
	IsDisable  int    `json:"disable"`    // 是否被禁用 1:是
	Pid        string `json:"pid"`
}

type UserInfo struct {
	Name       string `json:"name"`
	Account    string `json:"account"`
	CreateTime int64  `json:"createTime"` // 创建时间
	IsDisable  int    `json:"disable"`    // 是否被禁用 1:是
}

type ProjectAddUserReq struct {
	PId      string `json:"pid"`
	Accounts string `json:"accounts"`
}

type ProjectDelUserReq struct {
	PId     string `json:"pid"`
	Account string `json:"account"`
}

type UserModifyReq struct {
	Name    string `json:"name"`
	Account string `json:"account"`
}

type UserResetPasswordReq struct {
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Account   string `json:"account"`
}

type AdminCreateUserReq struct {
	Name      string `json:"name"`
	Account   string `json:"account"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	IsAdmin   int    `json:"isAdmin"` // 1是
}

type AdminDeleteUserReq struct {
	Account string `json:"account"`
}

type AdminDisableUserReq struct {
	Account   string `json:"account"`
	IsDisable int    `json:"isDisable"`
}

type DocumentDirCreateReq struct {
	PId     string `json:"pid"`
	DirName string `json:"dirName"`
}

type DocumentDirDeleteReq struct {
	PId   string `json:"pid"`
	DirId string `json:"dirId"`
}

type DocumentDirModifyReq struct {
	PId     string `json:"pid"`
	DirId   string `json:"dirId"`
	DirName string `json:"dirName"`
}

type DocumentListReq struct {
	PId   string `json:"pid"`
	DirId string `json:"dirId"`
}

type DocumentDirSortReq struct {
	PId     string   `json:"pid"`
	DirList []string `json:"dirList"`
}

type DocumentSortReq struct {
	PId     string   `json:"pid"`
	DirId   string   `json:"dirId"`
	DocList []string `json:"docList"`
}

type DocumentItemParam struct {
	PId   string `json:"pid"`
	DocId string `json:"docId"`
}

type DocumentItemResp struct {
	Content      *DocumentContent  `json:"content"`
	SnapshotList []*SnapshotItem   `json:"snapshotList"`
	BaseInfo     *entity.Document  `json:"baseInfo"`
	ReqCode      map[string]string `json:"reqCode"`
}

type DocumentContent struct {
	DocId                     string                           `json:"docId"`                     // 文档id
	ProjectId                 string                           `json:"projectId"`                 // 项目id
	Method                    string                           `json:"method"`                    // http 请求类型
	Name                      string                           `json:"name"`                      // 接口名
	Url                       string                           `json:"url"`                       // 接口url
	DescriptionHtml           string                           `json:"descriptionHtml"`           // 接口说明  html
	Description               string                           `json:"description"`               // 接口说明 md
	ReqHeader                 []*entity.ReqHeaderItem          `json:"reqHeader"`                 // 请求头
	ReqType                   define.ReqTypeCode               `json:"reqType"`                   // 请求类型
	ReqTypeName               string                           `json:"reqTypeName"`               // 请求类型名称
	ReqBodyJson               string                           `json:"reqBodyJson"`               // 请求参数 - json
	ReqBodyText               string                           `json:"reqBodyText"`               // 请求参数 - text
	ReqBodyFormData           []*entity.FormDataItem           `json:"reqBodyFormData"`           // 请求参数 - form-data
	ReqBodyXWWWFormUrlEncoded []*entity.XWWWFormUrlEncodedItem `json:"reqBodyXWWWFormUrlEncoded"` // 请求参数 - x-www-form-urlencoded"
	ReqBodyXml                string                           `json:"reqBodyXml"`                // 请求参数 - xml
	ReqBodyRaw                string                           `json:"reqBodyRaw"`                // 请求参数 - raw
	ReqBodyBinary             string                           `json:"reqBodyBinary"`             // 请求参数 - binary
	ReqBodyGraphQL            string                           `json:"reqBodyGraphQL"`            // 请求参数 - GraphQL
	ReqBodyInfo               []*entity.BodyInfoItem           `json:"reqBodyInfo"`               // 请求参数说明
	Resp                      []*entity.RespItem               `json:"resp"`                      // 请求响应
	UserAcc                   string                           `json:"userAcc"`                   // 创建者
	Date                      string                           `json:"date"`                      // 最后一次编辑(创建)时间
}

type SnapshotItem struct {
	SnapshotIdId  string `json:"snapshotId"` // 快照id
	UserAcc       string `json:"userAcc"`    // 操作者
	Operation     string `json:"operation"`  // 操作日志，文本信息
	CreateTime    int64  `json:"createTime"` // 创建时间
	CreateTimeStr string `json:"createTimeStr"`
}

type DocumentDeleteReq struct {
	PId   string `json:"pid"`
	DirId string `json:"dirId"`
	DocId string `json:"docId"`
}

type DocumentChangeDirReq struct {
	PId      string `json:"pid"`
	DirIdNew string `json:"dirIdNew"`
	DocId    string `json:"docId"`
}

type HomeProjectItem struct {
	ProjectId     string                    `json:"projectId"`
	Name          string                    `json:"name"`        // 项目名
	Description   string                    `json:"description"` // 项目简述
	CreateUserAcc string                    `json:"createUserAcc"`
	CreateDate    string                    `json:"createDate"`  // 创建时间
	IsOperation   int                       `json:"isOperation"` // 1 可以操作
	Private       define.ProjectPrivateCode `json:"private"`     // 1:公有(所有人可见)  2:私有
}

type DocumentDirListItem struct {
	Dir *DirRespItem   `json:"dir"`
	Doc []*DocRespItem `json:"doc"`
}

type DirRespItem struct {
	DirId string `json:"dirId"`
	Name  string `json:"name"`
}

type DocRespItem struct {
	DocId  string `json:"docId"`
	Method string `json:"method"`
	Title  string `json:"title"`
}

type DocumentDocListReq struct {
	PId     string   `json:"pid"`
	DocList []string `json:"docList"`
}

type CaseFuncGoReq struct {
	Text string `json:"text"`
}

type ToolGoStructToFieldReq struct {
	Text string `json:"text"`
}

type ToolGoStructToFieldItem struct {
	Field       string `json:"field"`       // 字段
	VarType     string `json:"varType"`     // 类型
	Description string `json:"description"` // 描述
}

type DocumentSnapshotItemReq struct {
	DocId      string `json:"docId"`
	SnapshotId string `json:"snapshotId"`
}

type DocumentSnapshotItemResp struct {
	Item         *entity.DocumentSnapshot `json:"item"`
	SnapshotList []*SnapshotItem          `json:"snapshotList"`
	ReqCode      map[string]string        `json:"reqCode"`
}

type MoveToRecycleBinReq struct {
	PId   string `json:"pid"`
	DocId string `json:"docId"`
}

type DocumentGetDocAllReq struct {
	PId   string `json:"pid"`
	DirId string `json:"dirId"`
}

type DocumentSearchReq struct {
	PId  string `json:"pid"`
	Word string `json:"word"`
}

type DocumentSearchResp struct {
	Count int                       `json:"count"`
	List  []*DocumentSearchRespItem `json:"list"`
}

type DocumentSearchRespItem struct {
	DocId   string `json:"docId"`
	Method  string `json:"method"`
	Title   string `json:"title"`
	Word    string `json:"word"`
	ModType string `json:"modType"`
	Score   int    `json:"score"`
}

type ToolImportReq struct {
}

type ToolExportReq struct {
	Project        string `json:"project"`
	ExportType     string `json:"exportType"`
	SourcePlatform string `json:"sourcePlatform"`
}
