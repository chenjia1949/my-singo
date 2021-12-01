// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ping": {
            "get": {
                "description": "系统状态检查",
                "tags": [
                    "系统探活接口"
                ],
                "summary": "状态检查页面",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登陆接口",
                "tags": [
                    "用户接口"
                ],
                "summary": "用户登陆",
                "parameters": [
                    {
                        "type": "string",
                        "description": "姓名",
                        "name": "user_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/logout": {
            "delete": {
                "description": "用户登出接口",
                "tags": [
                    "用户接口"
                ],
                "summary": "用户登出",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "description": "用户详情接口",
                "tags": [
                    "用户接口"
                ],
                "summary": "用户详情",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册接口",
                "tags": [
                    "用户接口"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "nickname",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "姓名",
                        "name": "user_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "再次确认密码",
                        "name": "password_confirm",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "陈佳实验项目",
	Description: "陈佳实验项目接口文档",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}