# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory
BIN := .tmp/bin
COPYRIGHT_YEARS := 2021-2023
LICENSE_IGNORE := --ignore /testdata/
# Set to use a different compiler. For example, `GO=go1.18rc1 make test`.
GO ?= go

.PHONY: help
help: ## Describe useful make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %s\n", $$1, $$2}'

.PHONY: all
all: run

.PHONY: clean
clean: ## Delete intermediate build artifacts
	@# -X only removes untracked files, -d recurses into directories, -f actually removes files/dirs
	git clean -Xdf

.PHONY: generate
generate: $(BIN)/buf $(BIN)/protoc-gen-go $(BIN)/protoc-gen-go-grpc $(BIN)/license-header ## Regenerate code and licenses
	rm -rf internal/gen
	PATH=$(abspath $(BIN)) $(BIN)/buf generate
	$(BIN)/license-header \
		--license-type apache \
		--copyright-holder "Buf Technologies, Inc." \
		--year-range "$(COPYRIGHT_YEARS)" $(LICENSE_IGNORE)

.PHONY: run
run: generate ## Run the compose setup (default)
	@# Disable buildkit, Podman doesn't support it and we don't need it.
	DOCKER_BUILDKIT=0 docker-compose up --build

$(BIN)/buf: Makefile
	@mkdir -p $(@D)
	GOBIN=$(abspath $(@D)) $(GO) install github.com/bufbuild/buf/cmd/buf@v1.17.0

$(BIN)/protoc-gen-go: Makefile go.mod
	@mkdir -p $(@D)
	GOBIN=$(abspath $(@D)) $(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0

.PHONY: $(BIN)/protoc-gen-go-grpc
$(BIN)/protoc-gen-go-grpc:
	@mkdir -p $(@D)
	GOBIN=$(abspath $(@D)) $(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

$(BIN)/license-header: Makefile
	@mkdir -p $(@D)
	GOBIN=$(abspath $(@D)) $(GO) install \
		  github.com/bufbuild/buf/private/pkg/licenseheader/cmd/license-header@90fa81df0e9ef86ed505956be1ab1e8ccd49aa52
