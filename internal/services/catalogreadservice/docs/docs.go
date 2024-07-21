// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mehdi Hadeli",
            "url": "https://github.com/mehdihadeli"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/products": {
            "get": {
                "description": "Get all products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all product",
                "parameters": [
                    {
                        "type": "string",
                        "name": "orderBy",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_features_getting_products_v1_dtos.GetProductsResponseDto"
                        }
                    }
                }
            }
        },
        "/api/v1/products/search": {
            "get": {
                "description": "Search products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Search products",
                "parameters": [
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_features_searching_products_v1_dtos.SearchProductsResponseDto"
                        }
                    }
                }
            }
        },
        "/api/v1/products/{id}": {
            "get": {
                "description": "Get product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_features_get_product_by_id_v1_dtos.GetProductByIdResponseDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_dto.ProductDto": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "productId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_features_get_product_by_id_v1_dtos.GetProductByIdResponseDto": {
            "type": "object",
            "properties": {
                "product": {
                    "$ref": "#/definitions/github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_dto.ProductDto"
                }
            }
        },
        "github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_features_getting_products_v1_dtos.GetProductsResponseDto": {
            "type": "object",
            "properties": {
                "products": {
                    "$ref": "#/definitions/utils.ListResult-github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_dto_ProductDto"
                }
            }
        },
        "github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_features_searching_products_v1_dtos.SearchProductsResponseDto": {
            "type": "object",
            "properties": {
                "products": {
                    "$ref": "#/definitions/utils.ListResult-github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_dto_ProductDto"
                }
            }
        },
        "utils.FilterModel": {
            "type": "object",
            "properties": {
                "comparison": {
                    "type": "string"
                },
                "field": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "utils.ListQuery": {
            "type": "object",
            "properties": {
                "filters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/utils.FilterModel"
                    }
                },
                "orderBy": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "utils.ListResult-github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_dto_ProductDto": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_mehdihadeli_go-food-delivery-microservices_internal_services_catalogreadservice_internal_products_dto.ProductDto"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "totalItems": {
                    "type": "integer"
                },
                "totalPage": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Catalogs Read-Service Api",
	Description:      "Catalogs Read-Service Api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
