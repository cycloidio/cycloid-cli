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

# docs
DOCFILES := $(shell grep -l '\[TOCBEGIN\]' $(shell find ./ -name '*md') 2>/dev/null)

# GO
GO_ENV             := docker-compose run --rm -u $(shell id -u):$(shell id -g) -e "GOCACHE=/tmp/gocache" -v $(GOCACHE):/tmp/gocache yd-go
# This needed to be this way because of the 'golint' it support or pkg or files, but not both
GO_SPKGS           := $(REPO_PATH)/http/... $(REPO_PATH)/services/...
# Setup the -ldflags build option for go here, interpolate the variable values
GO_LDFLAGS         ?= -ldflags \
	"-X $(REPO_PATH)/internal/version.Version=$(VERSION)\
	 -X $(REPO_PATH)/internal/version.Revision=$(REVISION)\
	 -X $(REPO_PATH)/internal/version.Branch=$(BRANCH)\
	 -X $(REPO_PATH)/internal/version.BuildOrigin=$(BUILD_ORIGIN)\
	 -X $(REPO_PATH)/internal/version.BuildDate=$(BUILD_DATE)"

# generate & fmt
GO_FILES        := $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./gen/*" -not -path "*/mock/*" -not -path "*/mock_*")

# testing
VERBOSE               := 0
COUNT                 := 0
TEST_OPTS             := -timeout 900s
COV_ON                := 1
COV_DEF_FILE          := profile.cov
TEST_PREFIX           := docker-compose run --rm -e "CYCLOID_TEST_BRANCH=${USER}" -v $(GOCACHE):/root/.cache/go-build youdeploy-cli
STOPLIGHT_API_VERSION := C6Dm8wDj2sZJMTj8B

ifeq ($(VERBOSE), 1)
	TEST_OPTS := $(TEST_OPTS) -v
endif

ifeq ($(COUNT), 1)
	TEST_OPTS := $(TEST_OPTS) -count=1
endif

ifeq ($(IS_CI), 1)
	export CYCLOID_TEST_BRANCH = ci_test
	TEST_PREFIX :=
	GO_ENV :=
endif

ifeq ($(COV_FILE),)
COV_FILE := $(COV_DEF_FILE)
endif

ifeq ($(COV_ON), 0)
	COV_TMP := profile.tmp
	COV_OPTS := -short -covermode=count -coverprofile=$(COV_TMP)
	COV_AGGREGATE := && (cat $(COV_TMP) | tail -n +2 >> $(COV_DEF_FILE) && rm -f $(COV_TMP))
endif

TEST_CMD := $(TEST_PREFIX) go test $(TEST_OPTS)
#
# Get the list of the go packages excluding gen, http and services using the local
# go installed otherwise the docker image
PKG_LIST := $(shell $(GO_ENV) go list ./... | grep -v "$(REPO_PATH)/\(gen\|http\|services\)")

.PHONY: help
help: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/:.*##/:##/' | column -t -s '##'

.PHONY: markdown-docs
markdown-docs:                           ## Create TOC for markdown files
	@for doc in $(DOCFILES); do \
	   (sed -i '/\[TOCBEGIN\]/,/\[TOCEND\]/{//!d}' $${doc} && \
           ./scripts/gh-md-toc $${doc} > TOCFILE && \
	   sed -i "/\[TOCBEGIN\]/ r TOCFILE" $${doc} && \
	   rm TOCFILE ) && \
	   echo "TOC generation for $${doc} done!" || \
	   (echo "TOC generation for $${doc} failed!" && exit 1); \
	 done

.PHONY: full-lint
full-lint: ## Run a wider range of linter against the sources of the project
	@$(GO_ENV) golangci-lint run --exclude-use-default=false --disable-all \
		--enable=golint --enable=govet --enable=goimports --enable=ineffassign \
		--enable=structcheck --enable=unused --enable=misspell \
		--enable=unconvert --enable=interfacer --deadline=5m

