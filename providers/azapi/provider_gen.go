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
            "default": {}
          },
          {
            "name": "create_headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "create_query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
          },
          {
            "name": "delete_headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "delete_query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
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
            "name": "read_headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "read_query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
          },
          {
            "name": "replace_triggers_external_values",
            "type": "dynamic",
            "optional": true
          },
          {
            "name": "replace_triggers_refs",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "response_export_values",
            "type": "dynamic",
            "optional": true
          },
          {
            "name": "retry",
            "nested_type": {
              "Attributes": [
                {
                  "name": "error_message_regex",
                  "type": [
                    "list",
                    "string"
                  ],
                  "required": true
                },
                {
                  "name": "interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": 10
                },
                {
                  "name": "max_interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": 180
                },
                {
                  "name": "multiplier",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": "1.5"
                },
                {
                  "name": "randomization_factor",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": "0.5"
                }
              ],
              "Nesting": 1
            },
            "optional": true
          },
          {
            "name": "type",
            "type": "string",
            "required": true
          },
          {
            "name": "update_headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "update_query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
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
            "default": {}
          },
          {
            "name": "create_headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "create_query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
          },
          {
            "name": "delete_headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "delete_query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
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
            "name": "read_headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "read_query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
          },
          {
            "name": "replace_triggers_external_values",
            "type": "dynamic",
            "optional": true
          },
          {
            "name": "replace_triggers_refs",
            "type": [
              "list",
              "string"
            ],
            "optional": true
          },
          {
            "name": "response_export_values",
            "type": "dynamic",
            "optional": true
          },
          {
            "name": "retry",
            "nested_type": {
              "Attributes": [
                {
                  "name": "error_message_regex",
                  "type": [
                    "list",
                    "string"
                  ],
                  "required": true
                },
                {
                  "name": "interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": 10
                },
                {
                  "name": "max_interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": 180
                },
                {
                  "name": "multiplier",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": "1.5"
                },
                {
                  "name": "randomization_factor",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": "0.5"
                }
              ],
              "Nesting": 1
            },
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
          },
          {
            "name": "update_headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "update_query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
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
            "name": "headers",
            "type": [
              "map",
              "string"
            ],
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
            "name": "query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
          },
          {
            "name": "resource_id",
            "type": "string",
            "required": true
          },
          {
            "name": "response_export_values",
            "type": "dynamic",
            "optional": true
          },
          {
            "name": "retry",
            "nested_type": {
              "Attributes": [
                {
                  "name": "error_message_regex",
                  "type": [
                    "list",
                    "string"
                  ],
                  "required": true
                },
                {
                  "name": "interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": 10
                },
                {
                  "name": "max_interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": 180
                },
                {
                  "name": "multiplier",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": "1.5"
                },
                {
                  "name": "randomization_factor",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": "0.5"
                }
              ],
              "Nesting": 1
            },
            "optional": true
          },
          {
            "name": "sensitive_output",
            "type": "dynamic",
            "computed": true,
            "sensitive": true
          },
          {
            "name": "sensitive_response_export_values",
            "type": "dynamic",
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
            "optional": true
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
            "name": "read_headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "read_query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
          },
          {
            "name": "resource_id",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "response_export_values",
            "type": "dynamic",
            "optional": true
          },
          {
            "name": "retry",
            "nested_type": {
              "Attributes": [
                {
                  "name": "error_message_regex",
                  "type": [
                    "list",
                    "string"
                  ],
                  "required": true
                },
                {
                  "name": "interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": 10
                },
                {
                  "name": "max_interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": 180
                },
                {
                  "name": "multiplier",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": "1.5"
                },
                {
                  "name": "randomization_factor",
                  "type": "number",
                  "optional": true,
                  "computed": true,
                  "default": "0.5"
                }
              ],
              "Nesting": 1
            },
            "optional": true
          },
          {
            "name": "type",
            "type": "string",
            "required": true
          },
          {
            "name": "update_headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "update_query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
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
  },
  "datasource_schemas": {
    "azapi_client_config": {
      "block": {
        "attributes": [
          {
            "name": "id",
            "type": "string",
            "computed": true
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true
          },
          {
            "name": "subscription_id",
            "type": "string",
            "computed": true
          },
          {
            "name": "subscription_resource_id",
            "type": "string",
            "computed": true
          },
          {
            "name": "tenant_id",
            "type": "string",
            "computed": true
          }
        ],
        "block_types": [
          {
            "type_name": "timeouts",
            "block": {
              "attributes": [
                {
                  "name": "read",
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
            "name": "headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "id",
            "type": "string",
            "computed": true
          },
          {
            "name": "identity",
            "nested_type": {
              "Attributes": [
                {
                  "name": "identity_ids",
                  "type": [
                    "list",
                    "string"
                  ],
                  "computed": true
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
                  "computed": true
                }
              ],
              "Nesting": 2
            },
            "computed": true
          },
          {
            "name": "location",
            "type": "string",
            "computed": true
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
            "name": "query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
          },
          {
            "name": "resource_id",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "response_export_values",
            "type": "dynamic",
            "optional": true
          },
          {
            "name": "retry",
            "nested_type": {
              "Attributes": [
                {
                  "name": "error_message_regex",
                  "type": [
                    "list",
                    "string"
                  ],
                  "required": true
                },
                {
                  "name": "interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true
                },
                {
                  "name": "max_interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true
                },
                {
                  "name": "multiplier",
                  "type": "number",
                  "optional": true,
                  "computed": true
                },
                {
                  "name": "randomization_factor",
                  "type": "number",
                  "optional": true,
                  "computed": true
                }
              ],
              "Nesting": 1
            },
            "optional": true
          },
          {
            "name": "tags",
            "type": [
              "map",
              "string"
            ],
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
            "type_name": "timeouts",
            "block": {
              "attributes": [
                {
                  "name": "read",
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
            "name": "headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "id",
            "type": "string",
            "computed": true
          },
          {
            "name": "method",
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
            "name": "query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
          },
          {
            "name": "resource_id",
            "type": "string",
            "optional": true
          },
          {
            "name": "response_export_values",
            "type": "dynamic",
            "optional": true
          },
          {
            "name": "retry",
            "nested_type": {
              "Attributes": [
                {
                  "name": "error_message_regex",
                  "type": [
                    "list",
                    "string"
                  ],
                  "required": true
                },
                {
                  "name": "interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true
                },
                {
                  "name": "max_interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true
                },
                {
                  "name": "multiplier",
                  "type": "number",
                  "optional": true,
                  "computed": true
                },
                {
                  "name": "randomization_factor",
                  "type": "number",
                  "optional": true,
                  "computed": true
                }
              ],
              "Nesting": 1
            },
            "optional": true
          },
          {
            "name": "sensitive_output",
            "type": "dynamic",
            "computed": true,
            "sensitive": true
          },
          {
            "name": "sensitive_response_export_values",
            "type": "dynamic",
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
                  "name": "read",
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
    "azapi_resource_id": {
      "block": {
        "attributes": [
          {
            "name": "id",
            "type": "string",
            "computed": true
          },
          {
            "name": "name",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "parent_id",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "parts",
            "type": [
              "map",
              "string"
            ],
            "computed": true
          },
          {
            "name": "provider_namespace",
            "type": "string",
            "computed": true
          },
          {
            "name": "resource_group_name",
            "type": "string",
            "computed": true
          },
          {
            "name": "resource_id",
            "type": "string",
            "optional": true,
            "computed": true
          },
          {
            "name": "subscription_id",
            "type": "string",
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
            "type_name": "timeouts",
            "block": {
              "attributes": [
                {
                  "name": "read",
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
    "azapi_resource_list": {
      "block": {
        "attributes": [
          {
            "name": "headers",
            "type": [
              "map",
              "string"
            ],
            "optional": true
          },
          {
            "name": "id",
            "type": "string",
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
            "required": true
          },
          {
            "name": "query_parameters",
            "type": [
              "map",
              [
                "list",
                "string"
              ]
            ],
            "optional": true
          },
          {
            "name": "response_export_values",
            "type": "dynamic",
            "optional": true
          },
          {
            "name": "retry",
            "nested_type": {
              "Attributes": [
                {
                  "name": "error_message_regex",
                  "type": [
                    "list",
                    "string"
                  ],
                  "required": true
                },
                {
                  "name": "interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true
                },
                {
                  "name": "max_interval_seconds",
                  "type": "number",
                  "optional": true,
                  "computed": true
                },
                {
                  "name": "multiplier",
                  "type": "number",
                  "optional": true,
                  "computed": true
                },
                {
                  "name": "randomization_factor",
                  "type": "number",
                  "optional": true,
                  "computed": true
                }
              ],
              "Nesting": 1
            },
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
                  "name": "read",
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
    ProviderSchemaInfo.Version = "2.3.0"
}
