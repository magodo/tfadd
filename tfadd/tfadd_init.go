package tfadd

import (
	"bytes"
	"fmt"
	"text/template"
)

func Init(providers []string) ([]byte, error) {
	if len(providers) == 0 {
		return nil, nil
	}

	type info struct {
		Name    string
		Source  string
		Version string
	}

	var infos []info
	for _, p := range providers {
		pinfo, ok := sdkProviderSchemas["registry.terraform.io/hashicorp/"+p]
		if !ok {
			return nil, fmt.Errorf("Unsupported provider %q\n", p)
		}
		infos = append(infos, info{
			Name:    p,
			Source:  "hashicorp/" + p,
			Version: pinfo.Version,
		})
	}

	out := bytes.Buffer{}
	if err := template.Must(template.New("setup").Parse(`terraform {
  required_providers {
  {{- range . }}
    {{.Name}} = {
      source = "hashicorp/{{.Name}}"
      version = "{{.Version}}"
    }
  {{- end }}
  }
}
`)).Execute(&out, infos); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
