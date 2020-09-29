ifeq ($(GOCACHE),)
	GOCACHE := $(HOME)/.cache/go-build
endif

ifneq (, $(shell which go))
	GOCACHE := $(shell go env GOCACHE)
endif

SHELL      := /bin/sh

REPO_PATH    := github.com/cycloidio/youdeploy-cli
BINARY       ?= cy
# VERSION example v1.0.47
VERSION      ?= $(shell git describe --tags --always)
REVISION     ?= $(shell git rev-parse --short HEAD 2> /dev/null  || echo 'unknown')
BRANCH       ?= $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown')
BUILD_ORIGIN ?= $(USER)@$(shell hostname -f)
BUILD_DATE   ?= $(shell date --utc -Iseconds)

# GO
# Setup the -ldflags build option for go here, interpolate the variable values
GO_LDFLAGS         ?= -ldflags \
	"-X $(REPO_PATH)/internal/version.Version=$(VERSION)\
	 -X $(REPO_PATH)/internal/version.Revision=$(REVISION)\
	 -X $(REPO_PATH)/internal/version.Branch=$(BRANCH)\
	 -X $(REPO_PATH)/internal/version.BuildOrigin=$(BUILD_ORIGIN)\
	 -X $(REPO_PATH)/internal/version.BuildDate=$(BUILD_DATE)"

SWAGGER_FILE ?= "gen-swagger/swagger.yml"
SWAGGER_GENERATE = rm -rf ./client; \
	mkdir ./client; \
	docker-compose run swagger generate client \
		--spec=$(SWAGGER_FILE) \
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

.PHONY: help
help: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/:.*##/:##/' | column -t -s '##'

.PHONY: build-plugin
build-plugin: ## Build a plugin from local code
	go build -buildmode=plugin -o "plugins/$(VERSION).so" cmd/cycloid.go

.PHONY: build-plugins-package
build-plugins-package: ## Build pkged.go package file from local plugins
	go get github.com/markbates/pkger/cmd/pkger
	pkger -include /plugins

.PHONY: build
build: ## Builds the binary
	GO111MODULE=on CGO_ENABLED=1 GOARCH=amd64 go build -o $(BINARY) $(GO_LDFLAGS) $(REPO_PATH)
	@#GO111MODULE=on CGO_ENABLED=1 GOARCH=amd64 go build -trimpath -o $(BINARY) $(GO_LDFLAGS) $(REPO_PATH)
	@#GO111MODULE=on CGO_ENABLED=0 GOARCH=amd64 go build -o $(BINARY) $(GO_LDFLAGS) $(REPO_PATH)

.PHONY: build-full
build-full: build-plugin build-plugins-package build ## Builds the binary, the plugin and the package

.PHONY: generate-local-client
generate-local-client: ## Generate client from local swagger file SWAGGER_FILE path
	$(SWAGGER_GENERATE)

.PHONY: generate-client
generate-client: ## Generate client from latest swagger file
	@mkdir -p ./gen-swagger
	@wget -O ./gen-swagger/swagger.yml https://docs.cycloid.io/api/swagger.yml
	@export SWAGGER_VERSION=$$(python -c 'import yaml, sys; y = yaml.safe_load(sys.stdin); print(y["info"]["version"])' < ./gen-swagger/swagger.yml); \
	if [ -z "$$SWAGGER_VERSION" ]; then echo "Unable to read version from swagger"; exit 1; fi; \
	export IS_GIT_TAG_EXIST=$$(git --no-pager tag -l $$SWAGGER_VERSION); \
	if [ -n "$$IS_GIT_TAG_EXIST" ]; then echo "Version tag $$SWAGGER_VERSION already exist in git"; exit 0; fi; \
	echo "Creating swagger files"; \
	$(SWAGGER_GENERATE) && \
	echo "Please run the following git commands:"; \
	echo "git add client" && \
	echo "git commit -m 'Bump swagger client to version $$SWAGGER_VERSION'" && \
	echo "git tag $$SWAGGER_VERSION"
	@rm -rf ./gen-swagger
