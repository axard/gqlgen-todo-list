#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

export CGO_ENABLED=0
export GO111MODULE=on
# export GOFLAGS="-mod=vendor"

go run                                                                \
   scripts/gqlgen.go generate                                         \
   --config=build/graphql/graphql.yaml
