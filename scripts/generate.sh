#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

export CGO_ENABLED=0
export GO111MODULE=on

go run github.com/99designs/gqlgen generate                                                        \
   --config=build/graphql/graphql.yaml
