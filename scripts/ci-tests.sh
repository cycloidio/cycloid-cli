#! /usr/bin/env sh

network="$(docker compose ps --format json | jq -rs '.[0].Networks')"
code_path="/go/src/github.com/cycloidio/${REPO_NAME:-cycloid-cli}"
go_version="${GOLANG_VERSION:-$(grep -i '^go ' go.mod | cut -d' ' -f2)}"

docker run -it --rm \
	-e API_LICENCE_KEY \
	--network "$network" \
	-v "${TEST_REPO_PATH:-$(pwd)}:${code_path}" \
	-w "${code_path}" \
	"cycloid/golang:${go_version}" \
	go test -v ./... -p 1 -failfast
