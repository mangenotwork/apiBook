package handler

import (
	"apiBook/common/docIE"
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"apiBook/internal/dao"
	"apiBook/internal/define"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/url"
	"regexp"
	"strings"
)

func ToolGoStructToField(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)

	param := &ToolGoStructToFieldReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	resp := make([]*ToolGoStructToFieldItem, 0)

	for _, v := range strings.Split(param.Text, "\n") {

		if isHaveStr(`(?is:type.*?struct)`, v) || v == "}" {
			continue
		}

		field := ""
		varType := ""
		description := ""

		section := make([]string, 0)

		for _, vItem := range strings.Split(v, " ") {
			if len([]byte(vItem)) > 0 {
				section = append(section, vItem)
			}
		}

		field = utils.CleaningStr(section[0])

		if len(section) > 1 {

			section1 := utils.CleaningStr(section[1])

			switch section1 {
			case "int64", "int", "int32", "uint64", "uint8", "uint16", "uint32", "float32", "float64":
				varType = "number"
			case "string":
				varType = "string"
			case "bool":
				varType = "boolean"
			}

			if varType == "" {

				if strings.Contains(section1, "*") {
					varType = "object"
				}
				if strings.Contains(section1, "[]") {
					varType = "array"
				}

			}
		}

		zsReg := regexp.MustCompile(`(?is:\/\*(.*?)\*\/)`)
		zsList := zsReg.FindAllStringSubmatch(v, -1)
		zsReg.FindStringSubmatch(v)

		if len(zsList) > 0 {
			if len(zsList[0]) > 0 {
				log.Info("zsList = ", zsList[0][len(zsList)])
				description = zsList[0][len(zsList)]
			}
		}

		zsSplit := strings.Split(v, "//")
		if len(zsSplit) > 1 {
			description = zsSplit[len(zsSplit)-1]
		}

		reg := regexp.MustCompile(`(?is:json:"(.*?)")`)
		list := reg.FindAllStringSubmatch(v, -1)
		reg.FindStringSubmatch(v)
		if len(list) > 0 {
			if len(list[0]) > 0 {
				field = list[0][len(list)]
			}
		}

		resp = append(resp, &ToolGoStructToFieldItem{
			Field:       field,
			VarType:     varType,
			Description: description,
		})

	}

	ctx.APIOutPut(resp, "")
	return
}

