// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://hc-dev.seafood-dev.com",
        "contact": {
            "name": "murasame29",
            "url": "https://twitter.com/fresh_salmon256",
            "email": "oogiriminister@gmail.com"
        },
        "license": {
            "name": "No-license"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hackathons": {
            "get": {
                "description": "List Hackathons",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hackathon"
                ],
                "summary": "List Hackathons",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "期間が長いかどうか？",
                        "name": "longTerm",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "締め切りが近いかどうか？",
                        "name": "nearDeadline",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "新着かどうか？",
                        "name": "new",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "description": "タグ",
                        "name": "tags",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.GetHackathon"
                            }
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            },
            "post": {
                "description": "Create Hackathon",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hackathon"
                ],
                "summary": "Create Hackathon",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "CreateHackathonRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateHackathon"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.CreateHackathon"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        },
        "/hackathons/{hackathon_id}": {
            "delete": {
                "description": "Delete Hackathons",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hackathon"
                ],
                "summary": "Delete Hackathons",
                "parameters": [
                    {
                        "type": "string",
                        "description": "request body",
                        "name": "hackathon_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.DeleteHackathon"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        },
        "/status_tags": {
            "get": {
                "description": "Get all StatusTag",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "StatusTag"
                ],
                "summary": "Get all StatusTag",
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.StatusTag"
                            }
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            },
            "post": {
                "description": "Create a new StatusTag",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "StatusTag"
                ],
                "summary": "Create a new StatusTag",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "CreateStatusTagRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateStatusTag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.StatusTag"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        },
        "/status_tags/{id}": {
            "put": {
                "description": "Update StatusTag by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "StatusTag"
                ],
                "summary": "Update StatusTag by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "status tag id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "CreateStatusTagRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateStatusTag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/response.StatusTag"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateHackathon": {
            "type": "object",
            "required": [
                "expired",
                "link",
                "name",
                "start_date",
                "term"
            ],
            "properties": {
                "expired": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "statuses[]": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "term": {
                    "type": "integer"
                }
            }
        },
        "request.CreateStatusTag": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "request.UpdateStatusTag": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "response.CreateHackathon": {
            "type": "object",
            "properties": {
                "expired": {
                    "type": "string"
                },
                "hackathon_id": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "status_tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.StatusTag"
                    }
                },
                "term": {
                    "type": "integer"
                }
            }
        },
        "response.DeleteHackathon": {
            "type": "object"
        },
        "response.GetHackathon": {
            "type": "object",
            "properties": {
                "expired": {
                    "type": "string"
                },
                "hackathon_id": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "status_tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.StatusTag"
                    }
                },
                "term": {
                    "type": "integer"
                }
            }
        },
        "response.StatusTag": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "api-dev.hack-portal.com",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Hack-Portal Backend API",
	Description:      "Hack-Portal Backend API serice",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
