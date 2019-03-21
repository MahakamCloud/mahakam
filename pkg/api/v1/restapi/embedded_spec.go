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
    "/apps": {
      "get": {
        "tags": [
          "apps"
        ],
        "operationId": "getApps",
        "responses": {
          "200": {
            "description": "list created applications",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/app"
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
          "apps"
        ],
        "operationId": "createApp",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/app"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/app"
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
    "/apps/values": {
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "apps"
        ],
        "operationId": "uploadAppValues",
        "parameters": [
          {
            "type": "file",
            "description": "App values.yaml to upload",
            "name": "values",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "App name",
            "name": "appName",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Owner of the app",
            "name": "owner",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Cluster name to deploy app",
            "name": "clusterName",
            "in": "formData"
          }
        ],
        "responses": {
          "201": {
            "description": "App values added"
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
    "/clusters": {
      "get": {
        "tags": [
          "clusters"
        ],
        "operationId": "getClusters",
        "parameters": [
          {
            "type": "string",
            "description": "Owner of clusters",
            "name": "owner",
            "in": "query"
          }
        ],
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
    },
    "/clusters/validate": {
      "post": {
        "tags": [
          "clusters"
        ],
        "operationId": "validateCluster",
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
            "description": "Validated",
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
    "/networks": {
      "get": {
        "tags": [
          "networks"
        ],
        "operationId": "getNetworks",
        "responses": {
          "200": {
            "description": "list created networks",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/network"
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
          "networks"
        ],
        "operationId": "createNetwork",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/network"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/network"
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
    "/networks/pools/ipPools": {
      "post": {
        "tags": [
          "networks"
        ],
        "operationId": "createIpPool",
        "parameters": [
          {
            "name": "ipPool",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ipPool"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created new ip pool",
            "schema": {
              "$ref": "#/definitions/ipPool"
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
    },
    "/networks/pools/ipPools/{poolId}": {
      "post": {
        "tags": [
          "networks"
        ],
        "operationId": "allocateOrReleaseFromIpPool",
        "parameters": [
          {
            "type": "string",
            "name": "poolId",
            "in": "path"
          },
          {
            "type": "string",
            "name": "allocatedIP",
            "in": "body"
          },
          {
            "type": "string",
            "name": "releasedIP",
            "in": "query"
          },
          {
            "enum": [
              "ALLOCATE",
              "RELEASE"
            ],
            "type": "string",
            "name": "action",
            "in": "query"
          }
        ],
        "responses": {
          "201": {
            "description": "Allocated new IP from IP pool",
            "schema": {
              "$ref": "#/definitions/allocatedIpPool"
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
    },
    "/nodes": {
      "get": {
        "tags": [
          "nodes"
        ],
        "operationId": "getNodes",
        "responses": {
          "200": {
            "description": "list created nodes",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/node"
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
          "nodes"
        ],
        "operationId": "createNode",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/node"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/node"
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
    }
  },
  "definitions": {
    "allocatedIpPool": {
      "type": "object",
      "properties": {
        "allocatedIp": {
          "type": "string"
        },
        "ipPoolId": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    },
    "app": {
      "type": "object",
      "required": [
        "name"
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
        "serviceFQDN": {
          "type": "string"
        },
        "status": {
          "type": "string"
        }
      }
    },
    "baseResource": {
      "type": "object",
      "properties": {
        "createdTime": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "kind": {
          "type": "string"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "key": {
                "type": "string"
              },
              "value": {
                "type": "string"
              }
            }
          }
        },
        "modifiedTime": {
          "type": "string",
          "format": "date-time"
        },
        "name": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "revision": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "cluster": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "componentFailures": {
          "type": "array",
          "items": {
            "type": "string"
          }
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
        "nodeFailures": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "nodeSize": {
          "type": "string",
          "default": "small",
          "enum": [
            "extrasmall",
            "small",
            "medium",
            "large",
            "extralarge"
          ]
        },
        "numNodes": {
          "type": "integer",
          "maximum": 10,
          "minimum": 1
        },
        "owner": {
          "type": "string",
          "minLength": 1
        },
        "podFailures": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ipPool": {
      "type": "object",
      "properties": {
        "cidr": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "readOnly": true
        },
        "ipPoolRangeEnd": {
          "type": "string"
        },
        "ipPoolRangeStart": {
          "type": "string"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "key": {
                "type": "string"
              },
              "value": {
                "type": "string"
              }
            }
          }
        },
        "reservedIPPools": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "metadata": {
      "type": "object",
      "properties": {
        "sshKeys": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "userdata": {
          "type": "string"
        }
      },
      "additionalProperties": {
        "type": "string",
        "format": "string"
      }
    },
    "network": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "dhcp": {
          "type": "string"
        },
        "gateway": {
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
        "nameserver": {
          "type": "string"
        },
        "networkCIDR": {
          "type": "string"
        },
        "networkName": {
          "type": "string"
        }
      }
    },
    "networkConfig": {
      "type": "object",
      "properties": {
        "fqdn": {
          "type": "string"
        },
        "gatewayIP": {
          "type": "string",
          "format": "ipv4"
        },
        "ip": {
          "type": "string",
          "format": "ipv4"
        },
        "ipMask": {
          "type": "string"
        },
        "mac": {
          "type": "string",
          "format": "mac"
        },
        "nameserverIP": {
          "type": "string",
          "format": "ipv4"
        }
      }
    },
    "node": {
      "type": "object",
      "required": [
        "name"
      ],
      "allOf": [
        {
          "$ref": "#/definitions/baseResource"
        }
      ],
      "properties": {
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "metadata": {
          "$ref": "#/definitions/metadata"
        },
        "modifiedAt": {
          "type": "string",
          "format": "date-time"
        },
        "name": {
          "type": "string"
        },
        "networkConfigs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/networkConfig"
          }
        },
        "nodespec": {
          "$ref": "#/definitions/nodeSpec"
        },
        "owner": {
          "type": "string"
        },
        "revision": {
          "type": "integer",
          "format": "uint64"
        },
        "status": {
          "$ref": "#/definitions/nodeStatus"
        }
      }
    },
    "nodeSpec": {
      "type": "object",
      "properties": {
        "cpu": {
          "type": "integer"
        },
        "memory": {
          "type": "string"
        }
      }
    },
    "nodeStatus": {
      "type": "object",
      "properties": {
        "host": {
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
    "/apps": {
      "get": {
        "tags": [
          "apps"
        ],
        "operationId": "getApps",
        "responses": {
          "200": {
            "description": "list created applications",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/app"
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
          "apps"
        ],
        "operationId": "createApp",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/app"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/app"
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
    "/apps/values": {
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "apps"
        ],
        "operationId": "uploadAppValues",
        "parameters": [
          {
            "type": "file",
            "description": "App values.yaml to upload",
            "name": "values",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "App name",
            "name": "appName",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Owner of the app",
            "name": "owner",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Cluster name to deploy app",
            "name": "clusterName",
            "in": "formData"
          }
        ],
        "responses": {
          "201": {
            "description": "App values added"
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
    "/clusters": {
      "get": {
        "tags": [
          "clusters"
        ],
        "operationId": "getClusters",
        "parameters": [
          {
            "type": "string",
            "description": "Owner of clusters",
            "name": "owner",
            "in": "query"
          }
        ],
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
    },
    "/clusters/validate": {
      "post": {
        "tags": [
          "clusters"
        ],
        "operationId": "validateCluster",
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
            "description": "Validated",
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
    "/networks": {
      "get": {
        "tags": [
          "networks"
        ],
        "operationId": "getNetworks",
        "responses": {
          "200": {
            "description": "list created networks",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/network"
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
          "networks"
        ],
        "operationId": "createNetwork",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/network"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/network"
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
    "/networks/pools/ipPools": {
      "post": {
        "tags": [
          "networks"
        ],
        "operationId": "createIpPool",
        "parameters": [
          {
            "name": "ipPool",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ipPool"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created new ip pool",
            "schema": {
              "$ref": "#/definitions/ipPool"
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
    },
    "/networks/pools/ipPools/{poolId}": {
      "post": {
        "tags": [
          "networks"
        ],
        "operationId": "allocateOrReleaseFromIpPool",
        "parameters": [
          {
            "type": "string",
            "name": "poolId",
            "in": "path"
          },
          {
            "type": "string",
            "name": "allocatedIP",
            "in": "body"
          },
          {
            "type": "string",
            "name": "releasedIP",
            "in": "query"
          },
          {
            "enum": [
              "ALLOCATE",
              "RELEASE"
            ],
            "type": "string",
            "name": "action",
            "in": "query"
          }
        ],
        "responses": {
          "201": {
            "description": "Allocated new IP from IP pool",
            "schema": {
              "$ref": "#/definitions/allocatedIpPool"
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
    },
    "/nodes": {
      "get": {
        "tags": [
          "nodes"
        ],
        "operationId": "getNodes",
        "responses": {
          "200": {
            "description": "list created nodes",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/node"
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
          "nodes"
        ],
        "operationId": "createNode",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/node"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/node"
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
    }
  },
  "definitions": {
    "allocatedIpPool": {
      "type": "object",
      "properties": {
        "allocatedIp": {
          "type": "string"
        },
        "ipPoolId": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    },
    "app": {
      "type": "object",
      "required": [
        "name"
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
        "serviceFQDN": {
          "type": "string"
        },
        "status": {
          "type": "string"
        }
      }
    },
    "baseResource": {
      "type": "object",
      "properties": {
        "createdTime": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "kind": {
          "type": "string"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "key": {
                "type": "string"
              },
              "value": {
                "type": "string"
              }
            }
          }
        },
        "modifiedTime": {
          "type": "string",
          "format": "date-time"
        },
        "name": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "revision": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "cluster": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "componentFailures": {
          "type": "array",
          "items": {
            "type": "string"
          }
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
        "nodeFailures": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "nodeSize": {
          "type": "string",
          "default": "small",
          "enum": [
            "extrasmall",
            "small",
            "medium",
            "large",
            "extralarge"
          ]
        },
        "numNodes": {
          "type": "integer",
          "maximum": 10,
          "minimum": 1
        },
        "owner": {
          "type": "string",
          "minLength": 1
        },
        "podFailures": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ipPool": {
      "type": "object",
      "properties": {
        "cidr": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "readOnly": true
        },
        "ipPoolRangeEnd": {
          "type": "string"
        },
        "ipPoolRangeStart": {
          "type": "string"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "key": {
                "type": "string"
              },
              "value": {
                "type": "string"
              }
            }
          }
        },
        "reservedIPPools": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "metadata": {
      "type": "object",
      "properties": {
        "sshKeys": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "userdata": {
          "type": "string"
        }
      },
      "additionalProperties": {
        "type": "string",
        "format": "string"
      }
    },
    "network": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "dhcp": {
          "type": "string"
        },
        "gateway": {
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
        "nameserver": {
          "type": "string"
        },
        "networkCIDR": {
          "type": "string"
        },
        "networkName": {
          "type": "string"
        }
      }
    },
    "networkConfig": {
      "type": "object",
      "properties": {
        "fqdn": {
          "type": "string"
        },
        "gatewayIP": {
          "type": "string",
          "format": "ipv4"
        },
        "ip": {
          "type": "string",
          "format": "ipv4"
        },
        "ipMask": {
          "type": "string"
        },
        "mac": {
          "type": "string",
          "format": "mac"
        },
        "nameserverIP": {
          "type": "string",
          "format": "ipv4"
        }
      }
    },
    "node": {
      "type": "object",
      "required": [
        "name"
      ],
      "allOf": [
        {
          "$ref": "#/definitions/baseResource"
        }
      ],
      "properties": {
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "metadata": {
          "$ref": "#/definitions/metadata"
        },
        "modifiedAt": {
          "type": "string",
          "format": "date-time"
        },
        "name": {
          "type": "string"
        },
        "networkConfigs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/networkConfig"
          }
        },
        "nodespec": {
          "$ref": "#/definitions/nodeSpec"
        },
        "owner": {
          "type": "string"
        },
        "revision": {
          "type": "integer",
          "format": "uint64"
        },
        "status": {
          "$ref": "#/definitions/nodeStatus"
        }
      }
    },
    "nodeSpec": {
      "type": "object",
      "properties": {
        "cpu": {
          "type": "integer"
        },
        "memory": {
          "type": "string"
        }
      }
    },
    "nodeStatus": {
      "type": "object",
      "properties": {
        "host": {
          "type": "string"
        }
      }
    }
  }
}`))
}
