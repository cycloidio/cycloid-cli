ifeq ($(GOCACHE),)
	GOCACHE := $(HOME)/.cache/go-build
endif

ifneq (, $(shell which go))
	GOCACHE := $(shell go env GOCACHE)
endif

SHELL      := /bin/sh

REPO_PATH    := github.com/cycloidio/youdeploy-cli
BINARY       ?= cy
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
