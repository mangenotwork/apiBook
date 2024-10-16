package docIE

import (
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"testing"
)

var swaggerYb = `{
  "info": {
    "title": "个人项目",
    "description": "",
    "version": "1.0.0"
  },
  "tags": [
    {
      "name": "宠物"
    },
    {
      "name": "aa"
    }
  ],
  "paths": {
    "/pet/{petId}": {
      "get": {
        "summary": "查询宠物详情",
        "deprecated": false,
        "description": "",
        "tags": [
          "宠物",
          "aa"
        ],
        "parameters": [
          {
            "name": "petId",
            "in": "path",
            "description": "宠物 ID",
            "required": true,
            "type": "string",
            "x-example": "1"
          },
          {
            "name": "aaaa",
            "in": "header",
            "description": "",
            "required": false,
            "type": "string",
            "x-example": "aaaaaa"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "type": "integer",
                  "minimum": 0,
                  "maximum": 0,
                  "description": "状态码"
                },
                "data": {
                  "$ref": "#/definitions/Pet",
                  "description": "宠物信息"
                }
              },
              "required": [
                "code",
                "data"
              ]
            }
          },
          "400": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "type": "integer"
                },
                "message": {
                  "type": "string"
                }
              },
              "required": [
                "code",
                "message"
              ]
            }
          },
          "404": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "type": "integer"
                },
                "message": {
                  "type": "string"
                }
              },
              "required": [
                "code",
                "message"
              ]
            }
          }
        },
        "security": [],
        "produces": [
          "application/json"
        ]
      },
      "delete": {
        "summary": "删除宠物信息",
        "deprecated": false,
        "description": "",
        "tags": [
          "宠物"
        ],
        "parameters": [
          {
            "name": "petId",
            "in": "path",
            "description": "Pet id to delete",
            "required": true,
            "type": "string"
          },
          {
            "name": "api_key",
            "in": "header",
            "description": "",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "type": "integer",
                  "minimum": 0,
                  "maximum": 0
                }
              },
              "required": [
                "code"
              ]
            }
          }
        },
        "security": [],
        "produces": [
          "application/json"
        ]
      }
    },
    "/pet": {
      "post": {
        "summary": "新建宠物信息",
        "deprecated": false,
        "description": "",
        "tags": [
          "宠物"
        ],
        "parameters": [
          {
            "name": "name",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": "宠物名"
          },
          {
            "name": "status",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": "宠物销售状态"
          }
        ],
        "responses": {
          "201": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "type": "integer",
                  "minimum": 0,
                  "maximum": 0
                },
                "data": {
                  "$ref": "#/definitions/Pet",
                  "description": "宠物信息"
                }
              },
              "required": [
                "code",
                "data"
              ]
            }
          }
        },
        "security": [],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "produces": [
          "application/json"
        ]
      },
      "put": {
        "summary": "修改宠物信息",
        "deprecated": false,
        "description": "",
        "tags": [
          "宠物"
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {}
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "type": "integer"
                },
                "data": {
                  "$ref": "#/definitions/Pet",
                  "description": "宠物信息"
                }
              },
              "required": [
                "code",
                "data"
              ]
            }
          },
          "404": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "object",
              "properties": {}
            }
          },
          "405": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "object",
              "properties": {}
            }
          }
        },
        "security": [],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ]
      }
    },
    "/pet/findByStatus": {
      "get": {
        "summary": "根据状态查找宠物列表",
        "deprecated": false,
        "description": "",
        "tags": [
          "宠物"
        ],
        "parameters": [
          {
            "name": "status",
            "in": "query",
            "description": "Status values that need to be considered for filter",
            "required": true,
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet",
                "description": "宠物信息"
              }
            }
          },
          "400": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "type": "integer"
                }
              },
              "required": [
                "code"
              ]
            }
          }
        },
        "security": [],
        "produces": [
          "application/json"
        ]
      }
    },
    "/pay/channel": {
      "post": {
        "summary": "a",
        "deprecated": false,
        "description": "",
        "tags": [],
        "parameters": [
          {
            "name": "sign",
            "in": "header",
            "description": "",
            "required": true,
            "type": "string",
            "x-example": "/<B70o;7W@3W,]dG<20q"
          },
          {
            "name": "source",
            "in": "header",
            "description": "",
            "required": true,
            "type": "string",
            "x-example": "3"
          },
          {
            "name": "token",
            "in": "header",
            "description": "",
            "required": true,
            "type": "string",
            "x-example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaWQiOjExMzIsImNsaWVudF9pZCI6IjU2ZDhmNzNjLTkzNzktNDc2OC05ODZhLWY2OTEwMWVkMzA1NiIsImV4cCI6MTcyNzUwNDE3OCwidWlkIjoxMjQ4fQ.bf6nIeqBhSPpfoiRJcf9nT1dMreZgR98kUbJGSqiZIg"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "product_select_list": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "product_id": {
                        "type": "integer",
                        "title": "中product_id",
                        "description": "说product_id"
                      },
                      "product_type": {
                        "type": "integer"
                      },
                      "select_count": {
                        "type": "integer"
                      }
                    },
                    "required": [
                      "product_id",
                      "product_type",
                      "select_count"
                    ]
                  },
                  "title": "中文名",
                  "description": "说明"
                }
              },
              "required": [
                "product_select_list"
              ]
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {},
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "type": "integer",
                  "title": "中文",
                  "description": "说明"
                },
                "msg": {
                  "type": "string"
                },
                "data": {
                  "type": "string"
                },
                "timestamp": {
                  "type": "integer"
                }
              },
              "required": [
                "code",
                "msg",
                "data",
                "timestamp"
              ]
            }
          }
        },
        "security": [],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ]
      }
    }
  },
  "swagger": "2.0",
  "definitions": {
    "Pet": {
      "required": [
        "name",
        "photoUrls",
        "id",
        "category",
        "tags",
        "status"
      ],
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "description": "宠物ID编号"
        },
        "category": {
          "$ref": "#/definitions/Category",
          "description": "分组"
        },
        "name": {
          "type": "string",
          "description": "名称",
          "examples": [
            "doggie"
          ]
        },
        "photoUrls": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "照片URL"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          },
          "description": "标签"
        },
        "status": {
          "type": "string",
          "description": "宠物销售状态",
          "enum": [
            "available",
            "pending",
            "sold"
          ]
        }
      }
    },
    "Category": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "description": "分组ID编号"
        },
        "name": {
          "type": "string",
          "description": "分组名称"
        }
      },
      "xml": {
        "name": "Category"
      }
    },
    "Tag": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "description": "标签ID编号"
        },
        "name": {
          "type": "string",
          "description": "标签名称"
        }
      },
      "xml": {
        "name": "Tag"
      }
    }
  },
  "securityDefinitions": {},
  "x-components": {}
}`

func Test_SwaggerImport(t *testing.T) {
	obj := NewSwaggerImport()

	err := obj.analysis(swaggerYb)
	if err != nil {
		t.Error(err)
		return
	}

	// t.Log(obj)

	project := obj.analysisProject("user", define.ProjectPrivate)
	t.Log("project: ", project)

	obj.analysisDoc(project, "user", func(doc *entity.DocumentContent) {
		t.Log("")
	})

}
