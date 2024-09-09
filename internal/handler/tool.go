package handler

import (
	"apiBook/common/ginHelper"
	"apiBook/common/log"
	"apiBook/common/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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

// 代码生成模板

// curl
// curl --location --request POST 'https://api.ecosmos.vip/shop/pingxx' \    // 可变的
//--header 'sign: /<B70o;7W@3W,]dG<20q' \
//--header 'source: 2' \
//--header 'User-Agent: apiBook/0.0.1 (https://github.com/mangenotwork/apiBook)' \   // 可变的
//--header 'Content-Type: application/json' \    // 可变的
//--header 'Accept: */*' \
//--header 'Host: api.ecosmos.vip' \
//--header 'Connection: keep-alive' \
//--data-raw '{
//  "activity_id":2
//}'

// wget
//wget --no-check-certificate --quiet \
//   --method POST \
//   --timeout=0 \
//   --header 'sign: /<B70o;7W@3W,]dG<20q' \
//   --header 'source: 2' \
//   --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
//   --header 'Content-Type: application/json' \
//   --header 'Accept: */*' \
//   --header 'Host: api.ecosmos.vip' \
//   --header 'Connection: keep-alive' \
//   --body-data '{
//  "activity_id":2
//}' \
//    'https://api.ecosmos.vip/shop/pingxx'

// PowerShell
//$headers = New-Object "System.Collections.Generic.Dictionary[[String],[String]]"
//$headers.Add("sign", "/<B70o;7W@3W,]dG<20q")
//$headers.Add("source", "2")
//$headers.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
//$headers.Add("Content-Type", "application/json")
//$headers.Add("Accept", "*/*")
//$headers.Add("Host", "api.ecosmos.vip")
//$headers.Add("Connection", "keep-alive")
//
//$body = "{
//`n  `"activity_id`":2
//`n}"
//
//$response = Invoke-RestMethod 'https://api.ecosmos.vip/shop/pingxx' -Method 'POST' -Headers $headers -Body $body
//$response | ConvertTo-Json

// js - fetch
//var myHeaders = new Headers();
//myHeaders.append("sign", "/<B70o;7W@3W,]dG<20q");
//myHeaders.append("source", "2");
//myHeaders.append("User-Agent", "Apifox/1.0.0 (https://apifox.com)");
//myHeaders.append("Content-Type", "application/json");
//myHeaders.append("Accept", "*/*");
//myHeaders.append("Host", "api.ecosmos.vip");
//myHeaders.append("Connection", "keep-alive");
//
//var raw = JSON.stringify({
//   "activity_id": 2
//});
//
//var requestOptions = {
//   method: 'POST',
//   headers: myHeaders,
//   body: raw,
//   redirect: 'follow'
//};
//
//fetch("https://api.ecosmos.vip/shop/pingxx", requestOptions)
//   .then(response => response.text())
//   .then(result => console.log(result))
//   .catch(error => console.log('error', error));

// js - axios
//var axios = require('axios');
//var data = JSON.stringify({
//   "activity_id": 2
//});
//
//var config = {
//   method: 'post',
//   url: 'https://api.ecosmos.vip/shop/pingxx',
//   headers: {
//      'sign': '/<B70o;7W@3W,]dG<20q',
//      'source': '2',
//      'User-Agent': 'Apifox/1.0.0 (https://apifox.com)',
//      'Content-Type': 'application/json',
//      'Accept': '*/*',
//      'Host': 'api.ecosmos.vip',
//      'Connection': 'keep-alive'
//   },
//   data : data
//};
//
//axios(config)
//.then(function (response) {
//   console.log(JSON.stringify(response.data));
//})
//.catch(function (error) {
//   console.log(error);
//});

// js - jquery
//var settings = {
//   "url": "https://api.ecosmos.vip/shop/pingxx",
//   "method": "POST",
//   "timeout": 0,
//   "headers": {
//      "sign": "/<B70o;7W@3W,]dG<20q",
//      "source": "2",
//      "User-Agent": "Apifox/1.0.0 (https://apifox.com)",
//      "Content-Type": "application/json",
//      "Accept": "*/*",
//      "Host": "api.ecosmos.vip",
//      "Connection": "keep-alive"
//   },
//   "data": JSON.stringify({
//      "activity_id": 2
//   }),
//};
//
//$.ajax(settings).done(function (response) {
//   console.log(response);
//});

// js - xhr
//// WARNING: For POST requests, body is set to null by browsers.
//var data = JSON.stringify({
//   "activity_id": 2
//});
//
//var xhr = new XMLHttpRequest();
//xhr.withCredentials = true;
//
//xhr.addEventListener("readystatechange", function() {
//   if(this.readyState === 4) {
//      console.log(this.responseText);
//   }
//});
//
//xhr.open("POST", "https://api.ecosmos.vip/shop/pingxx");
//xhr.setRequestHeader("sign", "/<B70o;7W@3W,]dG<20q");
//xhr.setRequestHeader("source", "2");
//xhr.setRequestHeader("User-Agent", "Apifox/1.0.0 (https://apifox.com)");
//xhr.setRequestHeader("Content-Type", "application/json");
//xhr.setRequestHeader("Accept", "*/*");
//xhr.setRequestHeader("Host", "api.ecosmos.vip");
//xhr.setRequestHeader("Connection", "keep-alive");
//
//xhr.send(data);

// java - Unirest
//Unirest.setTimeouts(0, 0);
//HttpResponse<String> response = Unirest.post("https://api.ecosmos.vip/shop/pingxx")
//   .header("sign", "/<B70o;7W@3W,]dG<20q")
//   .header("source", "2")
//   .header("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
//   .header("Content-Type", "application/json")
//   .header("Accept", "*/*")
//   .header("Host", "api.ecosmos.vip")
//   .header("Connection", "keep-alive")
//   .body("{\r\n  \"activity_id\":2\r\n}")
//   .asString();

// java - OkHttpClient
//OkHttpClient client = new OkHttpClient().newBuilder()
//   .build();
//MediaType mediaType = MediaType.parse("application/json");
//RequestBody body = RequestBody.create(mediaType, "{\r\n  \"activity_id\":2\r\n}");
//Request request = new Request.Builder()
//   .url("https://api.ecosmos.vip/shop/pingxx")
//   .method("POST", body)
//   .addHeader("sign", "/<B70o;7W@3W,]dG<20q")
//   .addHeader("source", "2")
//   .addHeader("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
//   .addHeader("Content-Type", "application/json")
//   .addHeader("Accept", "*/*")
//   .addHeader("Host", "api.ecosmos.vip")
//   .addHeader("Connection", "keep-alive")
//   .build();
//Response response = client.newCall(request).execute();

// swift
//import Foundation
//#if canImport(FoundationNetworking)
//import FoundationNetworking
//#endif
//
//var semaphore = DispatchSemaphore (value: 0)
//
//let parameters = "{\r\n  \"activity_id\":2\r\n}"
//let postData = parameters.data(using: .utf8)
//
//var request = URLRequest(url: URL(string: "https://api.ecosmos.vip/shop/pingxx")!,timeoutInterval: Double.infinity)
//request.addValue("/<B70o;7W@3W,]dG<20q", forHTTPHeaderField: "sign")
//request.addValue("2", forHTTPHeaderField: "source")
//request.addValue("Apifox/1.0.0 (https://apifox.com)", forHTTPHeaderField: "User-Agent")
//request.addValue("application/json", forHTTPHeaderField: "Content-Type")
//request.addValue("*/*", forHTTPHeaderField: "Accept")
//request.addValue("api.ecosmos.vip", forHTTPHeaderField: "Host")
//request.addValue("keep-alive", forHTTPHeaderField: "Connection")
//
//request.httpMethod = "POST"
//request.httpBody = postData
//
//let task = URLSession.shared.dataTask(with: request) { data, response, error in
//   guard let data = data else {
//      print(String(describing: error))
//      semaphore.signal()
//      return
//   }
//   print(String(data: data, encoding: .utf8)!)
//   semaphore.signal()
//}
//
//task.resume()
//semaphore.wait()

// go
//package main
//
//import (
//   "fmt"
//   "strings"
//   "net/http"
//   "io/ioutil"
//)
//
//func main() {
//
//   url := "https://api.ecosmos.vip/shop/pingxx"
//   method := "POST"
//
//   payload := strings.NewReader(`{
//  	"activity_id":2"
//}`)
//
//   client := &http.Client {
//   }
//   req, err := http.NewRequest(method, url, payload)
//
//   if err != nil {
//      fmt.Println(err)
//      return
//   }
//   req.Header.Add("sign", "/<B70o;7W@3W,]dG<20q")
//   req.Header.Add("source", "2")
//   req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
//   req.Header.Add("Content-Type", "application/json")
//   req.Header.Add("Accept", "*/*")
//   req.Header.Add("Host", "api.ecosmos.vip")
//   req.Header.Add("Connection", "keep-alive")
//
//   res, err := client.Do(req)
//   if err != nil {
//      fmt.Println(err)
//      return
//   }
//   defer res.Body.Close()
//
//   body, err := ioutil.ReadAll(res.Body)
//   if err != nil {
//      fmt.Println(err)
//      return
//   }
//   fmt.Println(string(body))
//}

// php - Request2
//<?php
//require_once 'HTTP/Request2.php';
//$request = new HTTP_Request2();
//$request->setUrl('https://api.ecosmos.vip/shop/pingxx');
//$request->setMethod(HTTP_Request2::METHOD_POST);
//$request->setConfig(array(
//   'follow_redirects' => TRUE
//));
//$request->setHeader(array(
//   'sign' => '/<B70o;7W@3W,]dG<20q',
//   'source' => '2',
//   'User-Agent' => 'Apifox/1.0.0 (https://apifox.com)',
//   'Content-Type' => 'application/json',
//   'Accept' => '*/*',
//   'Host' => 'api.ecosmos.vip',
//   'Connection' => 'keep-alive'
//));
//$request->setBody('{
//\n  "activity_id":2
//\n}');
//try {
//   $response = $request->send();
//   if ($response->getStatus() == 200) {
//      echo $response->getBody();
//   }
//   else {
//      echo 'Unexpected HTTP status: ' . $response->getStatus() . ' ' .
//      $response->getReasonPhrase();
//   }
//}
//catch(HTTP_Request2_Exception $e) {
//   echo 'Error: ' . $e->getMessage();
//}

// php - http Client
//<?php
//$client = new http\Client;
//$request = new http\Client\Request;
//$request->setRequestUrl('https://api.ecosmos.vip/shop/pingxx');
//$request->setRequestMethod('POST');
//$body = new http\Message\Body;
//$body->append('{
//  "activity_id":2
//}');
//$request->setBody($body);
//$request->setOptions(array());
//$request->setHeaders(array(
//   'sign' => '/<B70o;7W@3W,]dG<20q',
//   'source' => '2',
//   'User-Agent' => 'Apifox/1.0.0 (https://apifox.com)',
//   'Content-Type' => 'application/json',
//   'Accept' => '*/*',
//   'Host' => 'api.ecosmos.vip',
//   'Connection' => 'keep-alive'
//));
//$client->enqueue($request)->send();
//$response = $client->getResponse();
//echo $response->getBody();

// php - Client
//<?php
//$client = new Client();
//$headers = [
//   'sign' => '/<B70o;7W@3W,]dG<20q',
//   'source' => '2',
//   'User-Agent' => 'Apifox/1.0.0 (https://apifox.com)',
//   'Content-Type' => 'application/json',
//   'Accept' => '*/*',
//   'Host' => 'api.ecosmos.vip',
//   'Connection' => 'keep-alive'
//];
//$body = '{
//   "activity_id": 2
//}';
//$request = new Request('POST', 'https://api.ecosmos.vip/shop/pingxx', $headers, $body);
//$res = $client->sendAsync($request)->wait();
//echo $res->getBody();

// python - client
//import http.client
//import json
//
//conn = http.client.HTTPSConnection("api.ecosmos.vip")
//payload = json.dumps({
//   "activity_id": 2
//})
//headers = {
//   'sign': '/<B70o;7W@3W,]dG<20q',
//   'source': '2',
//   'User-Agent': 'Apifox/1.0.0 (https://apifox.com)',
//   'Content-Type': 'application/json',
//   'Accept': '*/*',
//   'Host': 'api.ecosmos.vip',
//   'Connection': 'keep-alive'
//}
//conn.request("POST", "/shop/pingxx", payload, headers)
//res = conn.getresponse()
//data = res.read()
//print(data.decode("utf-8"))

// python - requests
//import requests
//import json
//
//url = "https://api.ecosmos.vip/shop/pingxx"
//
//payload = json.dumps({
//   "activity_id": 2
//})
//headers = {
//   'sign': '/<B70o;7W@3W,]dG<20q',
//   'source': '2',
//   'User-Agent': 'Apifox/1.0.0 (https://apifox.com)',
//   'Content-Type': 'application/json',
//   'Accept': '*/*',
//   'Host': 'api.ecosmos.vip',
//   'Connection': 'keep-alive'
//}
//
//response = requests.request("POST", url, headers=headers, data=payload)
//
//print(response.text)

// c#
//var client = new RestClient("https://api.ecosmos.vip/shop/pingxx");
//client.Timeout = -1;
//var request = new RestRequest(Method.POST);
//request.AddHeader("sign", "/<B70o;7W@3W,]dG<20q");
//request.AddHeader("source", "2");
//client.UserAgent = "Apifox/1.0.0 (https://apifox.com)";
//request.AddHeader("Content-Type", "application/json");
//request.AddHeader("Accept", "*/*");
//request.AddHeader("Host", "api.ecosmos.vip");
//request.AddHeader("Connection", "keep-alive");
//var body = @"{
//" + "\n" +
//@"  ""activity_id"":2
//" + "\n" +
//@"}";
//request.AddParameter("application/json", body,  ParameterType.RequestBody);
//IRestResponse response = client.Execute(request);
//Console.WriteLine(response.Content);

// Objective-C
//#import <Foundation/Foundation.h>
//
//dispatch_semaphore_t sema = dispatch_semaphore_create(0);
//
//NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:[NSURL URLWithString:@"https://api.ecosmos.vip/shop/pingxx"]
//   cachePolicy:NSURLRequestUseProtocolCachePolicy
//   timeoutInterval:10.0];
//NSDictionary *headers = @{
//   @"sign": @"/<B70o;7W@3W,]dG<20q",
//   @"source": @"2",
//   @"User-Agent": @"Apifox/1.0.0 (https://apifox.com)",
//   @"Content-Type": @"application/json",
//   @"Accept": @"*/*",
//   @"Host": @"api.ecosmos.vip",
//   @"Connection": @"keep-alive"
//};
//
//[request setAllHTTPHeaderFields:headers];
//NSData *postData = [[NSData alloc] initWithData:[@"{\r\n  \"activity_id\":2\r\n}" dataUsingEncoding:NSUTF8StringEncoding]];
//[request setHTTPBody:postData];
//
//[request setHTTPMethod:@"POST"];
//
//NSURLSession *session = [NSURLSession sharedSession];
//NSURLSessionDataTask *dataTask = [session dataTaskWithRequest:request
//completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
//   if (error) {
//      NSLog(@"%@", error);
//      dispatch_semaphore_signal(sema);
//   } else {
//      NSHTTPURLResponse *httpResponse = (NSHTTPURLResponse *) response;
//      NSError *parseError = nil;
//      NSDictionary *responseDictionary = [NSJSONSerialization JSONObjectWithData:data options:0 error:&parseError];
//      NSLog(@"%@",responseDictionary);
//      dispatch_semaphore_signal(sema);
//   }
//}];
//[dataTask resume];
//dispatch_semaphore_wait(sema, DISPATCH_TIME_FOREVER);

// Dart
//var headers = {
//   'sign': '/<B70o;7W@3W,]dG<20q',
//   'source': '2',
//   'User-Agent': 'Apifox/1.0.0 (https://apifox.com)',
//   'Content-Type': 'application/json',
//   'Accept': '*/*',
//   'Host': 'api.ecosmos.vip',
//   'Connection': 'keep-alive'
//};
//var request = http.Request('POST', Uri.parse('https://api.ecosmos.vip/shop/pingxx'));
//request.body = json.encode({
//   "activity_id": 2
//});
//request.headers.addAll(headers);
//
//http.StreamedResponse response = await request.send();
//
//if (response.statusCode == 200) {
//   print(await response.stream.bytesToString());
//}
//else {
//   print(response.reasonPhrase);
//}
