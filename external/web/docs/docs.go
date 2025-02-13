// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Cria um novo usuário com os dados fornecidos e retorna os detalhes do usuário, incluindo um cookie de autenticação.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Cria um novo usuário",
                "parameters": [
                    {
                        "description": "Dados do novo usuário",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserInputDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Usuário criado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/user.UserOutputDto"
                        }
                    },
                    "400": {
                        "description": "Erro de validação dos dados",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/patients/save": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Recebe os dados de um novo paciente e os salva no sistema.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patients"
                ],
                "summary": "Salva um novo paciente",
                "parameters": [
                    {
                        "description": "Dados do paciente",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/patient.NewPatientInputDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Paciente salvo com sucesso",
                        "schema": {
                            "$ref": "#/definitions/patient.PatientOutputDto"
                        }
                    },
                    "400": {
                        "description": "Erro de validação dos dados",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/signin": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Autentica um usuário com email e senha fornecidos e retorna um cookie para manter a sessão do usuário.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Autentica um usuário e retorna um cookie de autenticação",
                "parameters": [
                    {
                        "description": "user email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "user password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Autenticação bem-sucedida",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Erro de validação dos dados",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/signout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Remove o cookie de autenticação do usuário, invalidando a sessão.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Desconecta o usuário removendo o cookie de autenticação",
                "responses": {
                    "200": {
                        "description": "Desconexão bem-sucedida",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "patient.NewPatientInputDto": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "patient.PatientOutputDto": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "user.CreateUserInputDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "origin": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "user.UserOutputDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://localhost:3000",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample Neofarma server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