func ToolReqCode(c *gin.Context) {
	ctx := ginHelper.NewGinCtx(c)
	param := &ReqCodeArg{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	codeType := ctx.Param("codeType")
	obj := NewReqCode(codeType)
	resp := obj.ReqCodeTemplate(param)
	ctx.APIOutPut(resp, "")
	return
}

func GetAllReqCode(param *ReqCodeArg) map[string]string {
	resp := make(map[string]string)
	for _, v := range []string{ReqCodeTypeCurl, ReqCodeTypeWget, ReqCodeTypePowerShell, ReqCodeTypeJSFetch, ReqCodeTypeJSAxios,
		ReqCodeTypeJSJquery, ReqCodeTypeJSXhr, ReqCodeTypeJavaUnirest, ReqCodeTypeJavaOkHttpClient, ReqCodeTypeSwift, ReqCodeTypeGo,
		ReqCodeTypePhpRequest2, ReqCodeTypePhpHttpClient, ReqCodeTypePhpClient, ReqCodeTypePythonClient, ReqCodeTypePythonRequests,
		ReqCodeTypeC, ReqCodeTypeCSharp, ReqCodeTypeObjectiveC, ReqCodeTypeRuby, ReqCodeTypeDart} {
		obj := NewReqCode(v)
		resp[v] = obj.ReqCodeTemplate(param)
	}
	return resp
}

const (
	ReqCodeTypeCurl             = "curl"
	ReqCodeTypeWget             = "wget"
	ReqCodeTypePowerShell       = "powerShell"
	ReqCodeTypeJSFetch          = "jsFetch"
	ReqCodeTypeJSAxios          = "jsAxios"
	ReqCodeTypeJSJquery         = "jsJquery"
	ReqCodeTypeJSXhr            = "jsXhr"
	ReqCodeTypeJavaUnirest      = "javaUnirest"
	ReqCodeTypeJavaOkHttpClient = "javaOkHttpClient"
	ReqCodeTypeSwift            = "swift"
	ReqCodeTypeGo               = "go"
	ReqCodeTypePhpRequest2      = "phpRequest2"
	ReqCodeTypePhpHttpClient    = "phpHttpClient"
	ReqCodeTypePhpClient        = "phpClient"
	ReqCodeTypePythonClient     = "pythonClient"
	ReqCodeTypePythonRequests   = "pythonRequests"
	ReqCodeTypeC                = "c"
	ReqCodeTypeCSharp           = "c#"
	ReqCodeTypeObjectiveC       = "objectiveC"
	ReqCodeTypeRuby             = "ruby"
	ReqCodeTypeDart             = "dart"
)

type MethodType string

const (
	POST    MethodType = "POST"
	GET     MethodType = "GET"
	HEAD    MethodType = "HEAD"
	PUT     MethodType = "PUT"
	DELETE  MethodType = "DELETE"
	PATCH   MethodType = "PATCH"
	OPTIONS MethodType = "OPTIONS"
	ANY     MethodType = ""
)

var DefaultHeader = map[string]string{
	"Accept":     "*/*",
	"Connection": "keep-alive",
}

var ContentTypeMap = map[string]string{
	"json":      "application/json",
	"form":      "application/x-www-form-urlencoded",
	"text":      "text/plain",
	"form-data": "multipart/form-data",
	"xml":       "application/xml",
	"stream":    "application/octet-stream",
}

type ReqCodeTemplateEr interface {
	ReqCodeTemplate(req *ReqCodeArg) string
}

type ReqCodeArg struct {
	Method      MethodType        `json:"method"`
	Url         string            `json:"url"`
	ContentType string            `json:"contentType"`
	Header      map[string]string `json:"header"`
	DataRaw     string            `json:"dataRaw"`
}

func NewReqCode(reqCodeType string) ReqCodeTemplateEr {
	switch reqCodeType {
	case ReqCodeTypeCurl:
		return &ReqCodeCurl{}
	case ReqCodeTypeWget:
		return &ReqCodeWget{}
	case ReqCodeTypePowerShell:
		return &ReqCodePowerShell{}
	case ReqCodeTypeJSFetch:
		return &ReqCodeJSFetch{}
	case ReqCodeTypeJSAxios:
		return &ReqCodeJSAxios{}
	case ReqCodeTypeJSJquery:
		return &ReqCodeJSJquery{}
	case ReqCodeTypeJSXhr:
		return &ReqCodeJSXhr{}
	case ReqCodeTypeJavaUnirest:
		return &ReqCodeJavaUnirest{}
	case ReqCodeTypeJavaOkHttpClient:
		return &ReqCodeJavaOkHttpClient{}
	case ReqCodeTypeSwift:
		return &ReqCodeSwift{}
	case ReqCodeTypeGo:
		return &ReqCodeGo{}
	case ReqCodeTypePhpRequest2:
		return &ReqCodePhpRequest2{}
	case ReqCodeTypePhpHttpClient:
		return &ReqCodePhpHttpClient{}
	case ReqCodeTypePhpClient:
		return &ReqCodePhpClient{}
	case ReqCodeTypePythonClient:
		return &ReqCodePythonClient{}
	case ReqCodeTypePythonRequests:
		return &ReqCodePythonRequests{}
	case ReqCodeTypeC:
		return &ReqCodeC{}
	case ReqCodeTypeCSharp:
		return &ReqCodeCSharp{}
	case ReqCodeTypeObjectiveC:
		return &ReqCodeObjectiveC{}
	case ReqCodeTypeRuby:
		return &ReqCodeRuby{}
	case ReqCodeTypeDart:
		return &ReqCodeDart{}
	}
	return nil
}

type ReqCodeCurl struct {
}

func (*ReqCodeCurl) ReqCodeTemplate(req *ReqCodeArg) string {

	header := fmt.Sprintf("--header 'User-Agent:%s' \\\n", utils.RandAgent())

	if req.ContentType == "json" {
		header += fmt.Sprintf("--header 'Content-Type:%s' \\\n", "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf("--header '%s:%s' \\\n", k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf("--header '%s:%s' \\\n", k, v)
	}

	tpl := fmt.Sprintf(`curl --location --request %s '%s' \
	%s --data-raw '%s'`,
		req.Method,  // %s 请求模式 GET or POST
		req.Url,     // %s 请求Url
		header,      // %s header join   //--header 'sign: /<B70o;7W@3W,]dG<20q' \
		req.DataRaw, // %s 请求参数  data-raw  json需要序列化
	)

	return tpl
}

type ReqCodeWget struct {
}

func (*ReqCodeWget) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf("--header 'User-Agent:%s' \\\n", utils.RandAgent())

	if req.ContentType == "json" {
		header += fmt.Sprintf("--header 'Content-Type:%s' \\\n", "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf("--header '%s:%s' \\\n", k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf("--header '%s:%s' \\\n", k, v)
	}

	tpl := fmt.Sprintf(`wget --no-check-certificate --quiet \
	--method %s \
	%s --body-data '%s' \
    '%s'`,
		req.Method,
		header,
		req.DataRaw,
		req.Url,
	)
	return tpl
}

type ReqCodePowerShell struct {
}

func (*ReqCodePowerShell) ReqCodeTemplate(req *ReqCodeArg) string {
	header := `$headers = New-Object "System.Collections.Generic.Dictionary[[String],[String]]"
`
	header += fmt.Sprintf(`$headers.Add("User-Agent", "%s")
		`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`$headers.Add("Content-Type", "%s")
		`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`$headers.Add("%s", "%s")
		`, k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`$headers.Add("%s", "%s")
		`, k, v)
	}
	tpl := fmt.Sprintf(`%s $body = "%s"
		$response = Invoke-RestMethod '%s' -Method '%s' -Headers $headers -Body $body
		$response | ConvertTo-Json
	`,
		header,
		req.DataRaw,
		req.Url,
		req.Method,
	)
	return tpl
}

type ReqCodeJSFetch struct {
}

func (*ReqCodeJSFetch) ReqCodeTemplate(req *ReqCodeArg) string {
	header := "var headers = new Headers();\n"
	header += fmt.Sprintf(`headers.append("User-Agent", "%s");
`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`headers.append("Content-Type", "%s");
`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`headers.append("%s", "%s");
`, k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`headers.append("%s", "%s");
`, k, v)
	}
	tpl := fmt.Sprintf(`%s
var requestOptions = {
    method: '%s', 
    headers: headers,
    body: JSON.stringify(%s),
    redirect: 'follow'
};
fetch("%s", requestOptions)
   .then(response => response.text())
   .then(result => console.log(result))
   .catch(error => console.log('error', error));
`,
		header,
		req.Method,
		req.DataRaw,
		req.Url,
	)
	return tpl
}

