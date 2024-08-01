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
