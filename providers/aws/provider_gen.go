// Auto-Generated Code; DO NOT EDIT.
package aws

import (
	"encoding/json"
	"fmt"
	"github.com/magodo/tfadd/schema/legacy"
	"os"
)

var ProviderSchemaInfo legacy.ProviderSchema

func init() {
	if err := json.Unmarshal(b, &ProviderSchemaInfo); err != nil {
		fmt.Fprintf(os.Stderr, "unmarshalling the provider schema: %s", err)
		os.Exit(1)
	}
    ProviderSchemaInfo.Version = "4.9.0"
}