type ReqCodeJSAxios struct {
}

func (*ReqCodeJSAxios) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf("\t'User-Agent':'%s',\n", utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf("\t'Content-Type':'%s',\n", "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf("\t'%s':'%s',\n", k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf("\t'%s':'%s',\n", k, v)
	}

	tpl := fmt.Sprintf(`var axios = require('axios');
var config = {
   method: '%s',
   url: '%s',
   headers: {
%s},
   data : JSON.stringify(%s)
};

axios(config)
.then(function (response) {
   console.log(JSON.stringify(response.data));
})
.catch(function (error) {
   console.log(error);
});
`,
		req.Method,
		req.Url,
		header,
		req.DataRaw,
	)
	return tpl
}

type ReqCodeJSJquery struct {
}

func (*ReqCodeJSJquery) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf("\t'User-Agent':'%s',\n", utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf("\t'Content-Type':'%s',\n", "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf("\t'%s':'%s',\n", k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf("\t'%s':'%s',\n", k, v)
	}

	tpl := fmt.Sprintf(`var settings = {
   "url": "%s",
   "method": "%s",
   "headers": {
%s},
   "data": JSON.stringify(%s),
};
$.ajax(settings).done(function (response) {
   console.log(response);
});`,
		req.Url,
		req.Method,
		header,
		req.DataRaw,
	)
	return tpl
}

type ReqCodeJSXhr struct {
}

