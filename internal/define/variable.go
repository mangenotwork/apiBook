package define

import "time"

var TimeStamp int64 = time.Now().Unix()

const (
	UserToken    string = "sign"
	TokenExpires int    = 60 * 60 * 24 * 7
)
