// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/admin": {
            "get": {
                "description": "Get all admin with data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminJWT"
                ],
                "summary": "Get All admin",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/admin.Admin"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create admin with data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Create admin",
                "parameters": [
                    {
                        "description": "admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.AdminSwagger"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/admin.Admin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/admin/token": {
            "post": {
                "description": "Get token for admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Get Token",
                "parameters": [
                    {
                        "description": "admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.InputAdminToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/admin/{id}": {
            "get": {
                "description": "Get Admin By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Get Admin By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "anything id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.Admin"
                        }
                    }
                }
            },
            "put": {
                "description": "update data admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AdminJWT"
                ],
                "summary": "Update Admin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id admin",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.AdminSwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "delete data admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Delete Admin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id admin",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/covid": {
            "get": {
                "description": "Get All Data Pasien with data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Get All Data Pasien",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/covid.Res_Detail_covid"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create data pasien",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Create data pasien",
                "parameters": [
                    {
                        "description": "covid",
                        "name": "pasien",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/covid.Input_Detail_covid"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/covid.Res_Detail_covid"
                        }
                    }
                }
            }
        },
        "/covid/{id}": {
            "get": {
                "description": "Get Data Pasien By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Get Data Pasien By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "anything id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/covid.Res_Detail_covid"
                        }
                    }
                }
            },
            "put": {
                "description": "Update data pasien",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Update data pasien",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "anything id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "covid",
                        "name": "pasien",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/covid.Input_Detail_covid"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/covid.Res_Detail_covid"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete data pasien",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Delete data pasien",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "anything id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/covid.Res_Detail_covid"
                        }
                    }
                }
            }
        },
        "/hospital": {
            "get": {
                "description": "Get all data rumah sakit with data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Get all data rumah sakit",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/covid.Res_Rumah_sakit"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create data rumah sakit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Create data rumah sakit",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "hospital",
                        "name": "rumah_sakit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/covid.Input_Rumah_sakit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/covid.Res_Rumah_sakit"
                        }
                    }
                }
            }
        },
        "/hospital/{id}": {
            "get": {
                "description": "Get all data rumah sakit with data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Get data rumah sakit by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "anything id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/covid.Res_Rumah_sakit"
                        }
                    }
                }
            },
            "put": {
                "description": "Update data rumah sakit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Update data rumah sakit",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "anything id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "hospital",
                        "name": "rumah_sakit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/covid.Input_Rumah_sakit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/covid.Res_Rumah_sakit"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete data rumah sakit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Delete data rumah sakit",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "anything id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/covid.Res_Rumah_sakit"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.Admin": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.AdminSwagger": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.InputAdminToken": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "covid.Input_Detail_covid": {
            "type": "object",
            "properties": {
                "alamat": {
                    "type": "string"
                },
                "covid": {
                    "type": "boolean"
                },
                "nama": {
                    "type": "string"
                },
                "rumah_sakit_id": {
                    "type": "integer"
                }
            }
        },
        "covid.Input_Rumah_sakit": {
            "type": "object",
            "properties": {
                "alamat": {
                    "type": "string"
                },
                "nama_rumah_sakit": {
                    "type": "string"
                }
            }
        },
        "covid.Res_Detail_covid": {
            "type": "object",
            "properties": {
                "alamat": {
                    "type": "string"
                },
                "covid": {
                    "type": "boolean"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nama": {
                    "type": "string"
                },
                "rumah_sakit": {
                    "$ref": "#/definitions/covid.Res_Rumah_sakit"
                },
                "rumah_sakit_id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "covid.Res_Rumah_sakit": {
            "type": "object",
            "properties": {
                "alamat": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nama_rumah_sakit": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Jasa Pengiriman",
	Description:      "Berikut API Jasa Pengiriman",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}