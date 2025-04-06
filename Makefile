include version.mk

LDFLAGS += -X 'github.com/litsea/gin-example/version.Version=$(VERSION)'
LDFLAGS += -X 'github.com/litsea/gin-example/version.GitRev=$(GITREV)'
LDFLAGS += -X 'github.com/litsea/gin-example/version.GitBranch=$(GITBRANCH)'
LDFLAGS += -X 'github.com/litsea/gin-example/version.BuildDate=$(DATE)'

.PHONY: build
build: ## Builds the binary locally
	go build -ldflags "all=$(LDFLAGS)" -o app .
