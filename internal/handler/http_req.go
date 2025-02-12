package handler

import (
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"fmt"
	"github.com/gin-gonic/gin"
	http2 "net/http"
	"strings"

	gt "github.com/mangenotwork/gathertool"
)

type RequesterExecuteParam struct {
	Method      string            `json:"method"`      // 请求类型
	Url         string            `json:"url"`         // 请求url
	Header      map[string]string `json:"header"`      // 请求header
	Body        string            `json:"body"`        // 请求body - 非fromData
	FromData    map[string]string `json:"fromData"`    // 请求body - fromData
	ContentType string            `json:"contentType"` // 请求ContentType
}

type RequesterExecuteResp struct {
	RespBody   []byte       `json:"respBody"`
	RespCode   int          `json:"respCode"`
	RespHeader http2.Header `json:"respHeader"`
}

func RequesterExecute(c *gin.Context) {

	ctx := ginHelper.NewGinCtx(c)

	param := &RequesterExecuteParam{}
	err := ctx.GetPostArgs(param)
	if err != nil {
		ctx.APIOutPutError(nil, err.Error())
		return
	}

	if !isMethod(param.Method) {
		ctx.APIOutPutError(fmt.Errorf("未知的请求类型"), "未知的请求类型")
		return
	}

	if len(param.Url) < 1 {
		ctx.APIOutPutError(fmt.Errorf("请求地址为空"), "请求地址为空")
		return
	}

	log.Info(param)

	var gtCtx *gt.Context

	switch strings.ToUpper(param.Method) {
	case "GET":
		gtCtx, err = gt.Get(param.Url)

	case "POST":
		if strings.Contains(param.ContentType, "from") {
			gtCtx, err = gt.PostForm(param.Url, param.FromData)
		} else {
			gtCtx, err = gt.Post(param.Url, []byte(param.Body), param.ContentType)
		}

	case "PUT":
		gtCtx, err = gt.Put(param.Url, []byte(param.Body), param.ContentType)

	case "DELETE":
		gtCtx, err = gt.Delete(param.Url)

	case "OPTIONS":
		gtCtx, err = gt.Options(param.Url)

	case "HEAD":
		gtCtx, err = gt.Request(param.Url, param.Method, []byte(param.Body), param.ContentType)

	}

	if err != nil {
		log.Error(err)
		ctx.APIOutPutError(err, err.Error())
		return
	}

	if gtCtx == nil {
		log.Error(err)
		ctx.APIOutPutError(fmt.Errorf("请求失败，响应对象为空"), "请求失败，响应对象为空")
		return
	}

	resp := &RequesterExecuteResp{
		RespBody:   gtCtx.RespBody,
		RespCode:   gtCtx.StateCode,
		RespHeader: gtCtx.Resp.Header,
	}

	ctx.APIOutPut(resp, "ok")
	return
}

func isMethod(method string) bool {
	rse := false
	method = strings.ToUpper(method)

	for _, v := range []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"} {
		if method == v {
			rse = true
			break
		}
	}

	return rse
}