func (*ReqCodeJSXhr) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf(`xhr.setRequestHeader("User-Agent", "%s");
`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`xhr.setRequestHeader("Content-Type", "%s");
`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`xhr.setRequestHeader("%s", "%s");
`, k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`xhr.setRequestHeader("%s", "%s");
`, k, v)
	}
	tpl := fmt.Sprintf(`var xhr = new XMLHttpRequest();
xhr.withCredentials = true;

xhr.addEventListener("readystatechange", function() {
   if(this.readyState === 4) {
      console.log(this.responseText);
   }
});

xhr.open("%s", "%s");
%s
xhr.send(JSON.stringify(%s))`,
		req.Method,
		req.Url,
		header,
		req.DataRaw)
	return tpl
}

type ReqCodeJavaUnirest struct {
}

func (*ReqCodeJavaUnirest) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf(`.header("User-Agent", "%s")
`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`.header("Content-Type", "%s")
`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`.header("%s", "%s");
`, k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`.header("%s", "%s");
`, k, v)
	}

	method := strings.ToLower(string(req.Method))

	data := strings.ReplaceAll(req.DataRaw, "\"", "\\\"")

	tpl := fmt.Sprintf(`Unirest.setTimeouts(0, 0);
HttpResponse<String> response = Unirest.%s("%s")
%s.body("%s")
   .asString();`,
		method,
		req.Url,
		header,
		data,
	)
	return tpl
}

type ReqCodeJavaOkHttpClient struct {
}

func (*ReqCodeJavaOkHttpClient) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf(`.addHeader("User-Agent", "%s")
`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`.addHeader("Content-Type", "%s")
`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`.addHeader("%s", "%s");
`, k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`.addHeader("%s", "%s");
`, k, v)
	}

	data := strings.ReplaceAll(req.DataRaw, "\"", "\\\"")

	tpl := fmt.Sprintf(`OkHttpClient client = new OkHttpClient().newBuilder().build();
MediaType mediaType = MediaType.parse("application/json");
RequestBody body = RequestBody.create(mediaType, "%s");
Request request = new Request.Builder()
   .url("%s")
   .method("%s", body)
%s.build();
Response response = client.newCall(request).execute();`,
		data,
		req.Url,
		req.Method,
		header,
	)
	return tpl
}

type ReqCodeSwift struct {
}

func (*ReqCodeSwift) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf(`request.addValue("%s", forHTTPHeaderField: "User-Agent");
`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`request.addValue("%s", forHTTPHeaderField: "Content-Type");
`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`request.addValue("%s", forHTTPHeaderField: "%s");
`, v, k)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`request.addValue("%s", forHTTPHeaderField: "%s");
`, v, k)
	}

	data := strings.ReplaceAll(req.DataRaw, "\"", "\\\"")

	tpl := fmt.Sprintf(`import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif

var semaphore = DispatchSemaphore (value: 0)

let parameters = "%s"
let postData = parameters.data(using: .utf8)

var request = URLRequest(url: URL(string: "%s")!,timeoutInterval: Double.infinity)
%s
request.httpMethod = "%s"
request.httpBody = postData

let task = URLSession.shared.dataTask(with: request) { data, response, error in
   guard let data = data else {
      print(String(describing: error))
      semaphore.signal()
      return
   }
   print(String(data: data, encoding: .utf8)!)
   semaphore.signal()
}

task.resume()
semaphore.wait()`,
		data,
		req.Url,
		header,
		req.Method)
	return tpl
}

type ReqCodeGo struct {
}

func (*ReqCodeGo) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf(`req.Header.Add("User-Agent", "%s")
	`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`req.Header.Add("Content-Type", "%s")
	`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`req.Header.Add("%s", "%s")
	`, k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`req.Header.Add("%s", "%s")
	`, k, v)
	}

	data := strings.ReplaceAll(req.DataRaw, "\"", "\\\"")

	tpl := fmt.Sprintf(`package main

import (
   "fmt"
   "strings"
   "net/http"
   "io/ioutil"
)

func main() {
   url := "%s"
   method := "%s"
   payload := strings.NewReader("%s")
   client := &http.Client {}
   req, err := http.NewRequest(method, url, payload)
   if err != nil {
      fmt.Println(err)
      return
   }
   %s
   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()
   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(string(body))
}`,
		req.Url,
		req.Method,
		data,
		header)
	return tpl
}

type ReqCodePhpRequest2 struct {
}

