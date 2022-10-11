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
        "contact": {
            "name": "Gaël G.",
            "url": "https://blog.gothuey.dev/"
        },
        "license": {
            "name": "MIT license",
            "url": "https://github.com/gaelgoth/swiss-deals-api/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health check 🩺"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.setupRoutes.HealthCheck"
                        }
                    }
                }
            }
        },
        "/deals/digitec": {
            "get": {
                "description": "receive the promotion of the day",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Galaxus 🐢"
                ],
                "summary": "Get current deal from Digitec",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GalaxusProduct"
                        }
                    }
                }
            }
        },
        "/deals/galaxus": {
            "get": {
                "description": "receive the promotion of the day",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Galaxus 🐢"
                ],
                "summary": "Get current deal from Galaxus",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GalaxusProduct"
                        }
                    }
                }
            }
        },
        "/deals/qbeer": {
            "get": {
                "description": "receive the promotion of the day",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Qoqa 🦦"
                ],
                "summary": "Get current deal from Qbeer",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.QoqaProduct"
                        }
                    }
                }
            }
        },
        "/deals/qlock": {
            "get": {
                "description": "receive the promotion of the day",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Qoqa 🦦"
                ],
                "summary": "Get current deal from Qlock",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.QoqaProduct"
                        }
                    }
                }
            }
        },
        "/deals/qooking": {
            "get": {
                "description": "receive the promotion of the day",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Qoqa 🦦"
                ],
                "summary": "Get current deal from Qooking",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.QoqaProduct"
                        }
                    }
                }
            }
        },
        "/deals/qoqa": {
            "get": {
                "description": "receive the promotion of the day",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Qoqa 🦦"
                ],
                "summary": "Get current deal from QoQa",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.QoqaProduct"
                        }
                    }
                }
            }
        },
        "/deals/qsport": {
            "get": {
                "description": "receive the promotion of the day",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Qoqa 🦦"
                ],
                "summary": "Get current deal from Qsport",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.QoqaProduct"
                        }
                    }
                }
            }
        },
        "/deals/qwine": {
            "get": {
                "description": "receive the promotion of the day",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Qoqa 🦦"
                ],
                "summary": "Get current deal from Qwine",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.QoqaProduct"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.GalaxusProduct": {
            "type": "object",
            "properties": {
                "image_url": {
                    "type": "string"
                },
                "offer_price": {
                    "type": "integer"
                },
                "regular_price": {
                    "type": "integer"
                },
                "remaining_stock": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "store": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "controllers.QoqaProduct": {
            "type": "object",
            "properties": {
                "image_url": {
                    "type": "string"
                },
                "offer_price": {
                    "type": "string"
                },
                "regular_price": {
                    "type": "string"
                },
                "remaining_stock": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "store": {
                    "type": "string"
                },
                "subtitle": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "main.setupRoutes.HealthCheck": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "deals-api.gothuey.dev",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swiss Deals API",
	Description:      "Aggregate deals of the day from Digitec, Galaxus, QoQa. Front-end available on http://deals.gothuey.dev/",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
