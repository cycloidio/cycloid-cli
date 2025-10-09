ifeq ($(GOCACHE),)
	GOCACHE := $(HOME)/.cache/go-build
endif

ifneq (, $(shell which go))
	GOCACHE := $(shell go env GOCACHE)
endif

SHELL      := /bin/sh
REPO_NAME   ?= cycloid-cli
MAKEFILE_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

# IMAGE BUILD
BINARY       ?= cy
# VERSION example v1.0.47
VERSION      ?= $(shell cat client/version)
REVISION     ?= $(shell git rev-parse --short HEAD 2> /dev/null  || echo 'unknown')
BRANCH       ?= $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown')
BUILD_ORIGIN ?= $(USER)@$(shell hostname -f)
BUILD_DATE   ?= $(shell date --utc -Iseconds)
DOCKER_COMPOSE ?= $(shell docker compose --help >/dev/null 2>&1 && echo "docker compose" || echo "docker-compose")

# GO
# Setup the -ldflags build option for go here, interpolate the variable values
GO_LDFLAGS ?= -ldflags \
	"-X $(REPO_PATH)/internal/version.Version=$(VERSION)\
	 -X $(REPO_PATH)/internal/version.Revision=$(REVISION)\
	 -X $(REPO_PATH)/internal/version.Branch=$(BRANCH)\
	 -X $(REPO_PATH)/internal/version.BuildOrigin=$(BUILD_ORIGIN)\
	 -X $(REPO_PATH)/internal/version.BuildDate=$(BUILD_DATE)"

# SWAGGER
SWAGGER_FILE ?= "swagger.yml"
SWAGGER_GENERATE = swagger generate client \
		--spec=$(SWAGGER_FILE) \
		--default-produces="application/vnd.cycloid.io.v1+json" \
		--target=./client \
		--name=api

BACKEND_TAG ?= staging

-include .env
-include .api_key

.PHONY: help
help: ## Show this help
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/:.*##/:##/' | column -t -s '##'

.env: ## generate the .env files with the required secrets for this repo
	@rm .env || true
	@CY_API_KEY=$${CY_SAAS_API_KEY?A valid API key to the cycloid org in our saas is required for this target. It must be provided via the CY_SAAS_API_KEY} \
		cy uri interpolate .env.sample > .env

.PHONY: build
build: ## Build all the binaries
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY) $(GO_LDFLAGS) $(REPO_PATH)
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY)-linux-amd64 $(GO_LDFLAGS) $(REPO_PATH)
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(BINARY)-linux-arm64 $(GO_LDFLAGS) $(REPO_PATH)
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BINARY)-windows-amd64 $(GO_LDFLAGS) $(REPO_PATH)
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(BINARY)-darwin-arm64 $(GO_LDFLAGS) $(REPO_PATH)
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(BINARY)-darwin-amd64 $(GO_LDFLAGS) $(REPO_PATH)

.PHONY: clean
clean: test-clean
	@rm -f .env
	go clean

.PHONY: test test-clean
test: .env ## Run end to end tests
		go test ./...

test-clean: .env ## Clean test cache
		go clean -testcache

.PHONY: client-delete client-generate client-generate-from-docs
client-delete: ## Resets old client folder
	rm -rf ./client && mkdir -p client

client-generate: client-delete ## Generate client from file at SWAGGER_FILE path
	# Used in CI, do not use docker compose here.
	echo "Creating swagger files"; \
	$(SWAGGER_GENERATE)
	@export SWAGGER_VERSION=$$(python -c 'import yaml, sys; y = yaml.safe_load(sys.stdin); print(y["info"]["version"])' < swagger.yml); \
	if [ -z "$$SWAGGER_VERSION" ]; then echo "Unable to read version from swagger"; exit 1; fi; \
	echo $$SWAGGER_VERSION > client/version; \
	go mod tidy

client-generate-from-docs: client-delete ## Generates client using docker and swagger from docs (version -> latest-api)
	@wget -O swagger.yml https://docs.cycloid.io/api/swagger.yml
	@export SWAGGER_VERSION=$$(python -c 'import yaml, sys; y = yaml.safe_load(sys.stdin); print(y["info"]["version"])' < swagger.yml); \
	if [ -z "$$SWAGGER_VERSION" ]; then echo "Unable to read version from swagger"; exit 1; fi; \
	echo $$SWAGGER_VERSION > client/version; \
	make client-generate && \
	echo "Please run the following git commands:"; \
	echo "git add client" && \
	echo "git commit -m 'Bump swagger client to version $$SWAGGER_VERSION'"

.PHONY: docker-login
docker-login: .env
	echo "$(SECRET_ACCESS)" | docker login rg.fr-par.scw.cloud/cycloidio/cycloid-backend -u $(ACCESS_KEY) --password-stdin

.PHONY: be-start be-stop be-reset
be-start: ## start the local backend
	$(DOCKER_COMPOSE) up -dV

be-stop: ## stop the local backend
	$(DOCKER_COMPOSE) down -v

be-reset: be-stop be-start ## reset the backend

be-db-connect: ## Connect to the local mysql
	$(DOCKER_COMPOSE) exec -it database mysql -uroot -pyoudeploy youdeploy

.PHONY: new-changelog-entry
new-changelog-entry: ## Create a new entry for unreleased element
	@docker run --rm -it -v $(CURDIR):/cycloid-cli -w /cycloid-cli cycloid/cycloid-toolkit:latest changie new

.PHONY: lint lint-go lint-sh
lint: lint-go lint-sh

lint-sh:
	@find . -type f -name '*.sh' -exec shellcheck -f gcc {} \;

lint-go: ## Lint the source code
	@golangci-lint run -v

.PHONY: format format-go format-sh
format: format-go format-sh

format-go: ## format the repo
	@gci write --skip-generated -s standard -s default -s "prefix(github.com/cycloidio)" . > /dev/null
	@goimports -w -local github.com/cycloidio .

format-sh:
	@shfmt -w .

.PHONY: ci-test
.ONEFILE:
ci-test:
	$(MAKEFILE_DIR)/scripts/ci-tests.sh
