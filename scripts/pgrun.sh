#!/bin/sh

set -o errexit
set -o pipefail

volume="pgdata"

if [ -n "$1" ]; then
    volume=$1
    echo "Using custom docker volume name: $volume"
else
    echo "Using default docker volume name: $volume"
fi

docker run                                                                     \
       --rm                                                                    \
       --name pg-docker                                                        \
       -e POSTGRES_PASSWORD=docker                                             \
       -d                                                                      \
       -p 5432:5432                                                            \
       -v $volume:/var/lib/postgresql/data  postgres