func (*ReqCodePhpRequest2) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf("\t'User-Agent' => '%s',\n", utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf("\t'Content-Type' => '%s',\n", "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf("\t'%s' => '%s',\n", k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf("\t'%s' => '%s',\n", k, v)
	}

	tpl := fmt.Sprintf(`<?php
require_once 'HTTP/Request2.php';
$request = new HTTP_Request2();
$request->setUrl('%s');
$request->setMethod(HTTP_Request2::METHOD_%s);
$request->setConfig(array(
   'follow_redirects' => TRUE
));
$request->setHeader(array(
%s));
$request->setBody('%s');
try {
   $response = $request->send();
   if ($response->getStatus() == 200) {
      echo $response->getBody();
   }
   else {
      echo 'Unexpected HTTP status: ' . $response->getStatus() . ' ' .
      $response->getReasonPhrase();
   }
}
catch(HTTP_Request2_Exception $e) {
   echo 'Error: ' . $e->getMessage();
}`,
		req.Url,
		req.Method,
		header,
		req.DataRaw)
	return tpl
}

type ReqCodePhpHttpClient struct {
}

func (*ReqCodePhpHttpClient) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf("\t'User-Agent' => '%s',\n", utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf("\t'Content-Type' => '%s',\n", "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf("\t'%s' => '%s',\n", k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf("\t'%s' => '%s',\n", k, v)
	}

	tpl := fmt.Sprintf(`<?php
$client = new http\Client;
$request = new http\Client\Request;
$request->setRequestUrl('%s');
$request->setRequestMethod('%s');
$body = new http\Message\Body;
$body->append('%s');
$request->setBody($body);
$request->setOptions(array());
$request->setHeaders(array(
%s));
$client->enqueue($request)->send();
$response = $client->getResponse();
echo $response->getBody();`,
		req.Url,
		req.Method,
		req.DataRaw,
		header)
	return tpl
}

type ReqCodePhpClient struct {
}

func (*ReqCodePhpClient) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf("\t'User-Agent' => '%s',\n", utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf("\t'Content-Type' => '%s',\n", "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf("\t'%s' => '%s',\n", k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf("\t'%s' => '%s',\n", k, v)
	}

	tpl := fmt.Sprintf(`<?php
$client = new Client();
$headers = [
%s
];
$body = '%s';
$request = new Request('%s', '%s', $headers, $body);
$res = $client->sendAsync($request)->wait();
echo $res->getBody();`,
		header,
		req.DataRaw,
		req.Method,
		req.Url)
	return tpl
}

type ReqCodePythonClient struct {
}

func (*ReqCodePythonClient) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf("\t'User-Agent': '%s',\n", utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf("\t'Content-Type': '%s',\n", "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf("\t'%s': '%s',\n", k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf("\t'%s': '%s',\n", k, v)
	}

	host := ""
	path := ""
	if obj, err := url.Parse(req.Url); err == nil {
		host = obj.Host
		path = obj.Path
	}

	tpl := fmt.Sprintf(`import http.client
import json

conn = http.client.HTTPSConnection("%s")
payload = json.dumps(%s)
headers = {
%s}
conn.request("%s", "%s", payload, headers)
res = conn.getresponse()
data = res.read()
print(data.decode("utf-8"))`,
		host,
		req.DataRaw,
		header,
		req.Method,
		path)
	return tpl
}

type ReqCodePythonRequests struct {
}

func (*ReqCodePythonRequests) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf("\t'User-Agent': '%s',\n", utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf("\t'Content-Type': '%s',\n", "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf("\t'%s': '%s',\n", k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf("\t'%s': '%s',\n", k, v)
	}

	tpl := fmt.Sprintf(`import requests
import json

url = "%s"
payload = json.dumps(%s)
headers = {
%s}
response = requests.request("%s", url, headers=headers, data=payload)
print(response.text)`,
		req.Url,
		req.DataRaw,
		header,
		req.Method)
	return tpl
}

type ReqCodeCSharp struct {
}

func (*ReqCodeCSharp) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf(`request.AddHeader("User-Agent", "%s");
`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`request.AddHeader("Content-Type", "%s");
`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`request.AddHeader("%s", "%s");
`, k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`request.AddHeader("%s", "%s");
`, k, v)
	}

	tpl := fmt.Sprintf(`var client = new RestClient("%s");
client.Timeout = -1;
var request = new RestRequest(Method.%s);
%svar body = @"%s";
request.AddParameter("application/json", body,  ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
Console.WriteLine(response.Content);`,
		req.Url,
		req.Method,
		header,
		req.DataRaw)
	return tpl
}

