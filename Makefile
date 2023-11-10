.DEFAULT_GOAL := help
PROJECT_BIN = $(shell pwd)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)
GOOS = linux
GOARCH = amd64
CGO_ENABLED = 1
LDFLAGS = "-w -s"
APP := $(notdir $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))

.PHONY: build run .install-air air .install-linter lint cert .install-formatter fmt help compose

release: ## Build release
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags=$(LDFLAGS) -o $(PROJECT_BIN)/$(APP) cmd/main/main.go

debug: ## Build debug and run
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build --gcflags="all=-N -l" -v -o $(PROJECT_BIN)/$(APP)_debug cmd/main/main.go
	$(PROJECT_BIN)/$(APP)_debug

.install-air: ## Install air
	[ -f $(PROJECT_BIN)/air ] || go install github.com/cosmtrek/air@latest && cp $(GOPATH)/bin/air $(PROJECT_BIN)

dev: .install-air ## Run dev server
	$(PROJECT_BIN)/air

.install-linter: ## Install linter
	[ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.54.2

lint: .install-linter ## Run linter
	$(PROJECT_BIN)/golangci-lint run ./...

cert: ## Make ssl cert's
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt

PATH_TO_PRETTIER_PLUGIN = "/opt/helix/node/lib/node_modules/prettier-plugin-go-template/lib/index.js"
PATH_TO_TEMPLATES = "assets/templates"
.install-formatter: ## Install prettier for formats html go templates. See https://github.com/NiklasPor/prettier-plugin-go-template
	npm -g list | grep -e "prettier" -e "prettier-plugin-go-template" || npm -g install prettier prettier-plugin-go-template

compose: ## Up docker-compose.override.yaml
	docker-compose up

fmt: .install-formatter ## [DEPRECATED] Format html go templates
	prettier --plugin $(PATH_TO_PRETTIER_PLUGIN) --parser go-template -w $(PATH_TO_TEMPLATES)

help:
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
