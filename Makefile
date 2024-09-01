include .local.env
include .credentials

# Project directories
LOCAL_BIN=$(CURDIR)/bin
CONF_DIR=$(CURDIR)/config
SOURCE_DIR=$(CURDIR)/cmd

# Go compilation arguments
GOARCH?=amd64
GOOS?=linux
GO_CGO=CGO_ENABLED=0
GO_LDFLAGS=-ldflags="-s -w"

# Tools versions
GOLINT_VER=v1.60.3
GOLINT_CACHE=$(CURDIR)/.golangci-lint-cache
GOOSE_VER=v3.21.1

# Image tag from hash
GIT_SHORT_HASH=$(shell git rev-parse --short HEAD)

SILENT = @

# Install dependences
PHONY: install-deps
install-deps:
	$(SILENT) GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINT_VER)
	$(SILENT) GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@$(GOOSE_VER)

# Download dependences
PHONY: get-deps
download-deps:
	$(SILENT) go mod download
	$(SILENT) go mod verify

PHONY: update-deps
update-deps:
	$(SILENT) go get -u ./...

# Base init
PHONY: init
init:
	$(SILENT) rm -rf $(LOCAL_BIN)
	$(SILENT) rm -rf $(GOLINT_CACHE)
	$(SILENT) mkdir -p $(LOCAL_BIN)
	$(SILENT) mkdir -p $(GOLINT_CACHE)
	make install-deps
	make download-deps

# Local linter run
PHONY: lint
lint:
	$(SILENT) $(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml
PHONY: lint-fix
lint-fix:
	$(SILENT) $(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml --fix

# Fast build
PHONY: build
build:
	$(SILENT) $(GO_CGO) GOARCH=$(GOARCH) GOOS=$(GOOS) go build $(GO_LDFLAGS) -o $(LOCAL_BIN)/$(APP_NAME) $(SOURCE_DIR)

# Prod build
PHONY: build-prod
build-prod:
	make build
	$(SILENT) upx --best $(LOCAL_BIN)/$(APP_NAME)

# Make run
PHONY: run
run:
	$(SILENT) $(GO_CMP_ARGS) go run $(SOURCE_DIR)

