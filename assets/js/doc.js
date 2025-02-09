// 渲染api文档内容
function loadApiDoc(apiDoc, snapshotList) {
    var docIdContent = $("#docIdContent")
    var apiName = $("#api-name")
    var apiMethod = $("#api-method")
    var apiUrl = $("#api-url")
    var apiDescription = $("#api-description")
    var apiReqHeader = $("#api-reqHeader")
    var apiReqType = $("#api-reqType")
    var apiReqTypeStr = $("#api-reqTypeStr")
    var apiReqBodyInfo = $("#api-reqBodyInfo")
    var apiRespType = $("#api-respType")
    var apiRespTypeStr = $("#api-respTypeStr")
    var apiRespBodyInfo = $("#api-respBodyInfo")

    docIdContent.empty();
    apiName.empty();
    apiMethod.empty();
    apiUrl.empty();
    apiDescription.empty();
    apiReqHeader.empty();
    apiReqTypeStr.empty();
    apiReqBodyInfo.empty();
    apiRespTypeStr.empty();
    apiRespBodyInfo.empty();

    docIdContent.append(apiDoc.docId);
    apiName.append(apiDoc.name);

    var methodStr = ""
    switch (apiDoc.method) {
        case "GET":methodStr = GETSpan;break;
        case "DELETE":methodStr = DELETESpan;break;
        case "POST":methodStr = POSTSpan;break;
        case "HEAD":methodStr = HEADSpan;break;
        case "PUT":methodStr = PUTSpan;break;
        case "OPTIONS":methodStr = OPTIONSSpan;break;
    }
    apiMethod.append(methodStr);

    apiUrl.append(apiDoc.url);
    apiDescription.html(apiDoc.descriptionHtml);

    if ( apiDoc.reqHeader !== undefined ) {
        for(var i=0; i<apiDoc.reqHeader.length; i++) {
            var item = apiDoc.reqHeader[i];
            apiReqHeader.append(apiReqHeaderTpl(item.field, item.varType, item.isRequired, item.description, item.example))
        }
    }

    apiReqTypeStr.append(apiDoc.reqTypeName);

    if ( apiDoc.reqBodyInfo !== undefined ) {
        for (var i = 0; i < apiDoc.reqBodyInfo.length; i++) {
            var item = apiDoc.reqBodyInfo[i];

            zs["jsonpath:" + item.field.replace(/\./g, ",")] = item.description

            apiReqBodyInfo.append(apiReqBodyInfoTpl(item.field, item.varType, item.isRequired, item.description, item.example, i))
            if (item.example.length > 20) {
                tippy('.reqExampleTxt' + i, {
                    content: apiDoc.reqBodyInfo[i].example
                });
            }

            if (item.description.length > 20) {
                tippy('.reqDescriptionTxt' + i, {
                    content: item.description
                });
            }
        }
    }

    var jsoneditorDiv = $("#jsoneditorDiv");
    var reqXmlDiv = $("#reqXmlDiv");
    var reqTxtDiv = $("#reqTxtDiv");
    var reqXmlDivBody = $("#reqXmlDiv .card-body");
    var reqTxtDivBody = $("#reqTxtDiv .card-body");
    jsoneditorDiv.hide();
    reqXmlDiv.hide();
    reqTxtDiv.hide();
    editor.set({});
    editor.expandAll();
    reqXmlDivBody.empty();
    reqTxtDivBody.empty();

    if (apiDoc.reqType === "json" && apiDoc.reqBodyJson !== "" ) {
        jsoneditorDiv.show();

        console.log(apiDoc.reqBodyJson)

        const initialJson = JSON.parse(apiDoc.reqBodyJson)
        editor.set(initialJson)
        editor.expandAll()
        const updatedJson = editor.get()
        const paths = getJsonPaths(updatedJson);
        $("#jsoneditor").find(".jsoneditor-field").each(function() {
            var id = $(this).prop("id")
            var regex = new RegExp("0,", "g");
            id = id.replace(regex, "");
            if (zs[id]){
                $(this).parent().parent().append('<td class="zsReqText zsText"> /* '+zs[id]+' */</td>')
            }
        })
    }

    if (apiDoc.reqType === "xml") {
        reqXmlDiv.show();
        reqXmlDivBody.append(apiDoc.reqBodyXml.replace(/\n/g, '<br>'));
    }

    if (apiDoc.reqType === "text") {
        reqTxtDiv.show();
        reqTxtDivBody.append(apiDoc.reqBodyText.replace(/\n/g, '<br>'));
    }

    if ( apiDoc.resp !== undefined && apiDoc.resp.length > 0 ) {
        apiRespTypeStr.append(apiDoc.resp[0].respTypeName);

        if (apiDoc.resp[0].respBodyInfo != null) {
            for(var i=0; i<apiDoc.resp[0].respBodyInfo.length; i++) {
                var item = apiDoc.resp[0].respBodyInfo[i];
                zsresp["jsonpath:"+item.field.replace(/\./g, ",")] = item.description
                apiRespBodyInfo.append(apiRespBodyInfoTpl(item.field, item.varType, item.description, item.example, i))
                if(item.example.length > 20) {
                    tippy('.respExampleTxt'+i, { content: item.example });
                }
                if(item.description.length>20) {
                    tippy('.respDescriptionTxt'+i, { content: item.description });
                }
            }
        }

        var jsoneditorRespDiv = $("#jsoneditorRespDiv");
        var respTxtDiv = $("#respTxtDiv");
        var respXmlDiv = $("#respXmlDiv");
        var respTxtDivBody = $("#respTxtDiv .card-body");
        var respXmlDivBody = $("#respXmlDiv .card-body");
        jsoneditorRespDiv.hide();
        respTxtDiv.hide();
        respXmlDiv.hide();
        editorResp.set({});
        editorResp.expandAll();
        respTxtDivBody.empty();
        respXmlDivBody.empty();

        if (apiDoc.resp[0].respType === "json" && apiDoc.resp[0].respBody !== "") {
            jsoneditorRespDiv.show();
            const initialJson = JSON.parse(apiDoc.resp[0].respBody);
            editorResp.set(initialJson);
            editorResp.expandAll();
            const updatedJson = editorResp.get();
            const paths = getJsonPaths(updatedJson);
            $("#jsoneditorResp").find(".jsoneditor-field").each(function() {
                var id = $(this).prop("id");
                var regex = new RegExp("0,", "g");
                id = id.replace(regex, "");
                if (zsresp[id]){
                    $(this).parent().parent().append('<td class="zsRespText zsText"> /* '+zsresp[id]+' */</td>')
                }
            })
        }

        if (apiDoc.resp[0].respType === "text" && apiDoc.resp[0].respBody !== "") {
            respTxtDiv.show();
            respTxtDivBody.append(apiDoc.resp[0].respBody.replace(/\n/g, '<br>'));
        }

        if (apiDoc.resp[0].respType === "xml" && apiDoc.resp[0].respBody !== "") {
            respXmlDiv.show();
            respXmlDivBody.append(apiDoc.resp[0].respBody.replace(/\n/g, '<br>'));
        }

        var snapshot = $("#snapshot");
        snapshot.empty()
        if (snapshotList != null) {
            for (var i=0; i< snapshotList.length; i++) {
                var item = snapshotList[i];
                snapshot.append('<li class="list-group-item" style="font-size: 12px;">'+ item.userAcc+ ' - ' + item.createTimeStr+
                    ' - ' + item.operation + ' | 镜像: <a type="button" class="btn btn-link" ' +
                    'onclick="openSnapshot(\''+apiDoc.docId+'\',\''+item.snapshotId+'\')">' + item.snapshotId +'</a> </li>')
            }
        }
    }


}

