// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "base"
                ],
                "summary": "Log in user, issue JWT into Cookie",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/auth.UserGet"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "In the realm of code where data flows,\nA whisper travels where the network goes.\nWith a simple command, a heartbeat is sent,\nTo check the connection, to see where it went.\n\n\"Ping,\" it calls, like a friendly chime,\nA request for response, a dance through time.\nFrom client to server, a message takes flight,\nIn packets it journeys, through day and through night.\n\n\"Are you there?\" it asks, with a digital sigh,\nA promise of data, a link to the sky.\nThe echo returns, a confirmation so sweet,\nA symphony of bytes, where two systems meet.\n\nIn the world of APIs, where services blend,\nPing is the guardian, the reliable friend.\nIt measures the distance, the latency's grace,\nA pulse of the network, a test of the space.\n\nSo here’s to the ping, in its silent decree,\nA bridge in the ether, connecting you and me.\nWith each little packet, a story unfolds,\nIn the language of tech, where the future beholds.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "base"
                ],
                "summary": "Ping API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "base"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/auth.UserGet"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.UserCreate": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "middleName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.UserGet": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "middleName": {
                    "type": "string"
                }
            }
        },
        "auth.UserLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "My magical API",
	Description:      "Learn & Grow API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
