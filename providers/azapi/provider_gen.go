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
        "attributes": {
          "body": {
            "type": "string",
            "optional": true,
            "default": "{}"
          },
          "ignore_casing": {
            "type": "bool",
            "optional": true,
            "default": false
          },
          "ignore_missing_property": {
            "type": "bool",
            "optional": true,
            "default": true
          },
          "locks": {
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          "name": {
            "type": "string",
            "required": true,
            "force_new": true
          },
          "output": {
            "type": "string",
            "computed": true
          },
          "parent_id": {
            "type": "string",
            "required": true,
            "force_new": true
          },
          "response_export_values": {
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          "type": {
            "type": "string",
            "required": true
          }
        }
      }
    },
    "azapi_resource": {
      "block": {
        "attributes": {
          "body": {
            "type": "string",
            "optional": true,
            "default": "{}"
          },
          "ignore_body_changes": {
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          "ignore_casing": {
            "type": "bool",
            "optional": true,
            "default": false
          },
          "ignore_missing_property": {
            "type": "bool",
            "optional": true,
            "default": true
          },
          "location": {
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          "locks": {
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          "name": {
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          "output": {
            "type": "string",
            "computed": true
          },
          "parent_id": {
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          "removing_special_chars": {
            "type": "bool",
            "optional": true,
            "default": false
          },
          "response_export_values": {
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          "schema_validation_enabled": {
            "type": "bool",
            "optional": true,
            "default": true
          },
          "tags": {
            "type": [
              "map",
              "string"
            ],
            "optional": true,
            "computed": true
          },
          "type": {
            "type": "string",
            "required": true
          }
        },
        "block_types": {
          "identity": {
            "nesting_mode": 3,
            "block": {
              "attributes": {
                "identity_ids": {
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true
                },
                "principal_id": {
                  "type": "string",
                  "computed": true
                },
                "tenant_id": {
                  "type": "string",
                  "computed": true
                },
                "type": {
                  "type": "string",
                  "required": true
                }
              }
            },
            "optional": true,
            "computed": true,
            "max_items": 1
          }
        }
      }
    },
    "azapi_resource_action": {
      "block": {
        "attributes": {
          "action": {
            "type": "string",
            "optional": true,
            "force_new": true
          },
          "body": {
            "type": "string",
            "optional": true
          },
          "locks": {
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          "method": {
            "type": "string",
            "optional": true,
            "default": "POST"
          },
          "output": {
            "type": "string",
            "computed": true
          },
          "resource_id": {
            "type": "string",
            "required": true,
            "force_new": true
          },
          "response_export_values": {
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          "type": {
            "type": "string",
            "required": true,
            "force_new": true
          }
        }
      }
    },
    "azapi_update_resource": {
      "block": {
        "attributes": {
          "body": {
            "type": "string",
            "optional": true,
            "default": "{}"
          },
          "ignore_body_changes": {
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          "ignore_casing": {
            "type": "bool",
            "optional": true,
            "default": false
          },
          "ignore_missing_property": {
            "type": "bool",
            "optional": true,
            "default": true
          },
          "locks": {
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          "name": {
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "exactly_one_of": [
              "name",
              "resource_id"
            ],
            "required_with": [
              "parent_id"
            ]
          },
          "output": {
            "type": "string",
            "computed": true
          },
          "parent_id": {
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "required_with": [
              "name"
            ]
          },
          "resource_id": {
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "exactly_one_of": [
              "name",
              "resource_id"
            ]
          },
          "response_export_values": {
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          "type": {
            "type": "string",
            "required": true
          }
        }
      }
    }
  }
}`)
	if err := json.Unmarshal(b, &ProviderSchemaInfo); err != nil {
		fmt.Fprintf(os.Stderr, "unmarshalling the provider schema: %s", err)
		os.Exit(1)
	}
    ProviderSchemaInfo.Version = "1.9.0"
}
