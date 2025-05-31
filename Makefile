.PHONY: build clean install test lint help

# Binary name
BINARY=agenter
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "0.1.0")
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Default target
help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the binary
	go build ${LDFLAGS} -o ${BINARY} .

clean: ## Remove build artifacts
	rm -f ${BINARY}
	rm -rf dist/

install: build ## Install binary to /usr/local/bin
	sudo cp ${BINARY} /usr/local/bin/

uninstall: ## Remove binary from /usr/local/bin
	sudo rm -f /usr/local/bin/${BINARY}

test: ## Run tests
	go test -v ./...

lint: ## Run linter
	go vet ./...

run: build ## Build and run
	./${BINARY}

# Development helpers
dev-check: build ## Run check command
	./${BINARY} check

dev-init: build ## Run init command
	./${BINARY} init