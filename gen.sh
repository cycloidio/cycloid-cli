#!/bin/bash

if [[ -z $1 ]]; then
    echo "uasge: version"
    exit 1
fi

echo "# Generate version $1"

rm "plugins/v${1}.so"

echo "- Generate CLI from git tags"
rm -rf cmd/*
for cli in $(ls -1 CLIs/ | sort); do

    rsync -av CLIs/$cli/ cmd/

    if [[ "v$1" == "$cli" ]]; then
        echo "Done $cli"
        break
    fi
done


if [[ $2 == "swagger" ]]; then
echo "- Generate swagger client"
rm -rf ./client
mkdir ./client
docker-compose run swagger generate client \
--spec=swagger-files/api-v10${1}.yml \
--default-produces="application/vnd.cycloid.io.v1+json" \
--target=./client \
--name=api \
--tags=Cycloid \
--tags="Organization External Backends" \
--tags="Organization Credentials" \
--tags="Organization projects" \
--tags="Service catalogs" \
--tags="Organization workers" \
--tags="Organization pipelines" \
--tags="Organization pipelines jobs" \
--tags="Organization pipelines jobs build" \
--tags="Organization Config Repositories" \
--tags="Organization Service Catalog Sources" \
--tags="Organizations" \
--tags="User" \
--tags="Organization members" \
--tags="Organizations"
fi

echo "- Generate plugin"
go build -buildmode=plugin -o "plugins/v${1}.so" cmd/cycloid.go
