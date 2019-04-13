GO ?= go

.PHONY: build test
.DEFAULT_GOAL := build

# Build
build:
$(GO) build -o main main.go

# Test
test:
$(GO) test -v ./...