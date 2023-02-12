GOCMD=go
BINARY_NAME=openfort-api
CONFIG=./conf/example.config.json

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all build
all: help

## Build:
build: ## Build the application binary into out/bin/
	mkdir -p out/bin
	GO111MODULE=on $(GOCMD) build -o out/bin/$(BINARY_NAME) ./cmd/openfort-api

build-linux: ## Build the application binary for Linux systems into out/bin/
	mkdir -p out/bin
	CGO_ENABLED=0 GOOS=linux $(GOCMD) build -o out/bin/$(BINARY_NAME) ./cmd/openfort-api

run: ## Run the application
	CONFIG=$(CONFIG) $(GOCMD) run ./out/bin/$(BINARY_NAME)

run-example: ## Run the example
	CONFIG=$(CONFIG) $(GOCMD) run ./api-example.go

test: ## Run the application tests
	$(GOCMD) test -v ./api/handlers

clean: ## Remove build related file
	rm -fr ./out

vendor: ## Copy all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor

## Docker:
start: ## Start the application in a Docker container
	docker-compose up -d

stop: ## Stop the application Docker container
	docker-compose down

db-remove: ## Erase the application database data
	rm -rf postgres-data

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)

