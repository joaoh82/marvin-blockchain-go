SHELL:=/bin/bash

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

VERSION=dev
COMMIT=$(shell git rev-parse HEAD)
GITDIRTY=$(shell git diff --quiet || echo 'dirty')

GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

APP_NAME := $(notdir $(CURDIR))
SERVICE_NAME := "marvin-blockchain"
IMAGE := "marvin-blockchain"

TARGET_MAX_CHAR_NUM=25
## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

all: help

.PHONY: install-dependencies
## Install dependencies for the service
install-dependencies: ## Install dependencies for the service
	go mod tidy

.PHONY: build
## Build the binary for the marvin blockchain
build:
	CGO_ENABLED=0 go build -o ./bin/marvin ./*.go

.PHONY: build-cli
## Build the binary for marvinclt
build-cli:
	CGO_ENABLED=0 go build -o ./bin/marvinctl ./cmd/marvinctl/*.go

.PHONY: run-tests
## Run the tests for the service
run-tests:
	go test -v ./...