function loadReqCode(param) {

    console.log("loadReqCode ....")

    AjaxPostNotAsync("/document/reqCode", param, function(data){
        console.log(data)
        var reqCodeMap = data.data

        console.log(reqCodeMap["jsFetch"])

        $("#code-jsFetch-code").html(reqCodeMap["jsFetch"]);
        $("#code-jsAxios pre code").html(reqCodeMap["jsAxios"]);
        $("#code-jsJquery pre code").html(reqCodeMap["jsJquery"]);
        $("#code-jsXhr pre code").html(reqCodeMap["jsXhr"]);
        $("#code-swift pre code").html(reqCodeMap["swift"]);
        $("#code-objectiveC pre code").html(reqCodeMap["objectiveC"]);
        $("#code-dart pre code").html(reqCodeMap["dart"]);
        $("#code-javaUnirest pre code").html(reqCodeMap["javaUnirest"]);
        $("#code-javaOkHttpClient pre code").html(reqCodeMap["javaOkHttpClient"]);
        $("#code-curl pre code").html(reqCodeMap["curl"]);
        $("#code-wget pre code").html(reqCodeMap["wget"]);
        $("#code-powerShell pre code").html(reqCodeMap["powerShell"]);
        $("#code-phpRequest2 pre code").html(reqCodeMap["phpRequest2"]);
        $("#code-phpHttpClient pre code").html(reqCodeMap["phpHttpClient"]);
        $("#code-phpClient pre code").html(reqCodeMap["phpClient"]);
        $("#code-pythonClient pre code").html(reqCodeMap["pythonClient"]);
        $("#code-pythonRequests pre code").html(reqCodeMap["pythonRequests"]);
        $("#code-c pre code").html(reqCodeMap["c"]);
        $("#code-CSharp pre code").html(reqCodeMap["c#"]);
        $("#code-ruby pre code").html(reqCodeMap["ruby"]);
        $("#code-go pre code").html(reqCodeMap["go"]);


        $('#pills-tab a:first').tab('show');

    })




}