.PHONY: fmt
fmt: ## Verify the sources of the project are goimports compliant
	@if [ "$(shell $(GO_ENV) goimports -l $(GO_FILES) | wc -l)" != "0" ]; then \
		echo "--- CHECK FAIL: Some files did not pass goimports $(shell $(GO_ENV) goimports -l $(GO_FILES))"; exit 2; \
	fi

.PHONY: lint
lint: ## Verify the sources of the project are goimports compliant
	@$(GO_ENV) golint -set_exit_status \
		$(shell \
			$(shell which go > /dev/null; if [ "$$?" != 0 ]; then echo '$(GO_ENV) go'; else echo 'go'; fi) \
			list ./... | grep -v -e "$(REPO_PATH)/gen" -e "$(REPO_PATH)/services/email/internal/template"\
		)
	@# Make sure that we don't call the methods present in "gen/restapi/configure_you_deploy.go"
	@egrep '\.ConfigureFlags\(\)|\.ConfigureAPI\(\)' -rq cmd http\
				&& if [ "$$?" -eq 0 ]; then \
					echo "Don't use the generated Server ConfigureFlags && ConfigureAPI functions; read the FAQ.\n" && \
					exit 1; \
				fi \
				|| true

.PHONY: build
build: ## Builds the binary
	GO111MODULE=on CGO_ENABLED=1 GOARCH=amd64 go build -o $(BINARY) $(GO_LDFLAGS) $(REPO_PATH)
	#GO111MODULE=on CGO_ENABLED=1 GOARCH=amd64 go build -trimpath -o $(BINARY) $(GO_LDFLAGS) $(REPO_PATH)
	#GO111MODULE=on CGO_ENABLED=0 GOARCH=amd64 go build -o $(BINARY) $(GO_LDFLAGS) $(REPO_PATH)

.PHONY: tidy
tidy: ## Clean unused dependencies
	@$(GO_ENV) go mod tidy

.PHONY: vendor
vendor: ## Fetch and vendor all packages specified under go.mod
	@$(GO_ENV) go mod vendor

.PHONY: ci
ci: check-migrations-manifest lint fmt test translations-diff  ## Check the migrations Manifest and run goimports, go vet & run the tests

.PHONY: go-docs
go-docs: ## Serve GoDoc Project docs (http://localhost:8888/pkg/github.com/cycloidio/youdeploy-http-api/)
	@echo "GoDoc is accessible via: http://localhost:8888/pkg/$(REPO_PATH)/"
	@docker-compose run --rm -p 8888:8888 youdeploy-api godoc -http :8888

.PHONY: test-package
test-package: ## Run tests against the given package(s) set in P variable
	@$(TEST_CMD) $(P)

# Run 2 sets of tests, parallel ones go_packages_list and sync ones GO_SPKGS because the GO_PPKGS depend on CC and overlap each other
.PHONY: test
test: dev-env-up ## Run all the tests (COV_ON=0 to create coverage file)
ifeq ($(COV_ON), 0)
	@echo 'mode: count' > $(COV_DEF_FILE)
endif
	@$(TEST_CMD) $(PKG_LIST)
	@err=0 ; for p in $(GO_SPKGS); do \
	    /bin/sh -c "$(TEST_CMD) $(COV_OPTS) -p=1 $$p" $(COV_AGGREGATE) || err=1;\
	done; \
	if [ $$err != "0" ]; then exit 1; fi;
	@rm -f $(COV_TMP)

.PHONY: cov
cov: ## Display code coverage COV_FILE (default profile.cov)
	@grep -v -E '(/mysql/|/endpoints/|logger|trace|bindata)' $(COV_DEF_FILE) > final-profile.cov
	@echo "Available coverage files: $(shell ls -1 *.cov)"
	@echo "Displaying coverage for: $(COV_FILE)"
	@$(TEST_PREFIX) go tool cover -func $(COV_FILE)


.PHONY: show-prs-not-in-master
show-prs-not-in-master: ## Show the list of PRs merged into develop but not into master
	@git fetch --all
	@for i in $$(git rev-list --oneline --first-parent  origin/master..origin/develop | awk '/Merge pull/{print $$5}' | tr -d '#'); do echo "https://$(REPO_PATH)/pull/$$i"; done
