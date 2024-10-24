package docIE

import (
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"testing"
)

var yapiYb = `[
  {
    "index": 0,
    "name": "公共分类",
    "desc": "公共分类",
    "add_time": 1728974187,
    "up_time": 1728974187,
    "list": [
      {
        "query_path": {
          "path": "/pay/channel",
          "params": []
        },
        "edit_uid": 0,
        "status": "undone",
        "type": "static",
        "req_body_is_json_schema": true,
        "res_body_is_json_schema": true,
        "api_opened": false,
        "index": 0,
        "tag": [],
        "_id": 89,
        "method": "POST",
        "title": "a",
        "desc": "",
        "path": "/pay/channel",
        "req_params": [],
        "req_body_form": [],
        "req_headers": [
          {
            "required": "1",
            "_id": "670f7bfff69bbce4cdcb3ccf",
            "name": "Content-Type",
            "value": "application/json"
          },
          {
            "required": "1",
            "_id": "670f7bfff69bbc97a5cb3cce",
            "name": "sign",
            "desc": ""
          },
          {
            "required": "1",
            "_id": "670f7bfff69bbc1543cb3ccd",
            "name": "source",
            "desc": ""
          },
          {
            "required": "1",
            "_id": "670f7bfff69bbc7b21cb3ccc",
            "name": "token",
            "desc": ""
          }
        ],
        "req_query": [],
        "req_body_type": "json",
        "res_body_type": "json",
        "res_body": "{\n  \"type\": \"object\",\n  \"properties\": {\n    \"code\": {\n      \"type\": \"integer\",\n      \"title\": \"中文\",\n      \"description\": \"说明\"\n    },\n    \"msg\": {\n      \"type\": \"string\"\n    },\n    \"data\": {\n      \"type\": \"string\"\n    },\n    \"timestamp\": {\n      \"type\": \"integer\"\n    }\n  },\n  \"required\": [\n    \"code\",\n    \"msg\",\n    \"data\",\n    \"timestamp\"\n  ]\n}",
        "req_body_other": "{\n  \"type\": \"object\",\n  \"properties\": {\n    \"product_select_list\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"object\",\n        \"properties\": {\n          \"product_id\": {\n            \"type\": \"integer\",\n            \"title\": \"中product_id\",\n            \"description\": \"说product_id\"\n          },\n          \"product_type\": {\n            \"type\": \"integer\"\n          },\n          \"select_count\": {\n            \"type\": \"integer\"\n          }\n        },\n        \"required\": [\n          \"product_id\",\n          \"product_type\",\n          \"select_count\"\n        ]\n      },\n      \"title\": \"中文名\",\n      \"description\": \"说明\"\n    }\n  },\n  \"required\": [\n    \"product_select_list\"\n  ]\n}",
        "project_id": 30,
        "catid": 96,
        "uid": 13,
        "add_time": 1729068031,
        "up_time": 1729068031,
        "__v": 0
      },
      {
        "query_path": {
          "path": "/test/test/1",
          "params": []
        },
        "edit_uid": 0,
        "status": "undone",
        "type": "static",
        "req_body_is_json_schema": true,
        "res_body_is_json_schema": true,
        "api_opened": false,
        "index": 0,
        "tag": [],
        "_id": 83,
        "method": "GET",
        "catid": 96,
        "title": "test1",
        "path": "/test/test/1",
        "project_id": 30,
        "req_params": [],
        "res_body_type": "json",
        "uid": 13,
        "add_time": 1728974206,
        "up_time": 1728974324,
        "req_query": [],
        "req_headers": [
          {
            "required": "1",
            "_id": "670e0df4f69bbc6dfacb3cc0",
            "name": "aaaa1",
            "value": "aaa2",
            "example": "aa333",
            "desc": "44444"
          }
        ],
        "req_body_form": [],
        "__v": 0,
        "desc": "<p>备注信息</p>\n",
        "markdown": "备注信息",
        "res_body": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"product_select_list\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"product_id\":{\"type\":\"number\"},\"product_type\":{\"type\":\"number\"},\"select_count\":{\"type\":\"number\"}},\"required\":[\"product_id\",\"product_type\",\"select_count\"]}}}}"
      }
    ]
  },
  {
    "index": 0,
    "name": "宠物",
    "add_time": 1729068031,
    "up_time": 1729068031,
    "list": [
      {
        "query_path": {
          "path": "/pet",
          "params": []
        },
        "edit_uid": 0,
        "status": "undone",
        "type": "static",
        "req_body_is_json_schema": true,
        "res_body_is_json_schema": true,
        "api_opened": false,
        "index": 0,
        "tag": [
          "宠物"
        ],
        "_id": 87,
        "method": "PUT",
        "title": "修改宠物信息",
        "desc": "",
        "path": "/pet",
        "req_params": [],
        "req_body_form": [],
        "req_headers": [
          {
            "required": "1",
            "_id": "670f7bfff69bbc3f16cb3cca",
            "name": "Content-Type",
            "value": "application/json"
          }
        ],
        "req_query": [],
        "req_body_type": "json",
        "res_body_type": "json",
        "res_body": "{\n  \"type\": \"object\",\n  \"properties\": {\n    \"code\": {\n      \"type\": \"integer\"\n    },\n    \"data\": {\n      \"required\": [\n        \"name\",\n        \"photoUrls\",\n        \"id\",\n        \"category\",\n        \"tags\",\n        \"status\"\n      ],\n      \"type\": \"object\",\n      \"properties\": {\n        \"id\": {\n          \"type\": \"integer\",\n          \"format\": \"int64\",\n          \"minimum\": 1,\n          \"description\": \"宠物ID编号\"\n        },\n        \"category\": {\n          \"type\": \"object\",\n          \"properties\": {\n            \"id\": {\n              \"type\": \"integer\",\n              \"format\": \"int64\",\n              \"minimum\": 1,\n              \"description\": \"分组ID编号\"\n            },\n            \"name\": {\n              \"type\": \"string\",\n              \"description\": \"分组名称\"\n            }\n          },\n          \"xml\": {\n            \"name\": \"Category\"\n          },\n          \"$$ref\": \"#/definitions/Category\"\n        },\n        \"name\": {\n          \"type\": \"string\",\n          \"description\": \"名称\",\n          \"examples\": [\n            \"doggie\"\n          ]\n        },\n        \"photoUrls\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"照片URL\"\n        },\n        \"tags\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"object\",\n            \"properties\": {\n              \"id\": {\n                \"type\": \"integer\",\n                \"format\": \"int64\",\n                \"minimum\": 1,\n                \"description\": \"标签ID编号\"\n              },\n              \"name\": {\n                \"type\": \"string\",\n                \"description\": \"标签名称\"\n              }\n            },\n            \"xml\": {\n              \"name\": \"Tag\"\n            },\n            \"$$ref\": \"#/definitions/Tag\"\n          },\n          \"description\": \"标签\"\n        },\n        \"status\": {\n          \"type\": \"string\",\n          \"description\": \"宠物销售状态\",\n          \"enum\": [\n            \"available\",\n            \"pending\",\n            \"sold\"\n          ]\n        }\n      },\n      \"$$ref\": \"#/definitions/Pet\"\n    }\n  },\n  \"required\": [\n    \"code\",\n    \"data\"\n  ]\n}",
        "req_body_other": "{\n  \"type\": \"object\",\n  \"properties\": {}\n}",
        "project_id": 30,
        "catid": 102,
        "uid": 13,
        "add_time": 1729068031,
        "up_time": 1729068031,
        "__v": 0
      },
      {
        "query_path": {
          "path": "/pet/{petId}",
          "params": []
        },
        "edit_uid": 0,
        "status": "undone",
        "type": "var",
        "req_body_is_json_schema": false,
        "res_body_is_json_schema": true,
        "api_opened": false,
        "index": 0,
        "tag": [
          "宠物"
        ],
        "_id": 85,
        "method": "DELETE",
        "title": "删除宠物信息",
        "desc": "",
        "path": "/pet/{petId}",
        "req_params": [
          {
            "_id": "670f7bfff69bbc07efcb3cc5",
            "name": "petId",
            "desc": "Pet id to delete"
          }
        ],
        "req_body_form": [],
        "req_headers": [
          {
            "required": "0",
            "_id": "670f7bfff69bbc0223cb3cc6",
            "name": "api_key",
            "desc": ""
          }
        ],
        "req_query": [],
        "req_body_type": "raw",
        "res_body_type": "json",
        "res_body": "{\n  \"type\": \"object\",\n  \"properties\": {\n    \"code\": {\n      \"type\": \"integer\",\n      \"minimum\": 0,\n      \"maximum\": 0\n    }\n  },\n  \"required\": [\n    \"code\"\n  ]\n}",
        "project_id": 30,
        "catid": 102,
        "uid": 13,
        "add_time": 1729068031,
        "up_time": 1729068031,
        "__v": 0
      },
      {
        "query_path": {
          "path": "/pet",
          "params": []
        },
        "edit_uid": 0,
        "status": "undone",
        "type": "static",
        "req_body_is_json_schema": false,
        "res_body_is_json_schema": true,
        "api_opened": false,
        "index": 0,
        "tag": [
          "宠物"
        ],
        "_id": 86,
        "method": "POST",
        "title": "新建宠物信息",
        "desc": "",
        "path": "/pet",
        "req_params": [],
        "req_body_form": [
          {
            "required": "1",
            "_id": "670f7bfff69bbccc6acb3cc8",
            "name": "name",
            "desc": "宠物名",
            "type": "text"
          },
          {
            "required": "1",
            "_id": "670f7bfff69bbc6ff6cb3cc7",
            "name": "status",
            "desc": "宠物销售状态",
            "type": "text"
          }
        ],
        "req_headers": [
          {
            "required": "1",
            "_id": "670f7bfff69bbc0b6ccb3cc9",
            "name": "Content-Type",
            "value": "application/x-www-form-urlencoded"
          }
        ],
        "req_query": [],
        "req_body_type": "form",
        "res_body_type": "json",
        "res_body": "{\n  \"type\": \"object\",\n  \"properties\": {\n    \"code\": {\n      \"type\": \"integer\",\n      \"minimum\": 0,\n      \"maximum\": 0\n    },\n    \"data\": {\n      \"required\": [\n        \"name\",\n        \"photoUrls\",\n        \"id\",\n        \"category\",\n        \"tags\",\n        \"status\"\n      ],\n      \"type\": \"object\",\n      \"properties\": {\n        \"id\": {\n          \"type\": \"integer\",\n          \"format\": \"int64\",\n          \"minimum\": 1,\n          \"description\": \"宠物ID编号\"\n        },\n        \"category\": {\n          \"type\": \"object\",\n          \"properties\": {\n            \"id\": {\n              \"type\": \"integer\",\n              \"format\": \"int64\",\n              \"minimum\": 1,\n              \"description\": \"分组ID编号\"\n            },\n            \"name\": {\n              \"type\": \"string\",\n              \"description\": \"分组名称\"\n            }\n          },\n          \"xml\": {\n            \"name\": \"Category\"\n          },\n          \"$$ref\": \"#/definitions/Category\"\n        },\n        \"name\": {\n          \"type\": \"string\",\n          \"description\": \"名称\",\n          \"examples\": [\n            \"doggie\"\n          ]\n        },\n        \"photoUrls\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"照片URL\"\n        },\n        \"tags\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"object\",\n            \"properties\": {\n              \"id\": {\n                \"type\": \"integer\",\n                \"format\": \"int64\",\n                \"minimum\": 1,\n                \"description\": \"标签ID编号\"\n              },\n              \"name\": {\n                \"type\": \"string\",\n                \"description\": \"标签名称\"\n              }\n            },\n            \"xml\": {\n              \"name\": \"Tag\"\n            },\n            \"$$ref\": \"#/definitions/Tag\"\n          },\n          \"description\": \"标签\"\n        },\n        \"status\": {\n          \"type\": \"string\",\n          \"description\": \"宠物销售状态\",\n          \"enum\": [\n            \"available\",\n            \"pending\",\n            \"sold\"\n          ]\n        }\n      },\n      \"$$ref\": \"#/definitions/Pet\"\n    }\n  },\n  \"required\": [\n    \"code\",\n    \"data\"\n  ]\n}",
        "project_id": 30,
        "catid": 102,
        "uid": 13,
        "add_time": 1729068031,
        "up_time": 1729068031,
        "__v": 0
      },
      {
        "query_path": {
          "path": "/pet/{petId}",
          "params": []
        },
        "edit_uid": 0,
        "status": "undone",
        "type": "var",
        "req_body_is_json_schema": false,
        "res_body_is_json_schema": true,
        "api_opened": false,
        "index": 0,
        "tag": [
          "宠物",
          "aa"
        ],
        "_id": 84,
        "method": "GET",
        "title": "查询宠物详情",
        "desc": "",
        "path": "/pet/{petId}",
        "req_params": [
          {
            "_id": "670f7bfff69bbc360bcb3cc1",
            "name": "petId",
            "desc": "宠物 ID"
          }
        ],
        "req_body_form": [],
        "req_headers": [
          {
            "required": "0",
            "_id": "670f7bfff69bbc768dcb3cc2",
            "name": "aaaa",
            "desc": ""
          }
        ],
        "req_query": [],
        "req_body_type": "raw",
        "res_body_type": "json",
        "res_body": "{\n  \"type\": \"object\",\n  \"properties\": {\n    \"code\": {\n      \"type\": \"integer\",\n      \"minimum\": 0,\n      \"maximum\": 0,\n      \"description\": \"状态码\"\n    },\n    \"data\": {\n      \"required\": [\n        \"name\",\n        \"photoUrls\",\n        \"id\",\n        \"category\",\n        \"tags\",\n        \"status\"\n      ],\n      \"type\": \"object\",\n      \"properties\": {\n        \"id\": {\n          \"type\": \"integer\",\n          \"format\": \"int64\",\n          \"minimum\": 1,\n          \"description\": \"宠物ID编号\"\n        },\n        \"category\": {\n          \"type\": \"object\",\n          \"properties\": {\n            \"id\": {\n              \"type\": \"integer\",\n              \"format\": \"int64\",\n              \"minimum\": 1,\n              \"description\": \"分组ID编号\"\n            },\n            \"name\": {\n              \"type\": \"string\",\n              \"description\": \"分组名称\"\n            }\n          },\n          \"xml\": {\n            \"name\": \"Category\"\n          },\n          \"$$ref\": \"#/definitions/Category\"\n        },\n        \"name\": {\n          \"type\": \"string\",\n          \"description\": \"名称\",\n          \"examples\": [\n            \"doggie\"\n          ]\n        },\n        \"photoUrls\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"照片URL\"\n        },\n        \"tags\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"object\",\n            \"properties\": {\n              \"id\": {\n                \"type\": \"integer\",\n                \"format\": \"int64\",\n                \"minimum\": 1,\n                \"description\": \"标签ID编号\"\n              },\n              \"name\": {\n                \"type\": \"string\",\n                \"description\": \"标签名称\"\n              }\n            },\n            \"xml\": {\n              \"name\": \"Tag\"\n            },\n            \"$$ref\": \"#/definitions/Tag\"\n          },\n          \"description\": \"标签\"\n        },\n        \"status\": {\n          \"type\": \"string\",\n          \"description\": \"宠物销售状态\",\n          \"enum\": [\n            \"available\",\n            \"pending\",\n            \"sold\"\n          ]\n        }\n      },\n      \"$$ref\": \"#/definitions/Pet\"\n    }\n  },\n  \"required\": [\n    \"code\",\n    \"data\"\n  ]\n}",
        "project_id": 30,
        "catid": 102,
        "uid": 13,
        "add_time": 1729068031,
        "up_time": 1729068031,
        "__v": 0
      },
      {
        "query_path": {
          "path": "/pet/findByStatus",
          "params": []
        },
        "edit_uid": 0,
        "status": "undone",
        "type": "static",
        "req_body_is_json_schema": false,
        "res_body_is_json_schema": true,
        "api_opened": false,
        "index": 0,
        "tag": [
          "宠物"
        ],
        "_id": 88,
        "method": "GET",
        "title": "根据状态查找宠物列表",
        "desc": "",
        "path": "/pet/findByStatus",
        "req_params": [],
        "req_body_form": [],
        "req_headers": [],
        "req_query": [
          {
            "required": "1",
            "_id": "670f7bfff69bbcb0becb3ccb",
            "name": "status",
            "desc": "Status values that need to be considered for filter"
          }
        ],
        "req_body_type": "raw",
        "res_body_type": "json",
        "res_body": "{\n  \"type\": \"array\",\n  \"items\": {\n    \"required\": [\n      \"name\",\n      \"photoUrls\",\n      \"id\",\n      \"category\",\n      \"tags\",\n      \"status\"\n    ],\n    \"type\": \"object\",\n    \"properties\": {\n      \"id\": {\n        \"type\": \"integer\",\n        \"format\": \"int64\",\n        \"minimum\": 1,\n        \"description\": \"宠物ID编号\"\n      },\n      \"category\": {\n        \"type\": \"object\",\n        \"properties\": {\n          \"id\": {\n            \"type\": \"integer\",\n            \"format\": \"int64\",\n            \"minimum\": 1,\n            \"description\": \"分组ID编号\"\n          },\n          \"name\": {\n            \"type\": \"string\",\n            \"description\": \"分组名称\"\n          }\n        },\n        \"xml\": {\n          \"name\": \"Category\"\n        },\n        \"$$ref\": \"#/definitions/Category\"\n      },\n      \"name\": {\n        \"type\": \"string\",\n        \"description\": \"名称\",\n        \"examples\": [\n          \"doggie\"\n        ]\n      },\n      \"photoUrls\": {\n        \"type\": \"array\",\n        \"items\": {\n          \"type\": \"string\"\n        },\n        \"description\": \"照片URL\"\n      },\n      \"tags\": {\n        \"type\": \"array\",\n        \"items\": {\n          \"type\": \"object\",\n          \"properties\": {\n            \"id\": {\n              \"type\": \"integer\",\n              \"format\": \"int64\",\n              \"minimum\": 1,\n              \"description\": \"标签ID编号\"\n            },\n            \"name\": {\n              \"type\": \"string\",\n              \"description\": \"标签名称\"\n            }\n          },\n          \"xml\": {\n            \"name\": \"Tag\"\n          },\n          \"$$ref\": \"#/definitions/Tag\"\n        },\n        \"description\": \"标签\"\n      },\n      \"status\": {\n        \"type\": \"string\",\n        \"description\": \"宠物销售状态\",\n        \"enum\": [\n          \"available\",\n          \"pending\",\n          \"sold\"\n        ]\n      }\n    },\n    \"$$ref\": \"#/definitions/Pet\"\n  }\n}",
        "project_id": 30,
        "catid": 102,
        "uid": 13,
        "add_time": 1729068031,
        "up_time": 1729068031,
        "__v": 0
      }
    ]
  }
]`

func Test_YApiImport(t *testing.T) {
	obj := NewYApiImport()

	err := obj.analysis(yapiYb)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(obj)

	project := obj.analysisProject("user", define.ProjectPrivate)
	t.Log("project: ", project)

	obj.analysisDoc(project, "user", "",
		func(project *entity.Project, dirName string) (string, bool) {
			return "", false
		},
		func(project *entity.Project, doc *entity.DocumentContent, dirId string) {

		},
	)

}
