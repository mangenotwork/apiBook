package docIE

import (
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"testing"
)

var apizzayb = `{
  "project_info": {
    "id": "0f42fac5e7d327a16c6c555f52eccb04",
    "user_id": "101551",
    "create_time": "2024-10-11 14:44:03",
    "name": "test",
    "comment": "test"
  },
  "categorys": [
    {
      "id": "0c2f5dcc81e117d36deb28bce69ef1bd",
      "name": "aaaaaa",
      "comment": "",
      "rank": 2,
      "parent_category_id": "",
      "api_list": [
        {
          "id": "25e37f53cdf642c18de1bd8c6efec3d2",
          "name": "aaaaaaa",
          "method": "POST",
          "url": "http://aaaaaaaaa",
          "type": "http",
          "header_params": [],
          "query_params": [],
          "body_params": [],
          "body_raw": "",
          "body_raw_example": "aaaaa",
          "raw_content_type": "JSON (application/json)",
          "cookie_params": [],
          "response_doc": "",
          "response_example": "aaaa",
          "response_example_params": [],
          "response_example_annotation": "",
          "markdown_content": null,
          "body_type": "raw",
          "progress": "developing",
          "rank": 1,
          "category_id": "0c2f5dcc81e117d36deb28bce69ef1bd",
          "is_post_type": "1",
          "mock": {
            "mock_model": ""
          },
          "test_model": ""
        }
      ],
      "sub_categorys": [
        {
          "id": "f2354aeabc41b6597e39a842b8aa2ea3",
          "name": "bbbbbb",
          "comment": "",
          "rank": 1,
          "parent_category_id": "0c2f5dcc81e117d36deb28bce69ef1bd",
          "api_list": [
            {
              "id": "0733e56430c4453e86c049b9d8a0e41e",
              "name": "bbbbbb",
              "method": "POST",
              "url": "http://bbbbb",
              "type": "http",
              "header_params": [],
              "query_params": [],
              "body_params": [],
              "body_raw": "",
              "body_raw_example": "",
              "raw_content_type": "JSON (application/json)",
              "cookie_params": [],
              "response_doc": "<p>bbbbbbbbb</p>",
              "response_example": "",
              "response_example_params": [],
              "response_example_annotation": "",
              "markdown_content": null,
              "body_type": "raw",
              "progress": "developing",
              "rank": 1,
              "category_id": "f2354aeabc41b6597e39a842b8aa2ea3",
              "is_post_type": "1",
              "mock": {
                "mock_model": ""
              },
              "test_model": ""
            }
          ]
        }
      ]
    },
    {
      "id": "b53ce4d80bae67d0714c4c2c4edb2d20",
      "name": "test",
      "comment": "",
      "rank": 3,
      "parent_category_id": "",
      "api_list": [
        {
          "id": "f39a746392f244df9499b8c3fb92f953",
          "name": "test",
          "method": "POST",
          "url": "http://test.test/aaaaa",
          "type": "http",
          "header_params": [
            {
              "id": "65c6a195f83c4234807dac4abd7164a1",
              "checked": 1,
              "key": "aaaa",
              "value": "",
              "desc": "aaaa",
              "type": "string",
              "eg": "aaaa",
              "require": 1,
              "isBindingGlobalDoc": 0,
              "enums": []
            }
          ],
          "query_params": [],
          "body_params": [
            {
              "id": "ac911d2eec6f4951aa30a6c862597aae",
              "checked": 1,
              "key": "product_select_list",
              "value": "",
              "desc": "aaaaaaa",
              "type": "array",
              "rtype": "Text",
              "eg": "",
              "require": 1,
              "isBindingGlobalDoc": 0,
              "enums": []
            },
            {
              "id": "bc0a406a92fc4d83b18c083fdcac27c1",
              "checked": 1,
              "key": "product_select_list.product_id",
              "value": "",
              "desc": "bbbbb",
              "type": "number",
              "rtype": "Text",
              "eg": "",
              "require": 1,
              "isBindingGlobalDoc": 0,
              "enums": []
            },
            {
              "id": "ca885254c00b4ccdb167549abe76004f",
              "checked": 1,
              "key": "product_select_list.product_type",
              "value": "",
              "desc": "ccccc",
              "type": "number",
              "rtype": "Text",
              "eg": "",
              "require": 1,
              "isBindingGlobalDoc": 0,
              "enums": []
            },
            {
              "id": "6a8f124190234faebefe067e46a69658",
              "checked": 1,
              "key": "product_select_list.select_count",
              "value": "",
              "desc": "dddddddd",
              "type": "number",
              "rtype": "Text",
              "eg": "",
              "require": 1,
              "isBindingGlobalDoc": 0,
              "enums": []
            }
          ],
          "body_raw": "",
          "body_raw_example": "{\r\n    \"product_select_list\": [\r\n        {\r\n            \"product_id\": 1310,\r\n            \"product_type\": 3,\r\n            \"select_count\": 1\r\n        },\r\n        {\r\n            \"product_id\": 1311,\r\n            \"product_type\": 3,\r\n            \"select_count\": 1\r\n        }\r\n    ]\r\n}",
          "raw_content_type": "JSON (application/json)",
          "cookie_params": [],
          "response_doc": "<p>test11111111111</p>",
          "response_example": "{\r\n  \"code\": 1001,\r\n  \"msg\": \"token invalid\",\r\n  \"data\": \"\",\r\n  \"timestamp\": 1727334419\r\n}",
          "response_example_params": [
            {
              "id": "7999c068448c4b9e8ef19838bfa1c207",
              "key": "code",
              "desc": "1111111",
              "type": "number",
              "require": 1
            },
            {
              "id": "cb5efd32b2c140b8af03bcdfa6b867b5",
              "key": "msg",
              "desc": "2222222",
              "type": "string",
              "require": 1
            },
            {
              "id": "7e7c40c8df8a4ef9bbe0fedba10ba127",
              "key": "data",
              "desc": "33333333",
              "type": "string",
              "require": 1
            },
            {
              "id": "3face451e494462194d0a208cfa7eba3",
              "key": "timestamp",
              "desc": "4444444",
              "type": "number",
              "require": 1
            }
          ],
          "response_example_annotation": "",
          "markdown_content": null,
          "body_type": "raw",
          "progress": "developing",
          "rank": 1,
          "category_id": "b53ce4d80bae67d0714c4c2c4edb2d20",
          "is_post_type": "1",
          "mock": {
            "mock_model": ""
          },
          "test_model": ""
        }
      ],
      "sub_categorys": []
    }
  ],
  "envirnments": [],
  "status_codes": null,
  "request_paramdocs": [],
  "response_models": [],
  "tester_flows": [],
  "version": "1.0"
}`

func Test_ApiZZAImport(t *testing.T) {
	obj := NewApiZZAImport()

	err := obj.analysis(apizzayb)
	if err != nil {
		t.Error(err)
		return
	}

	project := obj.analysisProject("user", define.ProjectPrivate)
	t.Log("project : ", project)

	obj.analysisDoc(project, "user", "",
		func(project *entity.Project, dirName string) (string, bool) {
			t.Log("dirName = ", dirName)
			return "", false
		},
		func(project *entity.Project, doc *entity.DocumentContent, dirId string) {
			t.Log("doc = ", doc)
			t.Log("================================\n\n")
		})

}
