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
        "/community/create": {
            "get": {
                "description": "need strings community name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Community Manage"
                ],
                "summary": "Get the Community by Name",
                "parameters": [
                    {
                        "description": "Create a new community needs Creator, Name \u0026 Description.",
                        "name": "CommunityInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CommunityInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. Create Community Success",
                        "schema": {
                            "$ref": "#/definitions/controller.CommunityResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters or Community already exists",
                        "schema": {
                            "$ref": "#/definitions/controller.CommunityResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.CommunityResponseMsg"
                        }
                    }
                }
            },
            "post": {
                "description": "need strings creator \u0026 community name \u0026 description \u0026 create time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Community Manage"
                ],
                "summary": "Create a new Community",
                "parameters": [
                    {
                        "description": "Create a new community needs Creator, Name \u0026 Description.",
                        "name": "CommunityInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CommunityInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. Create Community Success",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters or Community already exists",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        },
        "/file/scan": {
            "post": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "description": "need token in cookie, only get self files",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Static Resource"
                ],
                "summary": "Scan User files",
                "responses": {
                    "201": {
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. Scan Successfully",
                        "schema": {
                            "$ref": "#/definitions/controller.UserFiles"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        },
        "/file/upload": {
            "post": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "description": "need token in cookie, html file type input element include name attribute with value \"uploadFilename\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Static Resource"
                ],
                "summary": "User Uploads files including images, video etc.",
                "responses": {
                    "201": {
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. Upload Successfully",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters or No Enough Space",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        },
        "/resources/userfiles/{username}/{filename}": {
            "get": {
                "description": "Static files request, need to claim the username and filename in the url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Static Resource"
                ],
                "summary": "Request User Files",
                "responses": {}
            }
        },
        "/user/admin/delete": {
            "post": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "description": "need strings username in post request, need token in cookie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Manage"
                ],
                "summary": "Admin delete Users, cannot self delete",
                "parameters": [
                    {
                        "description": "username in post request body",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. Update Password Successfully",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        },
        "/user/admin/register": {
            "post": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "description": "only need strings username \u0026 password \u0026 ForAdmin, need token in cookie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Manage"
                ],
                "summary": "Register a new Admin User",
                "parameters": [
                    {
                        "description": "Admin User Register only needs Username, Password(encoded by md5) \u0026 ForAdmin with true.",
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
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters or User Has Existed",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        },
        "/user/follow": {
            "post": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "description": "need token in cookie, need username who is followed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Manage"
                ],
                "summary": "User Follow other users",
                "parameters": [
                    {
                        "description": "username in post request body",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. Follow Successfully",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters or User not exist or User has followed.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "only need strings username \u0026 password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Manage"
                ],
                "summary": "Admin / Regular User login",
                "parameters": [
                    {
                        "description": "only needs username and password",
                        "name": "UserInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. User Login Successfully",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters or Username / Password incorrect",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "description": "need strings username in post request, need token in cookie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Manage"
                ],
                "summary": "Admin / Regular User logout",
                "parameters": [
                    {
                        "description": "username in post request body",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/password": {
            "post": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "description": "need token in cookie, need Username, Password, NewPassword",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Manage"
                ],
                "summary": "Admin \u0026 Regular Update Password",
                "parameters": [
                    {
                        "description": "need Username, Password, NewPassword",
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
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. Update Password Successfully",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters or Password not match",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "only need strings username \u0026 password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Manage"
                ],
                "summary": "Register a new Regular User",
                "parameters": [
                    {
                        "description": "Regular User Register only needs Username, Password(encoded by md5) \u0026 ForAdmin with false.",
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
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters or User Has Existed",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        },
        "/user/unfollow": {
            "post": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "description": "need token in cookie, need username who is followed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Manage"
                ],
                "summary": "User Unfollow other users",
                "parameters": [
                    {
                        "description": "username in post request body",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. Unfollow Successfully",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters or User not exist.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "post": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "description": "need token in cookie, need Nickname, Birthday(yyyy-mm-dd), Gender(male / female / unknown), Department",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Manage"
                ],
                "summary": "Update user information including Nickname, Birthday(yyyy-mm-dd), Gender(male / female / unknown), Department",
                "parameters": [
                    {
                        "description": "need Nickname, Birthday(yyyy-mm-dd), Gender(male / female / unknown), Department",
                        "name": "NewUserInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.NewUserInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "\u003cb\u003eSuccess\u003c/b\u003e. Update Password Successfully",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Bad Parameters",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "\u003cb\u003eFailure\u003c/b\u003e. Server Internal Error.",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseMsg"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CommunityInfo": {
            "type": "object",
            "properties": {
                "Create_Time": {
                    "type": "string",
                    "example": "2020-01-01"
                },
                "Creator": {
                    "type": "string",
                    "example": "test1"
                },
                "Description": {
                    "type": "string",
                    "example": "this is a test community"
                },
                "Name": {
                    "type": "string",
                    "example": "community1"
                }
            }
        },
        "controller.CommunityResponseMsg": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "$ref": "#/definitions/model.Community"
                },
                "message": {
                    "type": "string",
                    "example": "process successfully"
                }
            }
        },
        "controller.NewUserInfo": {
            "type": "object",
            "properties": {
                "Birthday": {
                    "type": "string",
                    "example": "2022-02-30"
                },
                "Department": {
                    "type": "string",
                    "example": "CS:GO"
                },
                "Gender": {
                    "type": "string",
                    "example": "male/female/unknown"
                },
                "Nickname": {
                    "type": "string",
                    "example": "Peter Park"
                },
                "Username": {
                    "type": "string",
                    "example": "jamesbond21"
                }
            }
        },
        "controller.ResponseMsg": {
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
        "controller.UserFiles": {
            "type": "object",
            "properties": {
                "Filenames": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "\"xxx.jpg\"",
                        "\"xxx.png\"",
                        "\"xxx.gif\""
                    ]
                },
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
                    "example": "jamesbond21"
                }
            }
        },
        "model.Community": {
            "type": "object",
            "properties": {
                "create_Time": {
                    "type": "string",
                    "example": "2020-01-01"
                },
                "creator": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "num_Member": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiAuthToken": {
            "type": "apiKey",
            "name": "token",
            "in": "cookies"
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
