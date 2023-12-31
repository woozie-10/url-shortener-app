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
        "/": {
            "get": {
                "description": "Render the index page with HTML content",
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "home"
                ],
                "summary": "Render the index page",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Shorten a given URL using the API",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "url"
                ],
                "summary": "Shorten a given URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "form data of the URL to be shortened",
                        "name": "url",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Content-Type": {
                                "type": "string",
                                "description": "text/html; charset=utf-8"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Content-Type": {
                                "type": "string",
                                "description": "text/html; charset=utf-8"
                            }
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:2311",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "URL Shortener App",
	Description:      "This web application serves as a URL Shortener, built with Go and Gin. Users can easily shorten URLs through a user-friendly interface.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
