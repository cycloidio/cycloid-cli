#! /usr/bin/env sh

network="$(docker compose ps --format json | jq -rs '.[0].Networks')"
code_path="/go/src/github.com/cycloidio/${REPO_NAME:-cycloid-cli}"
go_version="${GOLANG_VERSION:-$(grep -i '^go ' go.mod | cut -d' ' -f2)}"

api_id="$(docker compose ps -q youdeploy-api)"
printf 'Waiting for backend'
while [ "$(docker inspect --format '{{.State.Health.Status}}' "$api_id" 2>/dev/null)" != "healthy" ]; do
	printf '.'
	sleep 5
done
printf '\n'

docker run -it --rm \
	-e API_LICENCE_KEY \
	-e CY_TEST_PROVISION_API \
	-e "CY_TEST_API_URL=${CY_TEST_API_URL:-http://youdeploy-api:3001}" \
	--network "$network" \
	-v "${TEST_REPO_PATH:-$(pwd)}:${code_path}" \
	-w "${code_path}" \
	"golang:${go_version}" \
	go test -v ./... -p 1 -failfast
