package define

import (
	"fmt"
	"time"
)

var (
	Version         = "v0.0.1"
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
