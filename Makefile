# Go Options
MODULE   := github.com/mbow/frequentlyUsedWords
LDFLAGS  := -w -s
BINDIR   := $(CURDIR)/bin
GOBIN    := $(BINDIR)
PATH     := $(GOBIN):$(PATH)
NAME     := frequentlyUsedWords

# Tools as dependencies
TOOLS += github.com/golangci/golangci-lint/cmd/golangci-lint

# Verbose output
ifdef VERBOSE
V = -v
else
.SILENT:
endif

# Default target
.DEFAULT_GOAL := all

# Make All targets
.PHONY: all
all: tools lint test build

# Download dependencies to go module cache
.PHONY: deps
deps:
	$(info Installing dependencies)
	GO111MODULE=on XDG_CONFIG_HOME=$(CURDIR)/configs go mod download

# Install tooling
.PHONY: tools
tools: deps $(TOOLS)

# Check tools
.PHONY: $(TOOLS)
$(TOOLS): %:
	GOBIN=$(GOBIN) go install $(V) $*

# Lint code
.PHONY: lint
lint: tools
	$(info Running Go code checkers and linters)
	$(GOBIN)/golangci-lint $(V) run --enable-all

# Builds Binary
.PHONY: build
build:
	$(info building binary to bin/$(NAME))
	@CGO_ENABLED=0 go build -o bin/$(NAME) -ldflags '$(LDFLAGS)' ./cmd

# Run test suite
.PHONY: test
test:
	go test $(V) -cpu=1,2,8 --race ./...

# Clean all the things
.PHONY: clean
clean:
	@rm $(GOBIN)
