// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API definitions for Mahakam Cloud Platform",
    "title": "Mahakam API",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/clusters": {
      "get": {
        "tags": [
          "clusters"
        ],
        "operationId": "getClusters",
        "responses": {
          "200": {
            "description": "list available clusters",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/cluster"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "clusters"
        ],
        "operationId": "createCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/clusters/describe": {
      "get": {
        "tags": [
          "clusters"
        ],
        "operationId": "describeClusters",
        "parameters": [
          {
            "type": "string",
            "description": "Cluster name",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "describe single cluster",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "app": {
      "type": "object",
      "required": [
        "name",
        "owner"
      ],
      "properties": {
        "chartURL": {
          "type": "string"
        },
        "chartValues": {
          "type": "string"
        },
        "clusterName": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "status": {
          "type": "string"
        }
      }
    },
    "cluster": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "clusterPlan": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "numNodes": {
          "type": "integer",
          "maximum": 10,
          "minimum": 3
        },
        "owner": {
          "type": "string",
          "minLength": 1
        },
        "status": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API definitions for Mahakam Cloud Platform",
    "title": "Mahakam API",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/clusters": {
      "get": {
        "tags": [
          "clusters"
        ],
        "operationId": "getClusters",
        "responses": {
          "200": {
            "description": "list available clusters",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/cluster"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "clusters"
        ],
        "operationId": "createCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/clusters/describe": {
      "get": {
        "tags": [
          "clusters"
        ],
        "operationId": "describeClusters",
        "parameters": [
          {
            "type": "string",
            "description": "Cluster name",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "describe single cluster",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "app": {
      "type": "object",
      "required": [
        "name",
        "owner"
      ],
      "properties": {
        "chartURL": {
          "type": "string"
        },
        "chartValues": {
          "type": "string"
        },
        "clusterName": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "status": {
          "type": "string"
        }
      }
    },
    "cluster": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "clusterPlan": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "numNodes": {
          "type": "integer",
          "maximum": 10,
          "minimum": 3
        },
        "owner": {
          "type": "string",
          "minLength": 1
        },
        "status": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}
