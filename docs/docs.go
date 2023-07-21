// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://seaffood.com/api",
        "contact": {
            "name": "murasame",
            "url": "https://twitter.com/fresh_salmon256",
            "email": "oogiriminister@gmail.com"
        },
        "license": {
            "name": "No-license",
            "url": "No-license"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/acccounts/{from_user_id}/follow": {
            "post": {
                "description": "Create Follow",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccountsFollow"
                ],
                "summary": "Create Follow",
                "parameters": [
                    {
                        "type": "string",
                        "description": "create Follow Request path",
                        "name": "from_user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "create Follow Request Body",
                        "name": "CreateFollowRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateFollowRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "succsss response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Follows"
                            }
                        }
                    },
                    "400": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove follow",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccountsFollow"
                ],
                "summary": "Remove follow",
                "parameters": [
                    {
                        "type": "string",
                        "description": "remove Follow Request path",
                        "name": "from_user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "remove Follow Request Body",
                        "name": "RemoveFollowRequestQueries",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateFollowRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "succsss response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Follows"
                            }
                        }
                    },
                    "400": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/accounts": {
            "post": {
                "description": "Create new account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Create new account",
                "parameters": [
                    {
                        "description": "Create Account Request Body",
                        "name": "CreateAccountRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateAccountRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "create succsss response",
                        "schema": {
                            "$ref": "#/definitions/api.CreateAccountResponses"
                        }
                    },
                    "400": {
                        "description": "bad request response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/accounts/{user_id}": {
            "get": {
                "description": "Get Any Account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Get account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get success response",
                        "schema": {
                            "$ref": "#/definitions/api.GetAccountResponses"
                        }
                    },
                    "400": {
                        "description": "bad request response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update process when it matches the person",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Update Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Account Request Body",
                        "name": "UpdateAccountRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UpdateAccountRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update succsss response",
                        "schema": {
                            "$ref": "#/definitions/api.UpdateAccountResponse"
                        }
                    },
                    "400": {
                        "description": "bad request response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Only you can delete your account (logical delete)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Remove Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "delete succsss response",
                        "schema": {
                            "$ref": "#/definitions/api.DeleteResponse"
                        }
                    },
                    "400": {
                        "description": "bad request response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/bookmarks": {
            "post": {
                "description": "Create new bookmark",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookmark"
                ],
                "summary": "Create new bookmark",
                "parameters": [
                    {
                        "description": "New Bookmark Request Body",
                        "name": "CreateBookmarkRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateBookmarkRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "create succsss response",
                        "schema": {
                            "$ref": "#/definitions/api.BookmarkResponse"
                        }
                    },
                    "400": {
                        "description": "bad request response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/bookmarks/{hackathon_id}": {
            "get": {
                "description": "Get bookmark",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookmark"
                ],
                "summary": "Get my bookmark",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delete Bookmark Request Body",
                        "name": "ListBookmarkRequestQueries",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "delete succsss response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.BookmarkResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete bookmark",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookmark"
                ],
                "summary": "delete bookmark",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delete Bookmark Request Body",
                        "name": "hackathon_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "delete succsss response",
                        "schema": {
                            "$ref": "#/definitions/api.BookmarkResponse"
                        }
                    },
                    "400": {
                        "description": "bad request response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error response",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.BookmarkResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "expired": {
                    "type": "string"
                },
                "hackathon_id": {
                    "type": "integer"
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
                "term": {
                    "type": "integer"
                }
            }
        },
        "api.CreateAccountRequestBody": {
            "type": "object",
            "required": [
                "locate_id",
                "show_locate",
                "show_rate",
                "user_id",
                "username"
            ],
            "properties": {
                "explanatory_text": {
                    "type": "string"
                },
                "frameworks": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "icon": {
                    "type": "string"
                },
                "locate_id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "show_locate": {
                    "type": "boolean"
                },
                "show_rate": {
                    "type": "boolean"
                },
                "tech_tags": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.CreateAccountResponses": {
            "type": "object",
            "properties": {
                "explanatory_text": {
                    "type": "string"
                },
                "frameworks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.Frameworks"
                    }
                },
                "icon": {
                    "type": "string"
                },
                "locate": {
                    "type": "string"
                },
                "rate": {
                    "type": "integer"
                },
                "show_locate": {
                    "type": "boolean"
                },
                "show_rate": {
                    "type": "boolean"
                },
                "tech_tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.TechTags"
                    }
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.CreateBookmarkRequestBody": {
            "type": "object",
            "properties": {
                "hackathon_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "api.CreateFollowRequestBody": {
            "type": "object",
            "required": [
                "to_user_id"
            ],
            "properties": {
                "to_user_id": {
                    "type": "string"
                }
            }
        },
        "api.DeleteResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                }
            }
        },
        "api.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.GetAccountResponses": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "explanatory_text": {
                    "type": "string"
                },
                "frameworks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.Frameworks"
                    }
                },
                "icon": {
                    "type": "string"
                },
                "locate": {
                    "type": "string"
                },
                "rate": {
                    "type": "integer"
                },
                "show_locate": {
                    "type": "boolean"
                },
                "show_rate": {
                    "type": "boolean"
                },
                "tech_tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.TechTags"
                    }
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.UpdateAccountRequestBody": {
            "type": "object",
            "properties": {
                "explanatory_text": {
                    "type": "string"
                },
                "hashed_password": {
                    "type": "string"
                },
                "locate_id": {
                    "type": "integer"
                },
                "rate": {
                    "type": "integer"
                },
                "show_locate": {
                    "type": "boolean"
                },
                "show_rate": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.UpdateAccountResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "explanatory_text": {
                    "type": "string"
                },
                "hashed_password": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "locate": {
                    "type": "string"
                },
                "rate": {
                    "type": "integer"
                },
                "show_locate": {
                    "type": "boolean"
                },
                "show_rate": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "db.Follows": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string"
                },
                "from_user_id": {
                    "type": "string"
                },
                "to_user_id": {
                    "type": "string"
                }
            }
        },
        "db.Frameworks": {
            "type": "object",
            "properties": {
                "framework": {
                    "type": "string"
                },
                "framework_id": {
                    "type": "integer"
                },
                "tech_tag_id": {
                    "type": "integer"
                }
            }
        },
        "db.TechTags": {
            "type": "object",
            "properties": {
                "language": {
                    "type": "string"
                },
                "tech_tag_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "https://seaffood.com",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Hack Hack Backend API",
	Description:      "HackPortal Backend API serice",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
