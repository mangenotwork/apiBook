package docIE

import (
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"testing"
)

var apiBookYb = `{
    "project": {
        "createDate": "2024-08-20 17:01:20",
        "description": "b33333",
        "name": "b33333"
    },
    "dir": [
        {
            "dirId": "2759422810336952320",
            "dirName": "123"
        },
        {
            "dirId": "default_1ae58da57aeb11f7",
            "dirName": "默认"
        },
        {
            "dirId": "recycleBin_1ae58da57aeb11f7",
            "dirName": "回收站"
        }
    ],
    "doc": [
        {
            "description": "11111",
            "dirId": "2566962039146577920",
            "docId": "11027a1f8f3cc6f1",
            "method": "GET",
            "name": "11111",
            "reqBodyInfo": [
                {
                    "field": "code",
                    "varType": "number",
                    "description": "",
                    "example": "0",
                    "isRequired": 1,
                    "sort": 0,
                    "isOpen": 0
                },
                {
                    "field": "msg",
                    "varType": "string",
                    "description": "",
                    "example": "",
                    "isRequired": 1,
                    "sort": 1,
                    "isOpen": 0
                },
                {
                    "field": "data",
                    "varType": "object",
                    "description": "",
                    "example": "[object Object]",
                    "isRequired": 1,
                    "sort": 2,
                    "isOpen": 0
                },
                {
                    "field": "data.key",
                    "varType": "object",
                    "description": "",
                    "example": "[object Object]",
                    "isRequired": 1,
                    "sort": 3,
                    "isOpen": 0
                },
                {
                    "field": "data.key.aaa",
                    "varType": "number",
                    "description": "",
                    "example": "1243",
                    "isRequired": 1,
                    "sort": 4,
                    "isOpen": 0
                },
                {
                    "field": "data.key.bbb",
                    "varType": "object",
                    "description": "",
                    "example": "[object Object]",
                    "isRequired": 1,
                    "sort": 5,
                    "isOpen": 0
                },
                {
                    "field": "data.key.bbb.ccc",
                    "varType": "string",
                    "description": "",
                    "example": "ccc",
                    "isRequired": 1,
                    "sort": 6,
                    "isOpen": 0
                },
                {
                    "field": "timestamp",
                    "varType": "number",
                    "description": "",
                    "example": "1730339510",
                    "isRequired": 1,
                    "sort": 7,
                    "isOpen": 0
                }
            ],
            "reqBodyJson": "{\"code\":0,\"msg\":\"\",\"data\":{\"key\":{\"aaa\":1243,\"bbb\":{\"ccc\":\"ccc\"}}},\"timestamp\":1730339510}",
            "reqHeader": [
                {
                    "field": "",
                    "varType": "",
                    "description": "",
                    "example": "",
                    "isRequired": 1,
                    "sort": 0,
                    "isOpen": 0
                }
            ],
            "reqType": "json",
            "resp": [
                {
                    "tag": "成功",
                    "respType": "json",
                    "respTypeName": "",
                    "respBody": "{\"code\":0,\"msg\":\"\",\"data\":{\"key\":{\"aaa\":1243,\"bbb\":2222}},\"timestamp\":1730339510}",
                    "respBodyInfo": [
                        {
                            "field": "code",
                            "varType": "number",
                            "description": "",
                            "example": "0",
                            "isRequired": 0,
                            "sort": 0,
                            "isOpen": 0
                        },
                        {
                            "field": "msg",
                            "varType": "string",
                            "description": "",
                            "example": "",
                            "isRequired": 0,
                            "sort": 1,
                            "isOpen": 0
                        },
                        {
                            "field": "data",
                            "varType": "object",
                            "description": "",
                            "example": "[object Object]",
                            "isRequired": 0,
                            "sort": 2,
                            "isOpen": 0
                        },
                        {
                            "field": "data.key",
                            "varType": "object",
                            "description": "",
                            "example": "[object Object]",
                            "isRequired": 0,
                            "sort": 3,
                            "isOpen": 0
                        },
                        {
                            "field": "data.key.aaa",
                            "varType": "number",
                            "description": "",
                            "example": "1243",
                            "isRequired": 0,
                            "sort": 4,
                            "isOpen": 0
                        },
                        {
                            "field": "data.key.bbb",
                            "varType": "number",
                            "description": "",
                            "example": "2222",
                            "isRequired": 0,
                            "sort": 5,
                            "isOpen": 0
                        },
                        {
                            "field": "timestamp",
                            "varType": "number",
                            "description": "",
                            "example": "1730339510",
                            "isRequired": 0,
                            "sort": 6,
                            "isOpen": 0
                        }
                    ]
                }
            ],
            "url": "11111"
        }
    ]
}`

func Test_ApiBookImport(t *testing.T) {
	obj := NewApiBookImport()

	err := obj.analysis(apiBookYb)
	if err != nil {
		t.Error(err)
		return
	}

	project := obj.analysisProject("user", define.ProjectPrivate)
	t.Log(project)

	obj.analysisDoc(project, "user",
		func(project *entity.Project, dir *entity.DocumentDir) {
			t.Log(dir)
		},
		func(project *entity.Project, doc *entity.DocumentContent, dirId string) {
			t.Log(doc)
			t.Log(dirId)
		},
	)

}
