package define

import (
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
