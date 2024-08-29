package schema

import tfpluginschema "github.com/magodo/tfpluginschema/schema"

type ProviderSchema struct {
	// The provder version. This is defined in the vcs.
	Version         string
	ResourceSchemas map[string]*Schema `json:"resource_schemas,omitempty"`
}

type Schema struct {
	Block *tfpluginschema.SchemaBlock `json:"block,omitempty"`
}
