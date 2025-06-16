// Auto-Generated Code; DO NOT EDIT.
package azuread

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
    "azuread_access_package": {
      "block": {
        "attributes": [
          {
            "name": "catalog_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "description",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "hidden",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          }
        ]
      }
    },
    "azuread_access_package_assignment_policy": {
      "block": {
        "attributes": [
          {
            "name": "access_package_id",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "description",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "duration_in_days",
            "type": "number",
            "optional": true,
            "force_new": false,
            "conflicts_with": [
              "expiration_date"
            ]
          },
          {
            "name": "expiration_date",
            "type": "string",
            "optional": true,
            "force_new": false,
            "conflicts_with": [
              "duration_in_days"
            ]
          },
          {
            "name": "extension_enabled",
            "type": "bool",
            "optional": true,
            "force_new": false
          }
        ],
        "block_types": [
          {
            "type_name": "approval_settings",
            "block": {
              "attributes": [
                {
                  "name": "approval_required",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "approval_required_for_extension",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "requestor_justification_required",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                }
              ],
              "block_types": [
                {
                  "type_name": "approval_stage",
                  "block": {
                    "attributes": [
                      {
                        "name": "alternative_approval_enabled",
                        "type": "bool",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "approval_timeout_in_days",
                        "type": "number",
                        "required": true,
                        "force_new": false
                      },
                      {
                        "name": "approver_justification_required",
                        "type": "bool",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "enable_alternative_approval_in_days",
                        "type": "number",
                        "optional": true,
                        "force_new": false
                      }
                    ],
                    "block_types": [
                      {
                        "type_name": "alternative_approver",
                        "block": {
                          "attributes": [
                            {
                              "name": "backup",
                              "type": "bool",
                              "optional": true,
                              "force_new": false
                            },
                            {
                              "name": "object_id",
                              "type": "string",
                              "optional": true,
                              "force_new": false
                            },
                            {
                              "name": "subject_type",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "required": false,
                        "optional": true,
                        "computed": false,
                        "force_new": false
                      },
                      {
                        "type_name": "primary_approver",
                        "block": {
                          "attributes": [
                            {
                              "name": "backup",
                              "type": "bool",
                              "optional": true,
                              "force_new": false
                            },
                            {
                              "name": "object_id",
                              "type": "string",
                              "optional": true,
                              "force_new": false
                            },
                            {
                              "name": "subject_type",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "required": false,
                        "optional": true,
                        "computed": false,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "assignment_review_settings",
            "block": {
              "attributes": [
                {
                  "name": "access_recommendation_enabled",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "access_review_timeout_behavior",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "approver_justification_required",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "duration_in_days",
                  "type": "number",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "enabled",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "review_frequency",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "review_type",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "starting_on",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                }
              ],
              "block_types": [
                {
                  "type_name": "reviewer",
                  "block": {
                    "attributes": [
                      {
                        "name": "backup",
                        "type": "bool",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "object_id",
                        "type": "string",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "subject_type",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "question",
            "block": {
              "attributes": [
                {
                  "name": "required",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "sequence",
                  "type": "number",
                  "optional": true,
                  "force_new": false
                }
              ],
              "block_types": [
                {
                  "type_name": "choice",
                  "block": {
                    "attributes": [
                      {
                        "name": "actual_value",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      }
                    ],
                    "block_types": [
                      {
                        "type_name": "display_value",
                        "block": {
                          "attributes": [
                            {
                              "name": "default_text",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ],
                          "block_types": [
                            {
                              "type_name": "localized_text",
                              "block": {
                                "attributes": [
                                  {
                                    "name": "content",
                                    "type": "string",
                                    "required": true,
                                    "force_new": false
                                  },
                                  {
                                    "name": "language_code",
                                    "type": "string",
                                    "required": true,
                                    "force_new": false
                                  }
                                ]
                              },
                              "nesting_mode": 2,
                              "required": false,
                              "optional": true,
                              "computed": false,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "min_items": 1,
                        "max_items": 1,
                        "required": true,
                        "optional": false,
                        "computed": false,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                },
                {
                  "type_name": "text",
                  "block": {
                    "attributes": [
                      {
                        "name": "default_text",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      }
                    ],
                    "block_types": [
                      {
                        "type_name": "localized_text",
                        "block": {
                          "attributes": [
                            {
                              "name": "content",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "language_code",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "required": false,
                        "optional": true,
                        "computed": false,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "min_items": 1,
                  "max_items": 1,
                  "required": true,
                  "optional": false,
                  "computed": false,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "requestor_settings",
            "block": {
              "attributes": [
                {
                  "name": "requests_accepted",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "scope_type",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                }
              ],
              "block_types": [
                {
                  "type_name": "requestor",
                  "block": {
                    "attributes": [
                      {
                        "name": "backup",
                        "type": "bool",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "object_id",
                        "type": "string",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "subject_type",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          }
        ]
      }
    },
    "azuread_access_package_catalog": {
      "block": {
        "attributes": [
          {
            "name": "description",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "externally_visible",
            "type": "bool",
            "optional": true,
            "default": true,
            "force_new": false
          },
          {
            "name": "published",
            "type": "bool",
            "optional": true,
            "default": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_access_package_catalog_role_assignment": {
      "block": {
        "attributes": [
          {
            "name": "catalog_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "principal_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "role_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_access_package_resource_catalog_association": {
      "block": {
        "attributes": [
          {
            "name": "catalog_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "resource_origin_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "resource_origin_system",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_access_package_resource_package_association": {
      "block": {
        "attributes": [
          {
            "name": "access_package_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "access_type",
            "type": "string",
            "optional": true,
            "default": "Member",
            "force_new": true
          },
          {
            "name": "catalog_resource_association_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_administrative_unit": {
      "block": {
        "attributes": [
          {
            "name": "description",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "hidden_membership_enabled",
            "type": "bool",
            "optional": true,
            "force_new": false
          },
          {
            "name": "members",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "prevent_duplicate_names",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          }
        ]
      }
    },
    "azuread_administrative_unit_member": {
      "block": {
        "attributes": [
          {
            "name": "administrative_unit_object_id",
            "type": "string",
            "optional": true,
            "force_new": true
          },
          {
            "name": "member_object_id",
            "type": "string",
            "optional": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_administrative_unit_role_member": {
      "block": {
        "attributes": [
          {
            "name": "administrative_unit_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "member_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "role_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_app_role_assignment": {
      "block": {
        "attributes": [
          {
            "name": "app_role_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "principal_display_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "principal_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "principal_type",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "resource_display_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "resource_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_application": {
      "block": {
        "attributes": [
          {
            "name": "app_role_ids",
            "type": [
              "map",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "client_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "description",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "device_only_auth_enabled",
            "type": "bool",
            "optional": true,
            "force_new": false
          },
          {
            "name": "disabled_by_microsoft",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "fallback_public_client_enabled",
            "type": "bool",
            "optional": true,
            "force_new": false
          },
          {
            "name": "group_membership_claims",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false
          },
          {
            "name": "identifier_uris",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false
          },
          {
            "name": "logo_image",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "logo_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "marketing_url",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "notes",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "oauth2_permission_scope_ids",
            "type": [
              "map",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "oauth2_post_response_required",
            "type": "bool",
            "optional": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "owners",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false
          },
          {
            "name": "prevent_duplicate_names",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          },
          {
            "name": "privacy_statement_url",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "publisher_domain",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "service_management_reference",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "sign_in_audience",
            "type": "string",
            "optional": true,
            "default": "AzureADMyOrg",
            "force_new": false
          },
          {
            "name": "support_url",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "tags",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "feature_tags"
            ]
          },
          {
            "name": "template_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "terms_of_service_url",
            "type": "string",
            "optional": true,
            "force_new": false
          }
        ],
        "block_types": [
          {
            "type_name": "api",
            "block": {
              "attributes": [
                {
                  "name": "known_client_applications",
                  "type": [
                    "set",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "mapped_claims_enabled",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "requested_access_token_version",
                  "type": "number",
                  "optional": true,
                  "default": 1,
                  "force_new": false
                }
              ],
              "block_types": [
                {
                  "type_name": "oauth2_permission_scope",
                  "block": {
                    "attributes": [
                      {
                        "name": "admin_consent_description",
                        "type": "string",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "admin_consent_display_name",
                        "type": "string",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "enabled",
                        "type": "bool",
                        "optional": true,
                        "default": true,
                        "force_new": false
                      },
                      {
                        "name": "id",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      },
                      {
                        "name": "type",
                        "type": "string",
                        "optional": true,
                        "default": "User",
                        "force_new": false
                      },
                      {
                        "name": "user_consent_description",
                        "type": "string",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "user_consent_display_name",
                        "type": "string",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "value",
                        "type": "string",
                        "optional": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 3,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "app_role",
            "block": {
              "attributes": [
                {
                  "name": "allowed_member_types",
                  "type": [
                    "set",
                    "string"
                  ],
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "description",
                  "type": "string",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "display_name",
                  "type": "string",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "enabled",
                  "type": "bool",
                  "optional": true,
                  "default": true,
                  "force_new": false
                },
                {
                  "name": "id",
                  "type": "string",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "value",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 3,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "feature_tags",
            "block": {
              "attributes": [
                {
                  "name": "custom_single_sign_on",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "enterprise",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "gallery",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "hide",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "required": false,
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "tags"
            ]
          },
          {
            "type_name": "optional_claims",
            "block": {
              "block_types": [
                {
                  "type_name": "access_token",
                  "block": {
                    "attributes": [
                      {
                        "name": "additional_properties",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "essential",
                        "type": "bool",
                        "optional": true,
                        "default": false,
                        "force_new": false
                      },
                      {
                        "name": "name",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      },
                      {
                        "name": "source",
                        "type": "string",
                        "optional": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                },
                {
                  "type_name": "id_token",
                  "block": {
                    "attributes": [
                      {
                        "name": "additional_properties",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "essential",
                        "type": "bool",
                        "optional": true,
                        "default": false,
                        "force_new": false
                      },
                      {
                        "name": "name",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      },
                      {
                        "name": "source",
                        "type": "string",
                        "optional": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                },
                {
                  "type_name": "saml2_token",
                  "block": {
                    "attributes": [
                      {
                        "name": "additional_properties",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "essential",
                        "type": "bool",
                        "optional": true,
                        "default": false,
                        "force_new": false
                      },
                      {
                        "name": "name",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      },
                      {
                        "name": "source",
                        "type": "string",
                        "optional": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "password",
            "block": {
              "attributes": [
                {
                  "name": "display_name",
                  "type": "string",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "end_date",
                  "type": "string",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "key_id",
                  "type": "string",
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "start_date",
                  "type": "string",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "value",
                  "type": "string",
                  "computed": true,
                  "sensitive": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 3,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "public_client",
            "block": {
              "attributes": [
                {
                  "name": "redirect_uris",
                  "type": [
                    "set",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "required_resource_access",
            "block": {
              "attributes": [
                {
                  "name": "resource_app_id",
                  "type": "string",
                  "required": true,
                  "force_new": false
                }
              ],
              "block_types": [
                {
                  "type_name": "resource_access",
                  "block": {
                    "attributes": [
                      {
                        "name": "id",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      },
                      {
                        "name": "type",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "min_items": 1,
                  "required": true,
                  "optional": false,
                  "computed": false,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 3,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "single_page_application",
            "block": {
              "attributes": [
                {
                  "name": "redirect_uris",
                  "type": [
                    "set",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "web",
            "block": {
              "attributes": [
                {
                  "name": "homepage_url",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "logout_url",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "redirect_uris",
                  "type": [
                    "set",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                }
              ],
              "block_types": [
                {
                  "type_name": "implicit_grant",
                  "block": {
                    "attributes": [
                      {
                        "name": "access_token_issuance_enabled",
                        "type": "bool",
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "id_token_issuance_enabled",
                        "type": "bool",
                        "optional": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "max_items": 1,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          }
        ]
      }
    },
    "azuread_application_api_access": {
      "block": {
        "attributes": [
          {
            "name": "api_client_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "role_ids",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false,
            "at_least_one_of": [
              "role_ids",
              "scope_ids"
            ]
          },
          {
            "name": "scope_ids",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false,
            "at_least_one_of": [
              "role_ids",
              "scope_ids"
            ]
          }
        ]
      }
    },
    "azuread_application_app_role": {
      "block": {
        "attributes": [
          {
            "name": "allowed_member_types",
            "type": [
              "set",
              "string"
            ],
            "required": true,
            "force_new": false
          },
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "description",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "role_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "value",
            "type": "string",
            "optional": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_application_certificate": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "encoding",
            "type": "string",
            "optional": true,
            "default": "pem",
            "force_new": true
          },
          {
            "name": "end_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "conflicts_with": [
              "end_date_relative"
            ]
          },
          {
            "name": "end_date_relative",
            "type": "string",
            "optional": true,
            "force_new": true,
            "conflicts_with": [
              "end_date"
            ]
          },
          {
            "name": "key_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "start_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "type",
            "type": "string",
            "optional": true,
            "force_new": true
          },
          {
            "name": "value",
            "type": "string",
            "required": true,
            "sensitive": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_application_fallback_public_client": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "enabled",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": true
          }
        ]
      }
    },
    "azuread_application_federated_identity_credential": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "audiences",
            "type": [
              "list",
              "string"
            ],
            "required": true,
            "force_new": false
          },
          {
            "name": "credential_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "description",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "issuer",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "subject",
            "type": "string",
            "required": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_application_from_template": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "application_object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "service_principal_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "service_principal_object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "template_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_application_identifier_uri": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "identifier_uri",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_application_known_clients": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "known_client_ids",
            "type": [
              "set",
              "string"
            ],
            "required": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_application_optional_claims": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ],
        "block_types": [
          {
            "type_name": "access_token",
            "block": {
              "attributes": [
                {
                  "name": "additional_properties",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "essential",
                  "type": "bool",
                  "optional": true,
                  "default": false,
                  "force_new": false
                },
                {
                  "name": "name",
                  "type": "string",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "source",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false,
            "at_least_one_of": [
              "access_token",
              "id_token",
              "saml2_token"
            ]
          },
          {
            "type_name": "id_token",
            "block": {
              "attributes": [
                {
                  "name": "additional_properties",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "essential",
                  "type": "bool",
                  "optional": true,
                  "default": false,
                  "force_new": false
                },
                {
                  "name": "name",
                  "type": "string",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "source",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false,
            "at_least_one_of": [
              "access_token",
              "id_token",
              "saml2_token"
            ]
          },
          {
            "type_name": "saml2_token",
            "block": {
              "attributes": [
                {
                  "name": "additional_properties",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "essential",
                  "type": "bool",
                  "optional": true,
                  "default": false,
                  "force_new": false
                },
                {
                  "name": "name",
                  "type": "string",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "source",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false,
            "at_least_one_of": [
              "access_token",
              "id_token",
              "saml2_token"
            ]
          }
        ]
      }
    },
    "azuread_application_owner": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "owner_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_application_password": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "end_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "conflicts_with": [
              "end_date_relative"
            ]
          },
          {
            "name": "end_date_relative",
            "type": "string",
            "optional": true,
            "force_new": true,
            "conflicts_with": [
              "end_date"
            ]
          },
          {
            "name": "key_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "rotate_when_changed",
            "type": [
              "map",
              "string"
            ],
            "optional": true,
            "force_new": true
          },
          {
            "name": "start_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "value",
            "type": "string",
            "computed": true,
            "sensitive": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_application_permission_scope": {
      "block": {
        "attributes": [
          {
            "name": "admin_consent_description",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "admin_consent_display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "scope_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "type",
            "type": "string",
            "optional": true,
            "default": "User",
            "force_new": false
          },
          {
            "name": "user_consent_description",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "user_consent_display_name",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "value",
            "type": "string",
            "required": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_application_pre_authorized": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "authorized_client_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "permission_ids",
            "type": [
              "set",
              "string"
            ],
            "required": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_application_redirect_uris": {
      "block": {
        "attributes": [
          {
            "name": "application_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "redirect_uris",
            "type": [
              "set",
              "string"
            ],
            "required": true,
            "force_new": false
          },
          {
            "name": "type",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_application_registration": {
      "block": {
        "attributes": [
          {
            "name": "client_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "description",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "disabled_by_microsoft",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "group_membership_claims",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false
          },
          {
            "name": "homepage_url",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "implicit_access_token_issuance_enabled",
            "type": "bool",
            "optional": true,
            "force_new": false
          },
          {
            "name": "implicit_id_token_issuance_enabled",
            "type": "bool",
            "optional": true,
            "force_new": false
          },
          {
            "name": "logout_url",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "marketing_url",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "notes",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "privacy_statement_url",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "publisher_domain",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "requested_access_token_version",
            "type": "number",
            "optional": true,
            "default": 2,
            "force_new": false
          },
          {
            "name": "service_management_reference",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "sign_in_audience",
            "type": "string",
            "optional": true,
            "default": "AzureADMyOrg",
            "force_new": false
          },
          {
            "name": "support_url",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "terms_of_service_url",
            "type": "string",
            "optional": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_authentication_strength_policy": {
      "block": {
        "attributes": [
          {
            "name": "allowed_combinations",
            "type": [
              "set",
              "string"
            ],
            "required": true,
            "force_new": false
          },
          {
            "name": "description",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_claims_mapping_policy": {
      "block": {
        "attributes": [
          {
            "name": "definition",
            "type": [
              "list",
              "string"
            ],
            "required": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_conditional_access_policy": {
      "block": {
        "attributes": [
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "state",
            "type": "string",
            "required": true,
            "force_new": false
          }
        ],
        "block_types": [
          {
            "type_name": "conditions",
            "block": {
              "attributes": [
                {
                  "name": "client_app_types",
                  "type": [
                    "list",
                    "string"
                  ],
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "insider_risk_levels",
                  "type": "string",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "service_principal_risk_levels",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "sign_in_risk_levels",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "user_risk_levels",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                }
              ],
              "block_types": [
                {
                  "type_name": "applications",
                  "block": {
                    "attributes": [
                      {
                        "name": "excluded_applications",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "included_applications",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false,
                        "exactly_one_of": [
                          "conditions.0.applications.0.included_applications",
                          "conditions.0.applications.0.included_user_actions"
                        ]
                      },
                      {
                        "name": "included_user_actions",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false,
                        "exactly_one_of": [
                          "conditions.0.applications.0.included_applications",
                          "conditions.0.applications.0.included_user_actions"
                        ]
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "min_items": 1,
                  "max_items": 1,
                  "required": true,
                  "optional": false,
                  "computed": false,
                  "force_new": false
                },
                {
                  "type_name": "client_applications",
                  "block": {
                    "attributes": [
                      {
                        "name": "excluded_service_principals",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "included_service_principals",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "max_items": 1,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                },
                {
                  "type_name": "devices",
                  "block": {
                    "block_types": [
                      {
                        "type_name": "filter",
                        "block": {
                          "attributes": [
                            {
                              "name": "mode",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "rule",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "max_items": 1,
                        "required": false,
                        "optional": true,
                        "computed": false,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "max_items": 1,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                },
                {
                  "type_name": "locations",
                  "block": {
                    "attributes": [
                      {
                        "name": "excluded_locations",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "included_locations",
                        "type": [
                          "list",
                          "string"
                        ],
                        "required": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "max_items": 1,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                },
                {
                  "type_name": "platforms",
                  "block": {
                    "attributes": [
                      {
                        "name": "excluded_platforms",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "included_platforms",
                        "type": [
                          "list",
                          "string"
                        ],
                        "required": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "max_items": 1,
                  "required": false,
                  "optional": true,
                  "computed": false,
                  "force_new": false
                },
                {
                  "type_name": "users",
                  "block": {
                    "attributes": [
                      {
                        "name": "excluded_groups",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "excluded_roles",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "excluded_users",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false
                      },
                      {
                        "name": "included_groups",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false,
                        "at_least_one_of": [
                          "conditions.0.users.0.included_groups",
                          "conditions.0.users.0.included_roles",
                          "conditions.0.users.0.included_users",
                          "conditions.0.users.0.included_guests_or_external_users"
                        ]
                      },
                      {
                        "name": "included_roles",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false,
                        "at_least_one_of": [
                          "conditions.0.users.0.included_groups",
                          "conditions.0.users.0.included_roles",
                          "conditions.0.users.0.included_users",
                          "conditions.0.users.0.included_guests_or_external_users"
                        ]
                      },
                      {
                        "name": "included_users",
                        "type": [
                          "list",
                          "string"
                        ],
                        "optional": true,
                        "force_new": false,
                        "at_least_one_of": [
                          "conditions.0.users.0.included_groups",
                          "conditions.0.users.0.included_roles",
                          "conditions.0.users.0.included_users",
                          "conditions.0.users.0.included_guests_or_external_users"
                        ]
                      }
                    ],
                    "block_types": [
                      {
                        "type_name": "excluded_guests_or_external_users",
                        "block": {
                          "attributes": [
                            {
                              "name": "guest_or_external_user_types",
                              "type": [
                                "list",
                                "string"
                              ],
                              "required": true,
                              "force_new": false
                            }
                          ],
                          "block_types": [
                            {
                              "type_name": "external_tenants",
                              "block": {
                                "attributes": [
                                  {
                                    "name": "members",
                                    "type": [
                                      "list",
                                      "string"
                                    ],
                                    "optional": true,
                                    "force_new": false
                                  },
                                  {
                                    "name": "membership_kind",
                                    "type": "string",
                                    "required": true,
                                    "force_new": false
                                  }
                                ]
                              },
                              "nesting_mode": 2,
                              "required": false,
                              "optional": true,
                              "computed": false,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "required": false,
                        "optional": true,
                        "computed": false,
                        "force_new": false
                      },
                      {
                        "type_name": "included_guests_or_external_users",
                        "block": {
                          "attributes": [
                            {
                              "name": "guest_or_external_user_types",
                              "type": [
                                "list",
                                "string"
                              ],
                              "required": true,
                              "force_new": false
                            }
                          ],
                          "block_types": [
                            {
                              "type_name": "external_tenants",
                              "block": {
                                "attributes": [
                                  {
                                    "name": "members",
                                    "type": [
                                      "list",
                                      "string"
                                    ],
                                    "optional": true,
                                    "force_new": false
                                  },
                                  {
                                    "name": "membership_kind",
                                    "type": "string",
                                    "required": true,
                                    "force_new": false
                                  }
                                ]
                              },
                              "nesting_mode": 2,
                              "required": false,
                              "optional": true,
                              "computed": false,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "required": false,
                        "optional": true,
                        "computed": false,
                        "force_new": false,
                        "at_least_one_of": [
                          "conditions.0.users.0.included_groups",
                          "conditions.0.users.0.included_roles",
                          "conditions.0.users.0.included_users",
                          "conditions.0.users.0.included_guests_or_external_users"
                        ]
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "min_items": 1,
                  "max_items": 1,
                  "required": true,
                  "optional": false,
                  "computed": false,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "min_items": 1,
            "max_items": 1,
            "required": true,
            "optional": false,
            "computed": false,
            "force_new": false
          },
          {
            "type_name": "grant_controls",
            "block": {
              "attributes": [
                {
                  "name": "authentication_strength_policy_id",
                  "type": "string",
                  "optional": true,
                  "force_new": false,
                  "at_least_one_of": [
                    "grant_controls.0.built_in_controls",
                    "grant_controls.0.authentication_strength_policy_id",
                    "grant_controls.0.terms_of_use"
                  ]
                },
                {
                  "name": "built_in_controls",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false,
                  "at_least_one_of": [
                    "grant_controls.0.built_in_controls",
                    "grant_controls.0.authentication_strength_policy_id",
                    "grant_controls.0.terms_of_use"
                  ]
                },
                {
                  "name": "custom_authentication_factors",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "operator",
                  "type": "string",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "terms_of_use",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false,
                  "at_least_one_of": [
                    "grant_controls.0.built_in_controls",
                    "grant_controls.0.authentication_strength_policy_id",
                    "grant_controls.0.terms_of_use"
                  ]
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false,
            "at_least_one_of": [
              "grant_controls",
              "session_controls"
            ]
          },
          {
            "type_name": "session_controls",
            "block": {
              "attributes": [
                {
                  "name": "application_enforced_restrictions_enabled",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "cloud_app_security_policy",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "disable_resilience_defaults",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "persistent_browser_mode",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "sign_in_frequency",
                  "type": "number",
                  "optional": true,
                  "force_new": false,
                  "required_with": [
                    "session_controls.0.sign_in_frequency_period"
                  ]
                },
                {
                  "name": "sign_in_frequency_authentication_type",
                  "type": "string",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "sign_in_frequency_interval",
                  "type": "string",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "sign_in_frequency_period",
                  "type": "string",
                  "optional": true,
                  "force_new": false,
                  "required_with": [
                    "session_controls.0.sign_in_frequency"
                  ]
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false,
            "at_least_one_of": [
              "grant_controls",
              "session_controls"
            ]
          }
        ]
      }
    },
    "azuread_custom_directory_role": {
      "block": {
        "attributes": [
          {
            "name": "description",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "enabled",
            "type": "bool",
            "required": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "template_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "version",
            "type": "string",
            "required": true,
            "force_new": false
          }
        ],
        "block_types": [
          {
            "type_name": "permissions",
            "block": {
              "attributes": [
                {
                  "name": "allowed_resource_actions",
                  "type": [
                    "set",
                    "string"
                  ],
                  "required": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 3,
            "min_items": 1,
            "required": true,
            "optional": false,
            "computed": false,
            "force_new": false
          }
        ]
      }
    },
    "azuread_directory_role": {
      "block": {
        "attributes": [
          {
            "name": "description",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "exactly_one_of": [
              "display_name",
              "template_id"
            ]
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "template_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "exactly_one_of": [
              "display_name",
              "template_id"
            ]
          }
        ]
      }
    },
    "azuread_directory_role_assignment": {
      "block": {
        "attributes": [
          {
            "name": "app_scope_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "conflicts_with": [
              "directory_scope_id"
            ]
          },
          {
            "name": "directory_scope_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "conflicts_with": [
              "app_scope_id"
            ]
          },
          {
            "name": "principal_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "role_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_directory_role_eligibility_schedule_request": {
      "block": {
        "attributes": [
          {
            "name": "directory_scope_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "justification",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "principal_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "role_definition_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_directory_role_member": {
      "block": {
        "attributes": [
          {
            "name": "member_object_id",
            "type": "string",
            "optional": true,
            "force_new": true
          },
          {
            "name": "role_object_id",
            "type": "string",
            "optional": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_group": {
      "block": {
        "attributes": [
          {
            "name": "administrative_unit_ids",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false
          },
          {
            "name": "assignable_to_role",
            "type": "bool",
            "optional": true,
            "force_new": true
          },
          {
            "name": "auto_subscribe_new_members",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "behaviors",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": true
          },
          {
            "name": "description",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "external_senders_allowed",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "hide_from_address_lists",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "hide_from_outlook_clients",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "mail",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "mail_enabled",
            "type": "bool",
            "optional": true,
            "force_new": false,
            "at_least_one_of": [
              "mail_enabled",
              "security_enabled"
            ]
          },
          {
            "name": "mail_nickname",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "members",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "dynamic_membership"
            ]
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_domain_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_group_type",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_netbios_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_sam_account_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_security_identifier",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_sync_enabled",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "owners",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "preferred_language",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "prevent_duplicate_names",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          },
          {
            "name": "provisioning_options",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": true
          },
          {
            "name": "proxy_addresses",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "security_enabled",
            "type": "bool",
            "optional": true,
            "force_new": false,
            "at_least_one_of": [
              "mail_enabled",
              "security_enabled"
            ]
          },
          {
            "name": "theme",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "types",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": true
          },
          {
            "name": "visibility",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "writeback_enabled",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          }
        ],
        "block_types": [
          {
            "type_name": "dynamic_membership",
            "block": {
              "attributes": [
                {
                  "name": "enabled",
                  "type": "bool",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "rule",
                  "type": "string",
                  "required": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false,
            "conflicts_with": [
              "members"
            ]
          }
        ]
      }
    },
    "azuread_group_member": {
      "block": {
        "attributes": [
          {
            "name": "group_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "member_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_group_role_management_policy": {
      "block": {
        "attributes": [
          {
            "name": "description",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "group_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "role_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ],
        "block_types": [
          {
            "type_name": "activation_rules",
            "block": {
              "attributes": [
                {
                  "name": "maximum_duration",
                  "type": "string",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "require_approval",
                  "type": "bool",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "require_justification",
                  "type": "bool",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "require_multifactor_authentication",
                  "type": "bool",
                  "optional": true,
                  "computed": true,
                  "force_new": false,
                  "conflicts_with": [
                    "activation_rules.0.required_conditional_access_authentication_context"
                  ]
                },
                {
                  "name": "require_ticket_info",
                  "type": "bool",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "required_conditional_access_authentication_context",
                  "type": "string",
                  "optional": true,
                  "computed": true,
                  "force_new": false,
                  "conflicts_with": [
                    "activation_rules.0.require_multifactor_authentication"
                  ]
                }
              ],
              "block_types": [
                {
                  "type_name": "approval_stage",
                  "block": {
                    "block_types": [
                      {
                        "type_name": "primary_approver",
                        "block": {
                          "attributes": [
                            {
                              "name": "object_id",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "type",
                              "type": "string",
                              "optional": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 3,
                        "min_items": 1,
                        "required": true,
                        "optional": false,
                        "computed": false,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "max_items": 1,
                  "required": false,
                  "optional": true,
                  "computed": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "type_name": "active_assignment_rules",
            "block": {
              "attributes": [
                {
                  "name": "expiration_required",
                  "type": "bool",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "expire_after",
                  "type": "string",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "require_justification",
                  "type": "bool",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "require_multifactor_authentication",
                  "type": "bool",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "require_ticket_info",
                  "type": "bool",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "type_name": "eligible_assignment_rules",
            "block": {
              "attributes": [
                {
                  "name": "expiration_required",
                  "type": "bool",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "name": "expire_after",
                  "type": "string",
                  "optional": true,
                  "computed": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "type_name": "notification_rules",
            "block": {
              "block_types": [
                {
                  "type_name": "active_assignments",
                  "block": {
                    "block_types": [
                      {
                        "type_name": "admin_notifications",
                        "block": {
                          "attributes": [
                            {
                              "name": "additional_recipients",
                              "type": [
                                "set",
                                "string"
                              ],
                              "optional": true,
                              "computed": true,
                              "force_new": false
                            },
                            {
                              "name": "default_recipients",
                              "type": "bool",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "notification_level",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "max_items": 1,
                        "required": false,
                        "optional": true,
                        "computed": true,
                        "force_new": false
                      },
                      {
                        "type_name": "approver_notifications",
                        "block": {
                          "attributes": [
                            {
                              "name": "additional_recipients",
                              "type": [
                                "set",
                                "string"
                              ],
                              "optional": true,
                              "computed": true,
                              "force_new": false
                            },
                            {
                              "name": "default_recipients",
                              "type": "bool",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "notification_level",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "max_items": 1,
                        "required": false,
                        "optional": true,
                        "computed": true,
                        "force_new": false
                      },
                      {
                        "type_name": "assignee_notifications",
                        "block": {
                          "attributes": [
                            {
                              "name": "additional_recipients",
                              "type": [
                                "set",
                                "string"
                              ],
                              "optional": true,
                              "computed": true,
                              "force_new": false
                            },
                            {
                              "name": "default_recipients",
                              "type": "bool",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "notification_level",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "max_items": 1,
                        "required": false,
                        "optional": true,
                        "computed": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "max_items": 1,
                  "required": false,
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "type_name": "eligible_activations",
                  "block": {
                    "block_types": [
                      {
                        "type_name": "admin_notifications",
                        "block": {
                          "attributes": [
                            {
                              "name": "additional_recipients",
                              "type": [
                                "set",
                                "string"
                              ],
                              "optional": true,
                              "computed": true,
                              "force_new": false
                            },
                            {
                              "name": "default_recipients",
                              "type": "bool",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "notification_level",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "max_items": 1,
                        "required": false,
                        "optional": true,
                        "computed": true,
                        "force_new": false
                      },
                      {
                        "type_name": "approver_notifications",
                        "block": {
                          "attributes": [
                            {
                              "name": "additional_recipients",
                              "type": [
                                "set",
                                "string"
                              ],
                              "optional": true,
                              "computed": true,
                              "force_new": false
                            },
                            {
                              "name": "default_recipients",
                              "type": "bool",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "notification_level",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "max_items": 1,
                        "required": false,
                        "optional": true,
                        "computed": true,
                        "force_new": false
                      },
                      {
                        "type_name": "assignee_notifications",
                        "block": {
                          "attributes": [
                            {
                              "name": "additional_recipients",
                              "type": [
                                "set",
                                "string"
                              ],
                              "optional": true,
                              "computed": true,
                              "force_new": false
                            },
                            {
                              "name": "default_recipients",
                              "type": "bool",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "notification_level",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "max_items": 1,
                        "required": false,
                        "optional": true,
                        "computed": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "max_items": 1,
                  "required": false,
                  "optional": true,
                  "computed": true,
                  "force_new": false
                },
                {
                  "type_name": "eligible_assignments",
                  "block": {
                    "block_types": [
                      {
                        "type_name": "admin_notifications",
                        "block": {
                          "attributes": [
                            {
                              "name": "additional_recipients",
                              "type": [
                                "set",
                                "string"
                              ],
                              "optional": true,
                              "computed": true,
                              "force_new": false
                            },
                            {
                              "name": "default_recipients",
                              "type": "bool",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "notification_level",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "max_items": 1,
                        "required": false,
                        "optional": true,
                        "computed": true,
                        "force_new": false
                      },
                      {
                        "type_name": "approver_notifications",
                        "block": {
                          "attributes": [
                            {
                              "name": "additional_recipients",
                              "type": [
                                "set",
                                "string"
                              ],
                              "optional": true,
                              "computed": true,
                              "force_new": false
                            },
                            {
                              "name": "default_recipients",
                              "type": "bool",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "notification_level",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "max_items": 1,
                        "required": false,
                        "optional": true,
                        "computed": true,
                        "force_new": false
                      },
                      {
                        "type_name": "assignee_notifications",
                        "block": {
                          "attributes": [
                            {
                              "name": "additional_recipients",
                              "type": [
                                "set",
                                "string"
                              ],
                              "optional": true,
                              "computed": true,
                              "force_new": false
                            },
                            {
                              "name": "default_recipients",
                              "type": "bool",
                              "required": true,
                              "force_new": false
                            },
                            {
                              "name": "notification_level",
                              "type": "string",
                              "required": true,
                              "force_new": false
                            }
                          ]
                        },
                        "nesting_mode": 2,
                        "max_items": 1,
                        "required": false,
                        "optional": true,
                        "computed": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "max_items": 1,
                  "required": false,
                  "optional": true,
                  "computed": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_group_without_members": {
      "block": {
        "attributes": [
          {
            "name": "administrative_unit_ids",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false
          },
          {
            "name": "assignable_to_role",
            "type": "bool",
            "optional": true,
            "force_new": true
          },
          {
            "name": "auto_subscribe_new_members",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "behaviors",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": true
          },
          {
            "name": "description",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "external_senders_allowed",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "hide_from_address_lists",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "hide_from_outlook_clients",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "mail",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "mail_enabled",
            "type": "bool",
            "optional": true,
            "force_new": false,
            "at_least_one_of": [
              "mail_enabled",
              "security_enabled"
            ]
          },
          {
            "name": "mail_nickname",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_domain_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_group_type",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_netbios_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_sam_account_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_security_identifier",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_sync_enabled",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "owners",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "preferred_language",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "prevent_duplicate_names",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          },
          {
            "name": "provisioning_options",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": true
          },
          {
            "name": "proxy_addresses",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "security_enabled",
            "type": "bool",
            "optional": true,
            "force_new": false,
            "at_least_one_of": [
              "mail_enabled",
              "security_enabled"
            ]
          },
          {
            "name": "theme",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "types",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": true
          },
          {
            "name": "visibility",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "writeback_enabled",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          }
        ],
        "block_types": [
          {
            "type_name": "dynamic_membership",
            "block": {
              "attributes": [
                {
                  "name": "enabled",
                  "type": "bool",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "rule",
                  "type": "string",
                  "required": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          }
        ]
      }
    },
    "azuread_invitation": {
      "block": {
        "attributes": [
          {
            "name": "redeem_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "redirect_url",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "user_display_name",
            "type": "string",
            "optional": true,
            "force_new": true
          },
          {
            "name": "user_email_address",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "user_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "user_type",
            "type": "string",
            "optional": true,
            "default": "Guest",
            "force_new": true
          }
        ],
        "block_types": [
          {
            "type_name": "message",
            "block": {
              "attributes": [
                {
                  "name": "additional_recipients",
                  "type": [
                    "list",
                    "string"
                  ],
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "body",
                  "type": "string",
                  "optional": true,
                  "force_new": false,
                  "conflicts_with": [
                    "message.0.language"
                  ]
                },
                {
                  "name": "language",
                  "type": "string",
                  "optional": true,
                  "force_new": false,
                  "conflicts_with": [
                    "message.0.body"
                  ]
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": true
          }
        ]
      }
    },
    "azuread_named_location": {
      "block": {
        "attributes": [
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          }
        ],
        "block_types": [
          {
            "type_name": "country",
            "block": {
              "attributes": [
                {
                  "name": "countries_and_regions",
                  "type": [
                    "list",
                    "string"
                  ],
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "country_lookup_method",
                  "type": "string",
                  "optional": true,
                  "default": "clientIpAddress",
                  "force_new": false
                },
                {
                  "name": "include_unknown_countries_and_regions",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": true,
            "exactly_one_of": [
              "ip",
              "country"
            ]
          },
          {
            "type_name": "ip",
            "block": {
              "attributes": [
                {
                  "name": "ip_ranges",
                  "type": [
                    "list",
                    "string"
                  ],
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "trusted",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": true,
            "exactly_one_of": [
              "ip",
              "country"
            ]
          }
        ]
      }
    },
    "azuread_privileged_access_group_assignment_schedule": {
      "block": {
        "attributes": [
          {
            "name": "assignment_type",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "duration",
            "type": "string",
            "optional": true,
            "force_new": false,
            "conflicts_with": [
              "expiration_date"
            ]
          },
          {
            "name": "expiration_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "duration"
            ]
          },
          {
            "name": "group_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "justification",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "permanent_assignment",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "principal_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "start_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "status",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "ticket_number",
            "type": "string",
            "optional": true,
            "force_new": false,
            "required_with": [
              "ticket_system"
            ]
          },
          {
            "name": "ticket_system",
            "type": "string",
            "optional": true,
            "force_new": false,
            "required_with": [
              "ticket_number"
            ]
          }
        ]
      }
    },
    "azuread_privileged_access_group_eligibility_schedule": {
      "block": {
        "attributes": [
          {
            "name": "assignment_type",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "duration",
            "type": "string",
            "optional": true,
            "force_new": false,
            "conflicts_with": [
              "expiration_date"
            ]
          },
          {
            "name": "expiration_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "duration"
            ]
          },
          {
            "name": "group_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "justification",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "permanent_assignment",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "principal_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "start_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "status",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "ticket_number",
            "type": "string",
            "optional": true,
            "force_new": false,
            "required_with": [
              "ticket_system"
            ]
          },
          {
            "name": "ticket_system",
            "type": "string",
            "optional": true,
            "force_new": false,
            "required_with": [
              "ticket_number"
            ]
          }
        ]
      }
    },
    "azuread_service_principal": {
      "block": {
        "attributes": [
          {
            "name": "account_enabled",
            "type": "bool",
            "optional": true,
            "default": true,
            "force_new": false
          },
          {
            "name": "alternative_names",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false
          },
          {
            "name": "app_role_assignment_required",
            "type": "bool",
            "optional": true,
            "force_new": false
          },
          {
            "name": "app_role_ids",
            "type": [
              "map",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "app_roles",
            "type": [
              "list",
              [
                "object",
                {
                  "allowed_member_types": [
                    "list",
                    "string"
                  ],
                  "description": "string",
                  "display_name": "string",
                  "enabled": "bool",
                  "id": "string",
                  "value": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "application_tenant_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "client_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "description",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "homepage_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "login_url",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "logout_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "notes",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "notification_email_addresses",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false
          },
          {
            "name": "oauth2_permission_scope_ids",
            "type": [
              "map",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "oauth2_permission_scopes",
            "type": [
              "list",
              [
                "object",
                {
                  "admin_consent_description": "string",
                  "admin_consent_display_name": "string",
                  "enabled": "bool",
                  "id": "string",
                  "type": "string",
                  "user_consent_description": "string",
                  "user_consent_display_name": "string",
                  "value": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "owners",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false
          },
          {
            "name": "preferred_single_sign_on_mode",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "redirect_uris",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "saml_metadata_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "service_principal_names",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "sign_in_audience",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "tags",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "features",
              "feature_tags"
            ]
          },
          {
            "name": "type",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "use_existing",
            "type": "bool",
            "optional": true,
            "force_new": false
          }
        ],
        "block_types": [
          {
            "type_name": "feature_tags",
            "block": {
              "attributes": [
                {
                  "name": "custom_single_sign_on",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "enterprise",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "gallery",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "hide",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "required": false,
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "features",
              "tags"
            ]
          },
          {
            "type_name": "features",
            "block": {
              "attributes": [
                {
                  "name": "custom_single_sign_on_app",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "enterprise_application",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "gallery_application",
                  "type": "bool",
                  "optional": true,
                  "force_new": false
                },
                {
                  "name": "visible_to_users",
                  "type": "bool",
                  "optional": true,
                  "default": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "required": false,
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "feature_tags",
              "tags"
            ]
          },
          {
            "type_name": "saml_single_sign_on",
            "block": {
              "attributes": [
                {
                  "name": "relay_state",
                  "type": "string",
                  "optional": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "max_items": 1,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          }
        ]
      }
    },
    "azuread_service_principal_certificate": {
      "block": {
        "attributes": [
          {
            "name": "encoding",
            "type": "string",
            "optional": true,
            "default": "pem",
            "force_new": true
          },
          {
            "name": "end_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "conflicts_with": [
              "end_date_relative"
            ]
          },
          {
            "name": "end_date_relative",
            "type": "string",
            "optional": true,
            "force_new": true,
            "conflicts_with": [
              "end_date"
            ]
          },
          {
            "name": "key_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "service_principal_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "start_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "type",
            "type": "string",
            "optional": true,
            "force_new": true
          },
          {
            "name": "value",
            "type": "string",
            "required": true,
            "sensitive": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_service_principal_claims_mapping_policy_assignment": {
      "block": {
        "attributes": [
          {
            "name": "claims_mapping_policy_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "service_principal_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_service_principal_delegated_permission_grant": {
      "block": {
        "attributes": [
          {
            "name": "claim_values",
            "type": [
              "set",
              "string"
            ],
            "required": true,
            "force_new": false
          },
          {
            "name": "resource_service_principal_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "service_principal_object_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "user_object_id",
            "type": "string",
            "optional": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_service_principal_password": {
      "block": {
        "attributes": [
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "end_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true,
            "conflicts_with": [
              "end_date_relative"
            ]
          },
          {
            "name": "end_date_relative",
            "type": "string",
            "optional": true,
            "force_new": true,
            "conflicts_with": [
              "end_date"
            ]
          },
          {
            "name": "key_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "rotate_when_changed",
            "type": [
              "map",
              "string"
            ],
            "optional": true,
            "force_new": true
          },
          {
            "name": "service_principal_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "start_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "value",
            "type": "string",
            "computed": true,
            "sensitive": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_service_principal_token_signing_certificate": {
      "block": {
        "attributes": [
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "end_date",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": true
          },
          {
            "name": "key_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "service_principal_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "start_date",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "thumbprint",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "value",
            "type": "string",
            "computed": true,
            "sensitive": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_synchronization_job": {
      "block": {
        "attributes": [
          {
            "name": "enabled",
            "type": "bool",
            "optional": true,
            "default": true,
            "force_new": false
          },
          {
            "name": "schedule",
            "type": [
              "list",
              [
                "object",
                {
                  "expiration": "string",
                  "interval": "string",
                  "state": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "service_principal_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "template_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    },
    "azuread_synchronization_job_provision_on_demand": {
      "block": {
        "attributes": [
          {
            "name": "service_principal_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "synchronization_job_id",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "triggers",
            "type": [
              "map",
              "string"
            ],
            "optional": true,
            "force_new": true
          }
        ],
        "block_types": [
          {
            "type_name": "parameter",
            "block": {
              "attributes": [
                {
                  "name": "rule_id",
                  "type": "string",
                  "required": true,
                  "force_new": true
                }
              ],
              "block_types": [
                {
                  "type_name": "subject",
                  "block": {
                    "attributes": [
                      {
                        "name": "object_id",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      },
                      {
                        "name": "object_type_name",
                        "type": "string",
                        "required": true,
                        "force_new": false
                      }
                    ]
                  },
                  "nesting_mode": 2,
                  "min_items": 1,
                  "required": true,
                  "optional": false,
                  "computed": false,
                  "force_new": true
                }
              ]
            },
            "nesting_mode": 2,
            "min_items": 1,
            "required": true,
            "optional": false,
            "computed": false,
            "force_new": true
          }
        ]
      }
    },
    "azuread_synchronization_secret": {
      "block": {
        "attributes": [
          {
            "name": "service_principal_id",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ],
        "block_types": [
          {
            "type_name": "credential",
            "block": {
              "attributes": [
                {
                  "name": "key",
                  "type": "string",
                  "required": true,
                  "force_new": false
                },
                {
                  "name": "value",
                  "type": "string",
                  "required": true,
                  "sensitive": true,
                  "force_new": false
                }
              ]
            },
            "nesting_mode": 2,
            "required": false,
            "optional": true,
            "computed": false,
            "force_new": false
          }
        ]
      }
    },
    "azuread_user": {
      "block": {
        "attributes": [
          {
            "name": "about_me",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "account_enabled",
            "type": "bool",
            "optional": true,
            "default": true,
            "force_new": false
          },
          {
            "name": "age_group",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "business_phones",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "city",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "company_name",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "consent_provided_for_minor",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "cost_center",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "country",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "creation_type",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "department",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "disable_password_expiration",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          },
          {
            "name": "disable_strong_password",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "division",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "employee_hire_date",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "employee_id",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "employee_type",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "external_user_state",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "fax_number",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "force_password_change",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          },
          {
            "name": "given_name",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "im_addresses",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "job_title",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "mail",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "mail_nickname",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "manager_id",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "mobile_phone",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "office_location",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "onpremises_distinguished_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_domain_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_immutable_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_sam_account_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_security_identifier",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_sync_enabled",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_user_principal_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "other_mails",
            "type": [
              "set",
              "string"
            ],
            "optional": true,
            "force_new": false
          },
          {
            "name": "password",
            "type": "string",
            "optional": true,
            "computed": true,
            "sensitive": true,
            "force_new": false
          },
          {
            "name": "postal_code",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "preferred_language",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "proxy_addresses",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "show_in_address_list",
            "type": "bool",
            "optional": true,
            "default": true,
            "force_new": false
          },
          {
            "name": "state",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "street_address",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "surname",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "usage_location",
            "type": "string",
            "optional": true,
            "force_new": false
          },
          {
            "name": "user_principal_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "user_type",
            "type": "string",
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_user_flow_attribute": {
      "block": {
        "attributes": [
          {
            "name": "attribute_type",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "data_type",
            "type": "string",
            "required": true,
            "force_new": true
          },
          {
            "name": "description",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": true
          }
        ]
      }
    }
  },
  "datasource_schemas": {
    "azuread_access_package": {
      "block": {
        "attributes": [
          {
            "name": "catalog_id",
            "type": "string",
            "optional": true,
            "force_new": false,
            "conflicts_with": [
              "object_id"
            ],
            "at_least_one_of": [
              "object_id",
              "display_name",
              "catalog_id"
            ],
            "required_with": [
              "display_name"
            ]
          },
          {
            "name": "description",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "object_id"
            ],
            "at_least_one_of": [
              "object_id",
              "display_name",
              "catalog_id"
            ],
            "required_with": [
              "catalog_id"
            ]
          },
          {
            "name": "hidden",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "at_least_one_of": [
              "object_id",
              "display_name",
              "catalog_id"
            ]
          }
        ]
      }
    },
    "azuread_access_package_catalog": {
      "block": {
        "attributes": [
          {
            "name": "description",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "object_id",
              "display_name"
            ]
          },
          {
            "name": "externally_visible",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "object_id",
              "display_name"
            ]
          },
          {
            "name": "published",
            "type": "bool",
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_access_package_catalog_role": {
      "block": {
        "attributes": [
          {
            "name": "description",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "display_name",
              "object_id"
            ]
          },
          {
            "name": "object_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "display_name",
              "object_id"
            ]
          },
          {
            "name": "template_id",
            "type": "string",
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_administrative_unit": {
      "block": {
        "attributes": [
          {
            "name": "description",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "members",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "visibility",
            "type": "string",
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_application": {
      "block": {
        "attributes": [
          {
            "name": "api",
            "type": [
              "list",
              [
                "object",
                {
                  "known_client_applications": [
                    "list",
                    "string"
                  ],
                  "mapped_claims_enabled": "bool",
                  "oauth2_permission_scopes": [
                    "list",
                    [
                      "object",
                      {
                        "admin_consent_description": "string",
                        "admin_consent_display_name": "string",
                        "enabled": "bool",
                        "id": "string",
                        "type": "string",
                        "user_consent_description": "string",
                        "user_consent_display_name": "string",
                        "value": "string"
                      }
                    ]
                  ],
                  "requested_access_token_version": "number"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "app_role_ids",
            "type": [
              "map",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "app_roles",
            "type": [
              "list",
              [
                "object",
                {
                  "allowed_member_types": [
                    "list",
                    "string"
                  ],
                  "description": "string",
                  "display_name": "string",
                  "enabled": "bool",
                  "id": "string",
                  "value": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "client_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "client_id",
              "display_name",
              "object_id",
              "identifier_uri"
            ]
          },
          {
            "name": "description",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "device_only_auth_enabled",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "disabled_by_microsoft",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "client_id",
              "display_name",
              "object_id",
              "identifier_uri"
            ]
          },
          {
            "name": "fallback_public_client_enabled",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "feature_tags",
            "type": [
              "list",
              [
                "object",
                {
                  "custom_single_sign_on": "bool",
                  "enterprise": "bool",
                  "gallery": "bool",
                  "hide": "bool"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "group_membership_claims",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "identifier_uri",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "client_id",
              "display_name",
              "object_id",
              "identifier_uri"
            ]
          },
          {
            "name": "identifier_uris",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "logo_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "marketing_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "notes",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "oauth2_permission_scope_ids",
            "type": [
              "map",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "oauth2_post_response_required",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "client_id",
              "display_name",
              "object_id",
              "identifier_uri"
            ]
          },
          {
            "name": "optional_claims",
            "type": [
              "list",
              [
                "object",
                {
                  "access_token": [
                    "list",
                    [
                      "object",
                      {
                        "additional_properties": [
                          "list",
                          "string"
                        ],
                        "essential": "bool",
                        "name": "string",
                        "source": "string"
                      }
                    ]
                  ],
                  "id_token": [
                    "list",
                    [
                      "object",
                      {
                        "additional_properties": [
                          "list",
                          "string"
                        ],
                        "essential": "bool",
                        "name": "string",
                        "source": "string"
                      }
                    ]
                  ],
                  "saml2_token": [
                    "list",
                    [
                      "object",
                      {
                        "additional_properties": [
                          "list",
                          "string"
                        ],
                        "essential": "bool",
                        "name": "string",
                        "source": "string"
                      }
                    ]
                  ]
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "owners",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "privacy_statement_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "public_client",
            "type": [
              "list",
              [
                "object",
                {
                  "redirect_uris": [
                    "list",
                    "string"
                  ]
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "publisher_domain",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "required_resource_access",
            "type": [
              "list",
              [
                "object",
                {
                  "resource_access": [
                    "list",
                    [
                      "object",
                      {
                        "id": "string",
                        "type": "string"
                      }
                    ]
                  ],
                  "resource_app_id": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "service_management_reference",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "sign_in_audience",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "single_page_application",
            "type": [
              "list",
              [
                "object",
                {
                  "redirect_uris": [
                    "list",
                    "string"
                  ]
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "support_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "tags",
            "type": [
              "set",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "terms_of_service_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "web",
            "type": [
              "list",
              [
                "object",
                {
                  "homepage_url": "string",
                  "implicit_grant": [
                    "list",
                    [
                      "object",
                      {
                        "access_token_issuance_enabled": "bool",
                        "id_token_issuance_enabled": "bool"
                      }
                    ]
                  ],
                  "logout_url": "string",
                  "redirect_uris": [
                    "list",
                    "string"
                  ]
                }
              ]
            ],
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_application_published_app_ids": {
      "block": {
        "attributes": [
          {
            "name": "result",
            "type": [
              "map",
              "string"
            ],
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_application_template": {
      "block": {
        "attributes": [
          {
            "name": "categories",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "display_name",
              "template_id"
            ]
          },
          {
            "name": "homepage_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "logo_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "publisher",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "supported_provisioning_types",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "supported_single_sign_on_modes",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "template_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "display_name",
              "template_id"
            ]
          }
        ]
      }
    },
    "azuread_client_config": {
      "block": {
        "attributes": [
          {
            "name": "client_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "tenant_id",
            "type": "string",
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_directory_object": {
      "block": {
        "attributes": [
          {
            "name": "object_id",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "type",
            "type": "string",
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_directory_role_templates": {
      "block": {
        "attributes": [
          {
            "name": "object_ids",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "role_templates",
            "type": [
              "list",
              [
                "object",
                {
                  "description": "string",
                  "display_name": "string",
                  "object_id": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_directory_roles": {
      "block": {
        "attributes": [
          {
            "name": "object_ids",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "roles",
            "type": [
              "list",
              [
                "object",
                {
                  "description": "string",
                  "display_name": "string",
                  "object_id": "string",
                  "template_id": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "template_ids",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_domains": {
      "block": {
        "attributes": [
          {
            "name": "admin_managed",
            "type": "bool",
            "optional": true,
            "force_new": false
          },
          {
            "name": "domains",
            "type": [
              "list",
              [
                "object",
                {
                  "admin_managed": "bool",
                  "authentication_type": "string",
                  "default": "bool",
                  "domain_name": "string",
                  "initial": "bool",
                  "root": "bool",
                  "supported_services": [
                    "list",
                    "string"
                  ],
                  "verified": "bool"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "include_unverified",
            "type": "bool",
            "optional": true,
            "force_new": false,
            "conflicts_with": [
              "only_default",
              "only_initial"
            ]
          },
          {
            "name": "only_default",
            "type": "bool",
            "optional": true,
            "force_new": false,
            "conflicts_with": [
              "only_initial",
              "only_root"
            ]
          },
          {
            "name": "only_initial",
            "type": "bool",
            "optional": true,
            "force_new": false,
            "conflicts_with": [
              "only_default",
              "only_root"
            ]
          },
          {
            "name": "only_root",
            "type": "bool",
            "optional": true,
            "force_new": false,
            "conflicts_with": [
              "only_default",
              "only_initial"
            ]
          },
          {
            "name": "supports_services",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_group": {
      "block": {
        "attributes": [
          {
            "name": "assignable_to_role",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "auto_subscribe_new_members",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "behaviors",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "description",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "display_name",
              "object_id",
              "mail_nickname"
            ]
          },
          {
            "name": "dynamic_membership",
            "type": [
              "list",
              [
                "object",
                {
                  "enabled": "bool",
                  "rule": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "external_senders_allowed",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "hide_from_address_lists",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "hide_from_outlook_clients",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "include_transitive_members",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false
          },
          {
            "name": "mail",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "mail_enabled",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "mail_nickname",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "display_name",
              "object_id",
              "mail_nickname"
            ]
          },
          {
            "name": "members",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "display_name",
              "object_id",
              "mail_nickname"
            ]
          },
          {
            "name": "onpremises_domain_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_group_type",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_netbios_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_sam_account_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_security_identifier",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_sync_enabled",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "owners",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "preferred_language",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "provisioning_options",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "proxy_addresses",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "security_enabled",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false
          },
          {
            "name": "theme",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "types",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "visibility",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "writeback_enabled",
            "type": "bool",
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_group_role_management_policy": {
      "block": {
        "attributes": [
          {
            "name": "description",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "group_id",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "role_id",
            "type": "string",
            "required": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_groups": {
      "block": {
        "attributes": [
          {
            "name": "display_name_prefix",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "display_names",
              "display_name_prefix",
              "object_ids",
              "return_all"
            ]
          },
          {
            "name": "display_names",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "display_names",
              "display_name_prefix",
              "object_ids",
              "return_all"
            ]
          },
          {
            "name": "ignore_missing",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false,
            "conflicts_with": [
              "return_all"
            ]
          },
          {
            "name": "mail_enabled",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "object_ids"
            ]
          },
          {
            "name": "object_ids",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "display_names",
              "display_name_prefix",
              "object_ids",
              "return_all"
            ]
          },
          {
            "name": "return_all",
            "type": "bool",
            "optional": true,
            "force_new": false,
            "conflicts_with": [
              "ignore_missing"
            ],
            "exactly_one_of": [
              "display_names",
              "display_name_prefix",
              "object_ids",
              "return_all"
            ]
          },
          {
            "name": "security_enabled",
            "type": "bool",
            "optional": true,
            "computed": true,
            "force_new": false,
            "conflicts_with": [
              "object_ids"
            ]
          }
        ]
      }
    },
    "azuread_named_location": {
      "block": {
        "attributes": [
          {
            "name": "country",
            "type": [
              "list",
              [
                "object",
                {
                  "countries_and_regions": [
                    "list",
                    "string"
                  ],
                  "country_lookup_method": "string",
                  "include_unknown_countries_and_regions": "bool"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "required": true,
            "force_new": false
          },
          {
            "name": "ip",
            "type": [
              "list",
              [
                "object",
                {
                  "ip_ranges": [
                    "list",
                    "string"
                  ],
                  "trusted": "bool"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_service_principal": {
      "block": {
        "attributes": [
          {
            "name": "account_enabled",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "alternative_names",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "app_role_assignment_required",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "app_role_ids",
            "type": [
              "map",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "app_roles",
            "type": [
              "list",
              [
                "object",
                {
                  "allowed_member_types": [
                    "list",
                    "string"
                  ],
                  "description": "string",
                  "display_name": "string",
                  "enabled": "bool",
                  "id": "string",
                  "value": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "application_tenant_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "client_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "client_id",
              "display_name",
              "object_id"
            ]
          },
          {
            "name": "description",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "client_id",
              "display_name",
              "object_id"
            ]
          },
          {
            "name": "feature_tags",
            "type": [
              "list",
              [
                "object",
                {
                  "custom_single_sign_on": "bool",
                  "enterprise": "bool",
                  "gallery": "bool",
                  "hide": "bool"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "features",
            "type": [
              "list",
              [
                "object",
                {
                  "custom_single_sign_on_app": "bool",
                  "enterprise_application": "bool",
                  "gallery_application": "bool",
                  "visible_to_users": "bool"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "homepage_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "login_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "logout_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "notes",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "notification_email_addresses",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "oauth2_permission_scope_ids",
            "type": [
              "map",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "oauth2_permission_scopes",
            "type": [
              "list",
              [
                "object",
                {
                  "admin_consent_description": "string",
                  "admin_consent_display_name": "string",
                  "enabled": "bool",
                  "id": "string",
                  "type": "string",
                  "user_consent_description": "string",
                  "user_consent_display_name": "string",
                  "value": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "client_id",
              "display_name",
              "object_id"
            ]
          },
          {
            "name": "preferred_single_sign_on_mode",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "redirect_uris",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "saml_metadata_url",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "saml_single_sign_on",
            "type": [
              "list",
              [
                "object",
                {
                  "relay_state": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "service_principal_names",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "sign_in_audience",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "tags",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "type",
            "type": "string",
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_service_principals": {
      "block": {
        "attributes": [
          {
            "name": "client_ids",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "client_ids",
              "display_names",
              "object_ids",
              "return_all"
            ]
          },
          {
            "name": "display_names",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "client_ids",
              "display_names",
              "object_ids",
              "return_all"
            ]
          },
          {
            "name": "ignore_missing",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false,
            "conflicts_with": [
              "return_all"
            ]
          },
          {
            "name": "object_ids",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "client_ids",
              "display_names",
              "object_ids",
              "return_all"
            ]
          },
          {
            "name": "return_all",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false,
            "conflicts_with": [
              "ignore_missing"
            ],
            "exactly_one_of": [
              "client_ids",
              "display_names",
              "object_ids",
              "return_all"
            ]
          },
          {
            "name": "service_principals",
            "type": [
              "list",
              [
                "object",
                {
                  "account_enabled": "bool",
                  "app_role_assignment_required": "bool",
                  "application_tenant_id": "string",
                  "client_id": "string",
                  "display_name": "string",
                  "object_id": "string",
                  "preferred_single_sign_on_mode": "string",
                  "saml_metadata_url": "string",
                  "service_principal_names": [
                    "list",
                    "string"
                  ],
                  "sign_in_audience": "string",
                  "tags": [
                    "list",
                    "string"
                  ],
                  "type": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_user": {
      "block": {
        "attributes": [
          {
            "name": "account_enabled",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "age_group",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "business_phones",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "city",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "company_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "consent_provided_for_minor",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "cost_center",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "country",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "creation_type",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "department",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "display_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "division",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "employee_hire_date",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "employee_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "employee_id",
              "mail",
              "mail_nickname",
              "object_id",
              "user_principal_name"
            ]
          },
          {
            "name": "employee_type",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "external_user_state",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "fax_number",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "given_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "im_addresses",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "job_title",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "mail",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "employee_id",
              "mail",
              "mail_nickname",
              "object_id",
              "user_principal_name"
            ]
          },
          {
            "name": "mail_nickname",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "employee_id",
              "mail",
              "mail_nickname",
              "object_id",
              "user_principal_name"
            ]
          },
          {
            "name": "manager_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "mobile_phone",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "object_id",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "employee_id",
              "mail",
              "mail_nickname",
              "object_id",
              "user_principal_name"
            ]
          },
          {
            "name": "office_location",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_distinguished_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_domain_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_immutable_id",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_sam_account_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_security_identifier",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_sync_enabled",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "onpremises_user_principal_name",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "other_mails",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "postal_code",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "preferred_language",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "proxy_addresses",
            "type": [
              "list",
              "string"
            ],
            "computed": true,
            "force_new": false
          },
          {
            "name": "show_in_address_list",
            "type": "bool",
            "computed": true,
            "force_new": false
          },
          {
            "name": "state",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "street_address",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "surname",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "usage_location",
            "type": "string",
            "computed": true,
            "force_new": false
          },
          {
            "name": "user_principal_name",
            "type": "string",
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "employee_id",
              "mail",
              "mail_nickname",
              "object_id",
              "user_principal_name"
            ]
          },
          {
            "name": "user_type",
            "type": "string",
            "computed": true,
            "force_new": false
          }
        ]
      }
    },
    "azuread_users": {
      "block": {
        "attributes": [
          {
            "name": "employee_ids",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "object_ids",
              "user_principal_names",
              "mail_nicknames",
              "mails",
              "employee_ids",
              "return_all"
            ]
          },
          {
            "name": "ignore_missing",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false,
            "conflicts_with": [
              "return_all"
            ]
          },
          {
            "name": "mail_nicknames",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "object_ids",
              "user_principal_names",
              "mail_nicknames",
              "mails",
              "employee_ids",
              "return_all"
            ]
          },
          {
            "name": "mails",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "object_ids",
              "user_principal_names",
              "mail_nicknames",
              "mails",
              "employee_ids",
              "return_all"
            ]
          },
          {
            "name": "object_ids",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "object_ids",
              "user_principal_names",
              "mail_nicknames",
              "mails",
              "employee_ids",
              "return_all"
            ]
          },
          {
            "name": "return_all",
            "type": "bool",
            "optional": true,
            "default": false,
            "force_new": false,
            "conflicts_with": [
              "ignore_missing"
            ],
            "exactly_one_of": [
              "object_ids",
              "user_principal_names",
              "mail_nicknames",
              "mails",
              "employee_ids",
              "return_all"
            ]
          },
          {
            "name": "user_principal_names",
            "type": [
              "list",
              "string"
            ],
            "optional": true,
            "computed": true,
            "force_new": false,
            "exactly_one_of": [
              "object_ids",
              "user_principal_names",
              "mail_nicknames",
              "mails",
              "employee_ids",
              "return_all"
            ]
          },
          {
            "name": "users",
            "type": [
              "list",
              [
                "object",
                {
                  "account_enabled": "bool",
                  "display_name": "string",
                  "employee_id": "string",
                  "mail": "string",
                  "mail_nickname": "string",
                  "object_id": "string",
                  "onpremises_immutable_id": "string",
                  "onpremises_sam_account_name": "string",
                  "onpremises_user_principal_name": "string",
                  "usage_location": "string",
                  "user_principal_name": "string"
                }
              ]
            ],
            "computed": true,
            "force_new": false
          }
        ]
      }
    }
  }
}`)
	if err := json.Unmarshal(b, &ProviderSchemaInfo); err != nil {
        fmt.Fprintf(os.Stderr, "unmarshalling the provider schema (azuread): %s", err)
		os.Exit(1)
	}
    ProviderSchemaInfo.Version = "3.4.0"
}
