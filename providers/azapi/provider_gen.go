// Auto-Generated Code; DO NOT EDIT.
package azapi

import (
	"encoding/json"
	"fmt"
	"github.com/magodo/tfadd/schema"
	"os"
)

var ProviderSchemaInfo schema.ProviderSchema

func init() {
    b := []byte(`{
  "Version": "",
  "resource_schemas": {
    "azapi_data_plane_resource": {
      "block": {
        "attributes": [
          {
            "name": "body",
            "type": "dynamic",
            "optional": true,
            "computed": true,
            "default": "{}"
          },
          {
            "name": "id",
            "type": "string",
            "computed": true
          },
          {
            "name": "ignore_casing",
            "type": "bool",
            "optional": true,
            "computed": true,
            "default": false
          },
          {
            "name": "ignore_missing_property",
            "type": "bool",
            "optional": true,
            "computed": true,
            "default": true
          },
          {
            "name": "locks",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "name",
            "type": "string",
            "required": true
          },
          {
            "name": "output",
            "type": "dynamic",
            "computed": true
          },
          {
            "name": "parent_id",
            "type": "string",
            "required": true
          },
          {
            "name": "response_export_values",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "type",
            "type": "string",
            "required": true
          }
        ],
        "block_types": [
          {
            "type_name": "timeouts",
            "block": {
              "attributes": [
                {
                  "name": "create",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "delete",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "read",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "update",
                  "type": "string",
                  "optional": true
                }
              ]
            },
            "nesting_mode": 1
          }
        ]
      }
    },
    "azapi_resource": {
      "block": {
        "attributes": [
          {
            "name": "body",
            "type": "dynamic",
            "optional": true,
            "computed": true,
            "default": "{}"
          },
          {
            "name": "id",
            "type": "string",
            "computed": true
          },
          {
            "name": "ignore_body_changes",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "ignore_casing",
            "type": "bool",
            "optional": true,
            "computed": true,
            "default": false
          },
          {
            "name": "ignore_missing_property",
            "type": "bool",
            "optional": true,
            "computed": true,
            "default": true
          },
          {
            "name": "location",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "locks",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "name",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "output",
            "type": "dynamic",
            "computed": true
          },
          {
            "name": "parent_id",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "removing_special_chars",
            "type": "bool",
            "optional": true,
            "computed": true,
            "default": false
          },
          {
            "name": "response_export_values",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "schema_validation_enabled",
            "type": "bool",
            "optional": true,
            "computed": true,
            "default": true
          },
          {
            "name": "tags",
            "type": [
              "map",
              "string"
            ],
            "optional": true,
            "computed": true
          },
          {
            "name": "type",
            "type": "string",
            "required": true
          }
        ],
        "block_types": [
          {
            "type_name": "identity",
            "block": {
              "attributes": [
                {
                  "name": "identity_ids",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true
                },
                {
                  "name": "principal_id",
                  "type": "string",
                  "computed": true
                },
                {
                  "name": "tenant_id",
                  "type": "string",
                  "computed": true
                },
                {
                  "name": "type",
                  "type": "string",
                  "required": true
                }
              ]
            },
            "nesting_mode": 2
          },
          {
            "type_name": "timeouts",
            "block": {
              "attributes": [
                {
                  "name": "create",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "delete",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "read",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "update",
                  "type": "string",
                  "optional": true
                }
              ]
            },
            "nesting_mode": 1
          }
        ]
      }
    },
    "azapi_resource_action": {
      "block": {
        "attributes": [
          {
            "name": "action",
            "type": "string",
            "optional": true
          },
          {
            "name": "body",
            "type": "dynamic",
            "optional": true
          },
          {
            "name": "id",
            "type": "string",
            "computed": true
          },
          {
            "name": "locks",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "method",
            "type": "string",
            "optional": true,
            "computed": true,
            "default": "POST"
          },
          {
            "name": "output",
            "type": "dynamic",
            "computed": true
          },
          {
            "name": "resource_id",
            "type": "string",
            "required": true
          },
          {
            "name": "response_export_values",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "type",
            "type": "string",
            "required": true
          },
          {
            "name": "when",
            "type": "string",
            "optional": true,
            "computed": true,
            "default": "apply"
          }
        ],
        "block_types": [
          {
            "type_name": "timeouts",
            "block": {
              "attributes": [
                {
                  "name": "create",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "delete",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "read",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "update",
                  "type": "string",
                  "optional": true
                }
              ]
            },
            "nesting_mode": 1
          }
        ]
      }
    },
    "azapi_update_resource": {
      "block": {
        "attributes": [
          {
            "name": "body",
            "type": "dynamic",
            "optional": true,
            "computed": true,
            "default": "{}"
          },
          {
            "name": "id",
            "type": "string",
            "computed": true
          },
          {
            "name": "ignore_body_changes",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "ignore_casing",
            "type": "bool",
            "optional": true,
            "computed": true,
            "default": false
          },
          {
            "name": "ignore_missing_property",
            "type": "bool",
            "optional": true,
            "computed": true,
            "default": true
          },
          {
            "name": "locks",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "name",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "output",
            "type": "dynamic",
            "computed": true
          },
          {
            "name": "parent_id",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "resource_id",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "response_export_values",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "type",
            "type": "string",
            "required": true
          }
        ],
        "block_types": [
          {
            "type_name": "timeouts",
            "block": {
              "attributes": [
                {
                  "name": "create",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "delete",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "read",
                  "type": "string",
                  "optional": true
                },
                {
                  "name": "update",
                  "type": "string",
                  "optional": true
                }
              ]
            },
            "nesting_mode": 1
          }
        ]
      }
    }
  }
}`)
	if err := json.Unmarshal(b, &ProviderSchemaInfo); err != nil {
        fmt.Fprintf(os.Stderr, "unmarshalling the provider schema (azapi): %s", err)
		os.Exit(1)
	}
    ProviderSchemaInfo.Version = "1.15.0"
}
