package handler

import (
	"testing"
)

func Test_ReqCodeTemplate(t *testing.T) {
	obj := NewReqCode(ReqCodeTypeC)
	t.Log(obj.ReqCodeTemplate(&ReqCodeArg{
		Method:      POST,
		Url:         "https://api.ecosmos.vip/webapi/activity/list",
		ContentType: "json",
		Header:      map[string]string{"aaaa": "aaaa"},
		DataRaw: `{
			"type": 0,
			"page": 1,
			"limit": 10
		}`,
	}))
}
