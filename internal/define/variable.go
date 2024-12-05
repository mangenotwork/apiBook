package define

import (
	"apiBook/common/utils"
	"fmt"
	"time"
)

var (
	Version         = "v0.0.6"
	TimeStamp int64 = time.Now().Unix()
)

const (
	UserToken    string = "sign"
	TokenExpires int    = 60 * 60 * 24 * 7
)

var (
	CsrfAuthKey = ""
	CsrfName    = ""
)

type ProjectPrivateCode int

const (
	ProjectNone    ProjectPrivateCode = iota
	ProjectPublic                     // 公有
	ProjectPrivate                    // 私有
)

type ReqTypeCode string

const (
	ReqTypeNone               = "none"
	ReqTypeText               = "text"
	ReqTypeFormData           = "form-data"
	ReqTypeXWWWFormUrlEncoded = "x-www-form-urlencoded"
	ReqTypeJson               = "json"
	ReqTypeXml                = "xml"
	ReqTypeRaw                = "raw"
	ReqTypeBinary             = "binary"
	ReqTypeGraphQL            = "GraphQL"
)

var ReqTypeArray = []string{ReqTypeNone, ReqTypeText, ReqTypeFormData, ReqTypeXWWWFormUrlEncoded, ReqTypeJson,
	ReqTypeXml, ReqTypeRaw, ReqTypeBinary, ReqTypeGraphQL}

var ReqTypeCodeNameMap = map[ReqTypeCode]string{
	ReqTypeText:               "text/plain",
	ReqTypeFormData:           "multipart/form-data",
	ReqTypeXWWWFormUrlEncoded: "application/x-www-form-urlencoded",
	ReqTypeJson:               "application/json",
	ReqTypeXml:                "application/xml",
	ReqTypeRaw:                "text/plain",
	ReqTypeBinary:             "application/octet-stream",
	ReqTypeGraphQL:            "application/json",
}

func (c ReqTypeCode) GetName() string {
	if value, ok := ReqTypeCodeNameMap[c]; ok {
		return value
	}
	return ""
}

const (
	DirDefault       = "default_%s"    // 默认目录Key
	DirRecycleBinKey = "recycleBin_%s" // 回收站目录Key
)

func GetDirDefault(pid string) string {
	return fmt.Sprintf(DirDefault, pid)
}

func GetDirRecycleBinKey(pid string) string {
	return fmt.Sprintf(DirRecycleBinKey, pid)
}

const (
	DirNameDefault    = "默认"
	DirNameRecycleBin = "回收站"
)

const (
	OperationLogCreateDoc = "创建接口文档"
	OperationLogModifyDoc = "修改了接口文档"
)

func GetSnapshotId() string {
	return fmt.Sprintf("%d%s", time.Now().Unix(), utils.NewShortCode())
}

type SourceCode string

const (
	SourceApiBook    SourceCode = "apiBook"
	SourceOpenApi301 SourceCode = "openApi301"
	SourceOpenApi310 SourceCode = "openApi310"
	SourceSwagger    SourceCode = "swagger"
	SourceApiZZA     SourceCode = "apizza"
	SourceYApi       SourceCode = "yapi"
)
