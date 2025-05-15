ifeq ($(GOCACHE),)
	GOCACHE := $(HOME)/.cache/go-build
endif

ifneq (, $(shell which go))
	GOCACHE := $(shell go env GOCACHE)
endif

SHELL      := /bin/sh

REPO_PATH  := github.com/cycloidio/cycloid-cli

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

# E2E tests
CY_API_URL         ?= "https://api-cli-test.staging.cycloid.io/"
CY_TEST_ROOT_ORG ?= "cycloid"
# You can get the key in the admin_api_key cred in the cli console
CY_TEST_API_KEY       ?=

# Local E2E tests
# Note! Requires access to the private cycloid BE, only acessible within the organisation
# AWS - ECR login
export AWS_ACCESS_KEY_ID 	  ?= $(shell vault read -field=access_key secret/cycloid/aws)
export AWS_SECRET_ACCESS_KEY ?= $(shell vault read -field=secret_key secret/cycloid/aws)
export AWS_DEFAULT_REGION    ?= eu-west-1
export AWS_ACCOUNT_ID        ?= $(shell vault read -field=account_id secret/cycloid/aws)
# Local BE
LOCAL_BE_GIT_PATH ?= ../youdeploy-http-api
YD_API_TAG        ?= staging
API_LICENCE_KEY   ?=

.PHONY: help
help: ## Show this help
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/:.*##/:##/' | column -t -s '##'

.PHONY: build
build: ## Builds the binary
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY) $(GO_LDFLAGS) $(REPO_PATH)
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY)-linux-amd64 $(GO_LDFLAGS) $(REPO_PATH)
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BINARY)-windows-amd64 $(GO_LDFLAGS) $(REPO_PATH)
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(BINARY)-darwin-arm64 $(GO_LDFLAGS) $(REPO_PATH)
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(BINARY)-darwin-amd64 $(GO_LDFLAGS) $(REPO_PATH)

.PHONY: test
test: ## Run end to end tests
	CY_API_URL=$(CY_API_URL) \
	  CY_TEST_ROOT_ORG=$(CY_TEST_ROOT_ORG) \
	  CY_TEST_API_KEY=$(CY_TEST_API_KEY) \
		go test ./... -v

.PHONY: delete-old-client
reset-old-client: ## Resets old client folder
	rm -rf ./client && mkdir -p client

# Used in CI, do not use docker compose here.
.PHONY: generate-client
generate-client: reset-old-client ## Generate client from file at SWAGGER_FILE path
	echo "Creating swagger files"; \
	$(SWAGGER_GENERATE)
	@export SWAGGER_VERSION=$$(python -c 'import yaml, sys; y = yaml.safe_load(sys.stdin); print(y["info"]["version"])' < swagger.yml); \
	if [ -z "$$SWAGGER_VERSION" ]; then echo "Unable to read version from swagger"; exit 1; fi; \
	echo $$SWAGGER_VERSION > client/version; \
	go mod tidy

.PHONY: generate-client-from-docs
generate-client-from-docs: reset-old-client ## Generates client using docker and swagger from docs (version -> latest-api)
	@wget -O swagger.yml https://docs.cycloid.io/api/swagger.yml
	@export SWAGGER_VERSION=$$(python -c 'import yaml, sys; y = yaml.safe_load(sys.stdin); print(y["info"]["version"])' < swagger.yml); \
	if [ -z "$$SWAGGER_VERSION" ]; then echo "Unable to read version from swagger"; exit 1; fi; \
	echo $$SWAGGER_VERSION > client/version; \
	make generate-client && \
	echo "Please run the following git commands:"; \
	echo "git add client" && \
	echo "git commit -m 'Bump swagger client to version $$SWAGGER_VERSION'"

.PHONY: docker-login
docker-login: ## Login to ecr, requires aws cli installed
	aws ecr get-login-password --region $(AWS_DEFAULT_REGION) | docker login --username AWS --password-stdin $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_DEFAULT_REGION).amazonaws.com/youdeploy-http-api

.PHONY: start-local-be
start-local-be: ## Starts local BE instance. Note! Only for cycloid developers
	@if [ ! -d ${LOCAL_BE_GIT_PATH} ]; then echo "Unable to find BE at LOCAL_BE_GIT_PATH"; exit 1; fi;
	@if [ -z "$$API_LICENCE_KEY" ]; then echo "API_LICENCE_KEY is not set"; exit 1; fi; \
	echo "Starting Local BE..."
	@echo "Generating fake data to be used in the tests..."
	@cd $(LOCAL_BE_GIT_PATH) && sed -i '/cost-explorer-es/d' config.yml
	@cd $(LOCAL_BE_GIT_PATH) && YD_API_TAG=${YD_API_TAG} API_LICENCE_KEY=${API_LICENCE_KEY} \
	$(DOCKER_COMPOSE) -f docker-compose.yml -f docker-compose.cli.yml up youdeploy-init
	@echo "Running BE server with the fake data generated..."
	@cd $(LOCAL_BE_GIT_PATH) && YD_API_TAG=${YD_API_TAG} API_LICENCE_KEY=${API_LICENCE_KEY} \
	$(DOCKER_COMPOSE) -f docker-compose.yml -f docker-compose.cli.yml up -d youdeploy-api

.PHONY: local-e2e-test
local-e2e-test: ## Launches local e2e tests. Note! Only for cycloid developers
	@if [ -z "$(shell curl -I --connect-timeout 2 "172.42.0.3:3001" 2>&1 | grep -w "500")" ]; then make start-local-be; fi;
	@echo "Local BE is up!"
	@echo "Running Local e2e tests!"
	@make test CY_TEST_ROOT_API_KEY=$(shell cat ${LOCAL_BE_GIT_PATH}/API_KEY)

.PHONY: delete-local-be
delete-local-be: ## Creates local BE instance and starts e2e tests. Note! Only for cycloid developers
	@if [ ! -d ${LOCAL_BE_GIT_PATH} ]; then echo "Unable to find BE at LOCAL_BE_GIT_PATH"; exit 1; fi;
	@echo "Deleting local BE instances !"
	@cd $(LOCAL_BE_GIT_PATH) && $(DOCKER_COMPOSE) down -v --remove-orphans

.PHONY: new-changelog-entry
new-changelog-entry: ## Create a new entry for unreleased element
	@echo ${PATH}
	docker run -it -v $(CURDIR):/cycloid-cli -w /cycloid-cli cycloid/cycloid-toolkit changie new

.PHONY: lint
lint: ## Lint the source code
	@echo -e "Running golangci-lint"
	@golangci-lint run -v

.PHONY: format-go
format-go:
	@gci write --skip-generated -s standard -s default -s "prefix(github.com/cycloidio)" . > /dev/null
	@goimports -w -local github.com/cycloidio .
