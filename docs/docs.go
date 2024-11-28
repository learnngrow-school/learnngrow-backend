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
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
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
        "/auth/me": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get current user data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.UserGet"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
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
        },
        "/courses/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Courses"
                ],
                "summary": "Gets All courses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/courses.CourseWithData"
                            }
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
                    "Courses"
                ],
                "summary": "Create a course",
                "parameters": [
                    {
                        "description": "Course",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/courses.CourseCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/courses.Course"
                        }
                    }
                }
            }
        },
        "/courses/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Courses"
                ],
                "summary": "Get one course by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/courses.CourseWithData"
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
        "/reviews": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "Get all school reviews",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/reviews.ReviewGet"
                            }
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
                    "Reviews"
                ],
                "summary": "Create review",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reviews.ReviewCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/reviews.ReviewGet"
                        }
                    }
                }
            }
        },
        "/teachers/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teachers"
                ],
                "summary": "Get all teachers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/teachers.TeacherGet"
                            }
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
                    "Teachers"
                ],
                "summary": "Create teacher",
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
                "responses": {}
            }
        },
        "/teachers/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teachers"
                ],
                "summary": "Get teacher by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Teacher ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/teachers.TeacherGet"
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
                "firstName",
                "lastName",
                "password",
                "phone"
            ],
            "properties": {
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
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "auth.UserGet": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "middleName": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "auth.UserLogin": {
            "type": "object",
            "required": [
                "password",
                "phone"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "courses.Category": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "courses.Course": {
            "type": "object",
            "required": [
                "grade",
                "price",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "grade": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "courses.CourseCreate": {
            "type": "object",
            "required": [
                "categoryId",
                "grade",
                "price",
                "subjectId",
                "title"
            ],
            "properties": {
                "categoryId": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "grade": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "subjectId": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "courses.CourseWithData": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/courses.Category"
                },
                "course": {
                    "$ref": "#/definitions/courses.Course"
                },
                "subject": {
                    "$ref": "#/definitions/courses.Subject"
                }
            }
        },
        "courses.Subject": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "reviews.ReviewCreate": {
            "type": "object",
            "properties": {
                "authorName": {
                    "type": "string"
                },
                "details": {
                    "type": "string"
                }
            }
        },
        "reviews.ReviewGet": {
            "type": "object",
            "properties": {
                "authorName": {
                    "type": "string"
                },
                "details": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "teachers.Teacher": {
            "type": "object",
            "properties": {
                "biography": {
                    "type": "string"
                },
                "subjectIds": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "teachers.TeacherGet": {
            "type": "object",
            "properties": {
                "teacherData": {
                    "$ref": "#/definitions/teachers.Teacher"
                },
                "userData": {
                    "$ref": "#/definitions/auth.UserGet"
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
