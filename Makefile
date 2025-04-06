include version.mk

LDFLAGS += -X 'github.com/litsea/gin-example/version.Version=$(VERSION)'
LDFLAGS += -X 'github.com/litsea/gin-example/version.GitRev=$(GITREV)'
LDFLAGS += -X 'github.com/litsea/gin-example/version.GitBranch=$(GITBRANCH)'
LDFLAGS += -X 'github.com/litsea/gin-example/version.BuildDate=$(DATE)'

# Checks

# Check for go
.PHONY: check-go
check-go:
	@which go > /dev/null 2>&1 || (echo "Error: go is not installed" && exit 1)

# Check for golangci-lint
.PHONY: check-golangci-lint
check-golangci-lint:
	@which golangci-lint > /dev/null 2>&1 || (echo "Error: golangci-lint is not installed" && exit 1)

# Targets that require the checks
vet: check-go
lint: check-golangci-lint
test: check-go
lint-fix: check-golangci-lint
build: check-go

.PHONY: vet
vet: ## Run vet
	go vet -race ./...

.PHONY: lint
lint: ## Run test
	golangci-lint run ./...

.PHONY: test
test: ## Run test
	go test -race ./...

.PHONY: lint-fix
lint-fix: ## Auto lint fix
	golangci-lint run --fix ./...

.PHONY: ci
ci: vet lint test ## Run CI (vet, lint, test)

.PHONY: build
build: ## Builds the binary locally
	go build -ldflags "all=$(LDFLAGS)" -o app .

## Help display.
## Pulls comments from beside commands and prints a nicely formatted
## display with the commands and their usage information.
.DEFAULT_GOAL := help

.PHONY: help
help: ## Prints this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
