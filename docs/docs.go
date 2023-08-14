// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "zhangxin",
            "email": "xin.zhang@hicom.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/sso/login": {
            "get": {
                "description": "用户登录接口,用于通过用户名密码登录系统",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "用户密码登录",
                "parameters": [
                    {
                        "description": "登录参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "根据名称查询用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "每页数据条数",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "第几页",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.FindByNameAndMobileResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "创建用户",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateUserCommand"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "{\"id\":\"1\", \"msg\": \"create success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                }
            }
        },
        "/user/{id}/password": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "更新用户密码",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新用户密码信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ChangePasswordCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"modify success\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/{id}/role": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "更新用户角色信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新用户角色信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ModifyUserRoleCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"modify success\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/{id}/status": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "更新用户状态",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新用户状态  1:启用  2：禁用",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ModifyUserStatusCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"modify success\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.FindByNameAndMobileResponse": {
            "type": "object",
            "properties": {
                "records": {
                    "description": "记录",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.User"
                    }
                },
                "total": {
                    "description": "总数",
                    "type": "integer"
                }
            }
        },
        "domain.ChangePasswordCommand": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.CreateUserCommand": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/enum.SysRole"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.ModifyUserRoleCommand": {
            "type": "object",
            "properties": {
                "role": {
                    "$ref": "#/definitions/enum.SysRole"
                }
            }
        },
        "domain.ModifyUserStatusCommand": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "integer"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "enable": {
                    "description": "1：启用   2：禁用",
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "last_modified_at": {
                    "type": "string"
                },
                "last_modified_by": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "description": "0:普通用户，1:管理员，2:虚拟用户",
                    "allOf": [
                        {
                            "$ref": "#/definitions/enum.SysRole"
                        }
                    ]
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "enum.SysRole": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "SysRoleGeneralUser",
                "SysRoleAdminUser",
                "SysRoleVirtualUser"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "127.0.0.1",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "运维系统",
	Description:      "运维系统 api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