type ReqCodeC struct {
}

func (*ReqCodeC) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf(`   headers = curl_slist_append(headers, "User-Agent: %s");
`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`   headers = curl_slist_append(headers, "Content-Type: %s");
`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`   headers = curl_slist_append(headers, "%s: %s");
`, k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`   headers = curl_slist_append(headers, "%s: %s");
`, k, v)
	}

	data := strings.ReplaceAll(req.DataRaw, "\"", "\\\"")

	scheme := ""
	if obj, err := url.Parse(req.Url); err == nil {
		scheme = obj.Scheme
	}

	tpl := fmt.Sprintf(`CURL *curl;
CURLcode res;
curl = curl_easy_init();
if(curl) {
   curl_easy_setopt(curl, CURLOPT_CUSTOMREQUEST, "%s");
   curl_easy_setopt(curl, CURLOPT_URL, "%s");
   curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);
   curl_easy_setopt(curl, CURLOPT_DEFAULT_PROTOCOL, "%s");
   struct curl_slist *headers = NULL;
%s   curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
   const char *data = "%s";
   curl_easy_setopt(curl, CURLOPT_POSTFIELDS, data);
   res = curl_easy_perform(curl);
}
curl_easy_cleanup(curl);`,
		req.Method,
		req.Url,
		scheme,
		header,
		data)
	return tpl
}

type ReqCodeObjectiveC struct {
}

func (*ReqCodeObjectiveC) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf(`		@"User-Agent": @"%s",
		`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`@"Content-Type": @"%s",
		`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`@"%s": @"%s",
		`, k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`@"%s": @"%s",
		`, k, v)
	}

	data := strings.ReplaceAll(req.DataRaw, "\"", "\\\"")

	tpl := fmt.Sprintf(`#import <Foundation/Foundation.h>

dispatch_semaphore_t sema = dispatch_semaphore_create(0);

NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:[NSURL URLWithString:@"%s"]
   cachePolicy:NSURLRequestUseProtocolCachePolicy
   timeoutInterval:10.0];
NSDictionary *headers = @{
%s};

[request setAllHTTPHeaderFields:headers];
NSData *postData = [[NSData alloc] initWithData:[@"%s" dataUsingEncoding:NSUTF8StringEncoding]];
[request setHTTPBody:postData];

[request setHTTPMethod:@"%s"];

NSURLSession *session = [NSURLSession sharedSession];
NSURLSessionDataTask *dataTask = [session dataTaskWithRequest:request
completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
   if (error) {
      NSLog(@"%%@", error);
      dispatch_semaphore_signal(sema);
   } else {
      NSHTTPURLResponse *httpResponse = (NSHTTPURLResponse *) response;
      NSError *parseError = nil;
      NSDictionary *responseDictionary = [NSJSONSerialization JSONObjectWithData:data options:0 error:&parseError];
      NSLog(@"%%@",responseDictionary);
      dispatch_semaphore_signal(sema);
   }
}];
[dataTask resume];
dispatch_semaphore_wait(sema, DISPATCH_TIME_FOREVER);`,
		req.Url,
		header,
		data,
		req.Method)
	return tpl
}

type ReqCodeRuby struct {
}

