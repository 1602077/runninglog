SHELL=/bin/bash
.DEFAULT_GOAL=help
REPO_ROOT := $(shell git rev-parse --show-toplevel)

ifeq (,$(shell go env GOBIN))
	GOBIN=$(BUILD_DIR)/gobin
else
	GOBIN=$(shell go env GOBIN)
endif
export PATH:=$(GOBIN):${PATH}


STRAVA_API_SPEC_LOCATION=api/strava/swagger.json
STRAVA_PKG_LOCATION=pkg/generated/strava

.PHONY: ci
ci: fmt lint test ## runs all stages in CI locally.

.PHONY: fmt
fmt: ## formats go backend codebase
	@[ -x "$(command -v gofumpt)" ] || go install mvdan.cc/gofumpt@latest
	@[ -x "$(command -v goimports-reviser)" ] || go install github.com/incu6us/goimports-reviser/v3@latest
	gofumpt -l -w .
	goimports-reviser ./...

.PHONY: lint
lint: ## lints go backend codebase
	@[ -x "$$(command -v golangci-lint)" ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.56.1
	go vet ./...
	# golangci-lint run

.PHONY: help
help: ## shows this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: generate
# deps: brew install swagger-codegen@2 maven
generate: ## generates strava go client from upstream swagger spec.
	wget https://developers.strava.com/swagger/swagger.json -O $(STRAVA_API_SPEC_LOCATION)
	swagger-codegen generate -i $(STRAVA_API_SPEC_LOCATION) -l go -o $(STRAVA_PKG_LOCATION)

.PHONY: tidy
tidy: ## runs go mod tidy
	go mod tidy

.PHONY: test-unit
test-unit: ## runs go unit tests.
	go test -v -short ./...

.PHONY: test
test: ## runs all go tests.
	go test -v ./...

