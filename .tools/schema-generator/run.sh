#!/bin/bash

MYDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MYNAME="$(basename "${BASH_SOURCE[0]}")"
ROOTDIR="$MYDIR/../.."

die() {
    echo "$@" >&2
    exit 1
}

usage() {
    cat << EOF
Usage: ./run.sh [options] 

Options:
    -h|--help           show this message

Arguments:
    provider_dir        The path to the provider repo
    provider_version    The version of the provider (e.g. v1.0.0)
EOF
}

main() {
    while :; do
        case $1 in
            -h|--help)
                usage
                exit 1
                ;;
            --)
                shift
                break
                ;;
            *)
                break
                ;;
        esac
        shift
    done

    local expect_n_arg
    expect_n_arg=2
    [[ $# = "$expect_n_arg" ]] || die "wrong arguments (expected: $expect_n_arg, got: $#)"

    provider_dir=$1
    provider_version=$2

    [[ -d $provider_dir ]] || die "no such directory: $provider_dir"

    pushd $provider_dir > /dev/null

    command -v jq > /dev/null || die "jq is not available, please install it"
    provider_name=$(go mod edit -json | jq .Module.Path | tr -d '"' | sed -n 's;^.\+terraform-provider-\(.\+\)$;\1;p')
    case $provider_name in
        google)
            google_pre_hook
            ;;
    esac

    target_location="./internal/tfadd"
    mkdir -p $target_location
    cp -r "$MYDIR/$provider_name/main.go" "$target_location"
    git checkout "$provider_version" || die "failed to checkout provider version $provider_version"

    cat << EOF >> go.mod
require (
    github.com/magodo/tfadd v0.0.0
)

replace github.com/magodo/tfadd => $ROOTDIR
EOF
    go mod tidy || die "failed to run go mod tidy"
    go mod vendor || die "failed to run go mod vendor"

    out=$(go run "$target_location/main.go") || die "failed to generate provider schema"
    cat << EOF > "$ROOTDIR/providers/$provider_name/provider_gen.go"
// Auto-Generated Code; DO NOT EDIT.
package $provider_name

import (
	"encoding/json"
	"fmt"
	"github.com/magodo/tfadd/schema/legacy"
	"os"
)

var ProviderSchemaInfo legacy.ProviderSchema

func init() {
    b := []byte(\`$out\`)
	if err := json.Unmarshal(b, &ProviderSchemaInfo); err != nil {
		fmt.Fprintf(os.Stderr, "unmarshalling the provider schema: %s", err)
		os.Exit(1)
	}
    ProviderSchemaInfo.Version = "${provider_version#v}"
}
EOF

    popd > /dev/null
}

google_pre_hook() {
    # Remove the scripts directory as it will fail `go mod tidy` as one of the imported package is not public
    mv scripts .scripts.del

    sed -i 's;go 1.16;go 1.18;' go.mod
}

main "$@"
