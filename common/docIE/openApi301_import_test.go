package docIE

import (
	"apiBook/internal/define"
	"apiBook/internal/entity"
	"testing"
)

var yb = `{
  "openapi": "3.0.1",
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
            "example": "1",
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "aaaa",
            "in": "header",
            "description": "",
            "required": false,
            "example": "aaaaaa",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
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
                      "$ref": "#/components/schemas/Pet",
                      "description": "宠物信息"
                    }
                  },
                  "required": [
                    "code",
                    "data"
                  ]
                },
                "examples": {
                  "1": {
                    "summary": "成功示例",
                    "value": {
                      "code": 0,
                      "data": {
                        "name": "Hello Kity",
                        "photoUrls": [
                          "http://dummyimage.com/400x400"
                        ],
                        "id": 3,
                        "category": {
                          "id": 71,
                          "name": "Cat"
                        },
                        "tags": [
                          {
                            "id": 22,
                            "name": "Cat"
                          }
                        ],
                        "status": "sold"
                      }
                    }
                  }
                }
              }
            },
            "headers": {}
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
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
            "headers": {}
          },
          "404": {
            "description": "",
            "content": {
              "application/json": {
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
            "headers": {}
          }
        },
        "security": []
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
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "api_key",
            "in": "header",
            "description": "",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
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
                },
                "examples": {
                  "1": {
                    "summary": "成功示例",
                    "value": {
                      "code": 0
                    }
                  }
                }
              }
            },
            "headers": {}
          }
        },
        "security": []
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
        "parameters": [],
        "requestBody": {
          "content": {
            "application/x-www-form-urlencoded": {
              "schema": {
                "type": "object",
                "properties": {
                  "name": {
                    "description": "宠物名",
                    "example": "Hello Kitty",
                    "type": "string"
                  },
                  "status": {
                    "description": "宠物销售状态",
                    "example": "sold",
                    "type": "string"
                  }
                },
                "required": [
                  "name",
                  "status"
                ]
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "code": {
                      "type": "integer",
                      "minimum": 0,
                      "maximum": 0
                    },
                    "data": {
                      "$ref": "#/components/schemas/Pet",
                      "description": "宠物信息"
                    }
                  },
                  "required": [
                    "code",
                    "data"
                  ]
                },
                "examples": {
                  "1": {
                    "summary": "成功示例",
                    "value": {
                      "code": 0,
                      "data": {
                        "name": "Hello Kity",
                        "photoUrls": [
                          "http://dummyimage.com/400x400"
                        ],
                        "id": 3,
                        "category": {
                          "id": 71,
                          "name": "Cat"
                        },
                        "tags": [
                          {
                            "id": 22,
                            "name": "Cat"
                          }
                        ],
                        "status": "sold"
                      }
                    }
                  }
                }
              }
            },
            "headers": {}
          }
        },
        "security": []
      },
      "put": {
        "summary": "修改宠物信息",
        "deprecated": false,
        "description": "",
        "tags": [
          "宠物"
        ],
        "parameters": [],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {}
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "code": {
                      "type": "integer"
                    },
                    "data": {
                      "$ref": "#/components/schemas/Pet",
                      "description": "宠物信息"
                    }
                  },
                  "required": [
                    "code",
                    "data"
                  ]
                },
                "examples": {
                  "1": {
                    "summary": "成功示例",
                    "value": {
                      "code": 0,
                      "data": {
                        "name": "Hello Kity",
                        "photoUrls": [
                          "http://dummyimage.com/400x400"
                        ],
                        "id": 3,
                        "category": {
                          "id": 71,
                          "name": "Cat"
                        },
                        "tags": [
                          {
                            "id": 22,
                            "name": "Cat"
                          }
                        ],
                        "status": "sold"
                      }
                    }
                  }
                }
              }
            },
            "headers": {}
          },
          "404": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {}
                }
              }
            },
            "headers": {}
          },
          "405": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {}
                }
              }
            },
            "headers": {}
          }
        },
        "security": []
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
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/Pet",
                    "description": "宠物信息"
                  }
                },
                "examples": {
                  "1": {
                    "summary": "成功示例",
                    "value": {
                      "code": 0,
                      "data": [
                        {
                          "name": "Hello Kity",
                          "photoUrls": [
                            "http://dummyimage.com/400x400"
                          ],
                          "id": 3,
                          "category": {
                            "id": 71,
                            "name": "Cat"
                          },
                          "tags": [
                            {
                              "id": 22,
                              "name": "Cat"
                            }
                          ],
                          "status": "sold"
                        },
                        {
                          "name": "White Dog",
                          "photoUrls": [
                            "http://dummyimage.com/400x400"
                          ],
                          "id": 3,
                          "category": {
                            "id": 71,
                            "name": "Dog"
                          },
                          "tags": [
                            {
                              "id": 22,
                              "name": "Dog"
                            }
                          ],
                          "status": "sold"
                        }
                      ]
                    }
                  }
                }
              }
            },
            "headers": {}
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
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
            "headers": {}
          }
        },
        "security": []
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
            "example": "/<B70o;7W@3W,]dG<20q",
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "source",
            "in": "header",
            "description": "",
            "required": true,
            "example": "3",
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "token",
            "in": "header",
            "description": "",
            "required": true,
            "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaWQiOjExMzIsImNsaWVudF9pZCI6IjU2ZDhmNzNjLTkzNzktNDc2OC05ODZhLWY2OTEwMWVkMzA1NiIsImV4cCI6MTcyNzUwNDE3OCwidWlkIjoxMjQ4fQ.bf6nIeqBhSPpfoiRJcf9nT1dMreZgR98kUbJGSqiZIg",
            "schema": {
              "type": "string"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
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
              },
              "example": {
                "product_select_list": [
                  {
                    "product_id": 1310,
                    "product_type": 3,
                    "select_count": 1
                  },
                  {
                    "product_id": 1311,
                    "product_type": 3,
                    "select_count": 1
                  }
                ]
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
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
                },
                "examples": {
                  "1": {
                    "summary": "成功示例",
                    "value": {
                      "code": 1001,
                      "msg": "token invalid",
                      "data": "",
                      "timestamp": 1727334419
                    }
                  }
                }
              }
            },
            "headers": {}
          }
        },
        "security": []
      }
    }
  },
  "components": {
    "schemas": {
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
            "$ref": "#/components/schemas/Category",
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
              "$ref": "#/components/schemas/Tag"
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
    "securitySchemes": {}
  },
  "servers": []
}`

func Test_OpenApi301(t *testing.T) {
	obj := NewOpenApi301()

	err := obj.analysis(yb)
	if err != nil {
		t.Error(err)
		return
	}

	project := obj.getProject("user", define.ProjectPrivate)
	t.Log(project)

	dir := "dirdirdir"

	obj.analysisDoc(project, "user",
		func(doc *entity.DocumentContent) {
			t.Log(doc)
			t.Log(dir)
		},
	)

}
