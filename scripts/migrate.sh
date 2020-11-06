#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

docker run                                                                     \
       --rm                                                                    \
       -u $(id -u):$(id -g)                                                    \
       -v $(pwd):/migrations                                                   \
       -w /migrations                                                          \
       --network host                                                          \
       migrate/migrate "$@"
