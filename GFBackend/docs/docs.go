// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/fongziyjun16/SE/tree/backend",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/register": {
            "post": {
                "description": "get strings by username \u0026 password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Manage"
                ],
                "summary": "Register a new User",
                "parameters": [
                    {
                        "description": "Regular User Register only needs Username, Password(encoded by md5) \u0026 ForAdmin with false. Admin User Register needs Username, Password(encoded by md5) \u0026 ForAdmin with true.",
                        "name": "UserInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UserInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. User Register Successfully",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "406": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. User Has Existed",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "example": "process successfully"
                }
            }
        },
        "controller.UserInfo": {
            "type": "object",
            "properties": {
                "ForAdmin": {
                    "type": "boolean"
                },
                "NewPassword": {
                    "type": "string",
                    "example": "3ecb441b741bcd433288f5557e4b9118"
                },
                "Password": {
                    "type": "string",
                    "example": "f9ae5f68ae1e7f7f3fc06053e9b9b539"
                },
                "Username": {
                    "type": "string",
                    "example": "yingjiechen21"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "http://167.71.166.120:10010",
	BasePath:         "/gf/api",
	Schemes:          []string{},
	Title:            "Gator Forum Backend API",
	Description:      "This is the Gator Forum Backend Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
