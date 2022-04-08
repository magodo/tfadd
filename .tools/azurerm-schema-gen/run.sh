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
    provider_name       The name of the provider (e.g. azurerm)
    provider_dir        The path to the provider repo
    provider_version    The version of the provider (e.g. v1.0.0)
EOF
}

main() {
    declare -A target_locations=(["azurerm"]="internal/tools")

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
    expect_n_arg=3
    [[ $# = "$expect_n_arg" ]] || die "wrong arguments (expected: $expect_n_arg, got: $#)"

    provider_name=$1
    provider_dir=$2
    provider_version=$3

    [[ -d $provider_dir ]] || die "no such directory: $provider_dir"
    [[ -z ${target_locations[$provider_name]} ]] && die "unknown provider name: $provider_name"

    target_location="$provider_dir/${target_locations[$provider_name]}"

    cp -r "$MYDIR/$provider_name/main.go" "$target_location"
    pushd $provider_dir > /dev/null
    git checkout "$provider_version" || die "failed to checkout provider version $provider_version"
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

var ProviderVersion = "${provider_version#v}"

var ProviderSchemaInfo legacy.ProviderSchema

func init() {
    b := []byte(\`$out\`)
	if err := json.Unmarshal(b, &ProviderSchemaInfo); err != nil {
		fmt.Fprintf(os.Stderr, "unmarshalling the provider schema: %s", err)
		os.Exit(1)
	}
}
EOF
    popd > /dev/null
}

main "$@"