function jsonToBodyInfoItem(jsonObj, parentKey) {
    var jsonInfo = jsonObj
    for (var key in jsonInfo) {
        var reqField = $('.reqField').map(function () {
            return $(this).val();
        }).get();

        var reqVarType = $('.reqVarType').map(function () {
            return $(this).val();
        }).get();

        var reqExample = $('.reqExample').map(function () {
            return $(this).val();
        }).get();

        var keyType = typeof jsonInfo[key]
        if (keyType === "object") {
            keyType = isArrayOrObject(jsonInfo[key])
        }
        if (keyType === "string" && jsonInfo[key].length > 188) {
            jsonInfo[key] = jsonInfo[key].substring(0, 188) + "..."
        }
        var flag = false
        for (var has in reqField) {
            if (reqField[has] === key) {
                if (reqVarType[has] !== keyType) {
                    $('.reqVarType:eq(' + has + ')').val(keyType);
                }
                if (reqExample[has] !== jsonInfo[key]) {
                    $('.reqExample:eq(' + has + ')').val(jsonInfo[key]);
                }
                flag = true
            }
        }
        var field = key
        if (parentKey !== "") {
            field = parentKey + "." + key
        }
        if (!flag && reqField[0] === "") {
            $(".reqField").first().val(field);
            $(".reqVarType").first().val(keyType);
            $(".reqExample").first().val(jsonInfo[key]);
        } else if (!flag) {
            $("#setReqTable").append(addReqTpl(field, keyType, jsonInfo[key], "", "1"));
        }

        if (keyType === "object") {
            var newParentKey = parentKey
            if (newParentKey === "") {
                newParentKey = key
            } else {
                newParentKey = parentKey + "." + key
            }
            jsonToBodyInfoItem(jsonInfo[key], newParentKey)
        }

        if (keyType === "array") {
            if (isArrayOrObject(jsonInfo[key][0]) === "object") {
                var newParentKey = parentKey
                if (newParentKey === "") {
                    newParentKey = key
                } else {
                    newParentKey = parentKey + "." + key
                }
                jsonToBodyInfoItem(jsonInfo[key][0], newParentKey)
            }
        }
    }
}

function jsonToBodyInfo() {
    var jsonInfo = editorAdd.get();
    jsonToBodyInfoItem(jsonInfo, "")
}

function copyReqCodeJson(reqCode) {
    var txt = $("#"+reqCode+" code").text();
    console.log(txt);
    copyContent(txt);
}

function copyUrl() {
    var url = $('#api-url').text();
    copyContent(url)
    ToastShow("已复制到剪切板");
}