func (*ReqCodeRuby) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf(`request["User-Agent"] = "%s"
`, utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf(`request["Content-Type"] = "%s"
`, "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf(`request["%s"] = "%s"
`, k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf(`request["%s"] = "%s"
`, k, v)
	}

	tpl := fmt.Sprintf(`require "uri"
require "json"
require "net/http"

url = URI("%s")

https = Net::HTTP.new(url.host, url.port)
https.use_ssl = true

request = Net::HTTP::%s.new(url)
%s
request.body = JSON.dump(%s)

response = https.request(request)
puts response.read_body`,
		req.Url,
		strings.ToTitle(string(req.Method)),
		header,
		req.DataRaw)
	return tpl
}

type ReqCodeDart struct {
}

func (*ReqCodeDart) ReqCodeTemplate(req *ReqCodeArg) string {
	header := fmt.Sprintf("\t'User-Agent': '%s',\n", utils.RandAgent())
	if req.ContentType == "json" {
		header += fmt.Sprintf("\t'Content-Type': '%s',\n", "application/json")
	}
	for k, v := range DefaultHeader {
		header += fmt.Sprintf("\t'%s': '%s',\n", k, v)
	}
	for k, v := range req.Header {
		header += fmt.Sprintf("\t'%s': '%s',\n", k, v)
	}

	tpl := fmt.Sprintf(`var headers = {
%s};
var request = http.Request('%s', Uri.parse('%s'));
request.body = json.encode(%s);
request.headers.addAll(headers);

http.StreamedResponse response = await request.send();

if (response.statusCode == 200) {
   print(await response.stream.bytesToString());
}
else {
   print(response.reasonPhrase);
}`,
		header,
		req.Method,
		req.Url,
		req.DataRaw,
	)
	return tpl
}

func ToolImport(c *gin.Context) {

	ctx := ginHelper.NewGinCtx(c)

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	log.Info(form)

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	sourcePlatform := define.SourceCode(form.Value["sourcePlatform"][0])

	obj, err := docIE.NewDocImport(sourcePlatform)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	files := form.File["files"][0]

	src, err := files.Open()
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}
	defer src.Close()

	buffer := make([]byte, files.Size)
	_, err = src.Read(buffer)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	fileContent := string(buffer)

	log.Info(fileContent)

	projectId := form.Value["project"][0]
	if len(projectId) == 0 {
		err = obj.Whole(fileContent, userAcc, define.ProjectPublic)
		if err != nil {
			ctx.APIOutPutError(err, "导入失败")
			return
		}
	} else {
		_, err = dao.NewProjectDao().Get(projectId, userAcc, false)
		if err != nil {
			ctx.APIOutPutError(err, err.Error())
			return
		}
		err = obj.Increment(fileContent, projectId, userAcc, "")
		if err != nil {
			ctx.APIOutPutError(err, "导入失败")
			return
		}
	}

	log.SendOperationLog(userAcc, fmt.Sprintf("导入成功"))

	ctx.APIOutPut("导入成功", "导入成功")
	return
}

func ToolExport(c *gin.Context) {

	ctx := ginHelper.NewGinCtx(c)

	param := &ToolExportReq{}
	err := ctx.GetPostArgs(&param)
	if err != nil {
		ctx.APIOutPutError(fmt.Errorf("参数错误"), "参数错误")
		return
	}

	if param.Project == "" {
		ctx.APIOutPutError(fmt.Errorf("项目id为空"), "参数错误")
		return
	}

	userAcc := ctx.GetString("userAcc")
	if userAcc == "" {
		ctx.AuthErrorOut()
		return
	}

	projectInfo, err := dao.NewProjectDao().Get(param.Project, userAcc, false)
	if err != nil {
		ctx.AuthErrorOut()
		return
	}

	switch param.ExportType {

	case "json":
		data, err := docIE.NewDocExport(define.SourceCode(param.SourcePlatform))
		if err != nil {
			log.Error(err)
			ctx.APIOutPutError(err, "")
			return
		}

		log.Info(param)

		jsonData := data.ExportJson(param.Project)
		reader := strings.NewReader(jsonData)

		fileName := fmt.Sprintf("%s-%s-%s.json", stringToUnicode(projectInfo.Name), param.SourcePlatform, utils.NowDateNotLine())
		log.Info("fileName = ", fileName)

		c.Header("Content-Disposition", fileName)
		c.Header("Content-Type", "application/json")

		_, err = io.Copy(c.Writer, reader)
		if err != nil {
			ctx.APIOutPutError(err, "")
		}

		log.SendOperationLog(userAcc, fmt.Sprintf("导出成功"))

		return

	case "pdf":
		// todo ...
		ctx.APIOutPutError(fmt.Errorf("todo..."), "")
		return

	case "word":
		// todo ...
		ctx.APIOutPutError(fmt.Errorf("todo..."), "")
		return

	}

	ctx.APIOutPutError(fmt.Errorf("未知导出类型"), "")
	return
}

func stringToUnicode(s string) string {
	var result string
	for _, char := range s {
		result += fmt.Sprintf("\\u%04x", char)
	}
	return result
}
