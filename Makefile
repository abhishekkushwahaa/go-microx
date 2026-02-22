BINARY_NAME=go-microx
VERSION=v1.0.0
BUILD_DIR=bin

.PHONY: all build clean install test help

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/go-microx/main.go

clean:
	@echo "Cleaning..."
	@if exist $(BUILD_DIR) rd /s /q $(BUILD_DIR)
	@go clean

install:
	@echo "Installing $(BINARY_NAME)..."
	@go install ./cmd/go-microx/main.go

test:
	@echo "Running tests..."
	@go test -v ./...

help:
	@echo "Makefile for $(BINARY_NAME)"
	@echo ""
	@echo "Usage:"
	@echo "  make build    Build the binary"
	@echo "  make clean    Remove build artifacts"
	@echo "  make install  Install the binary to GOPATH/bin"
	@echo "  make test     Run tests"
	@echo "  make help     Show this help message"
