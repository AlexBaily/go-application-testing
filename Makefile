.PHONY: test lint fmt check build image-build

IMAGE_NAME ?= go-application-testing
IMAGE_TAG ?= latest

fmt:
	@echo "Running go fmt..."
	gofmt -s -w .

lint:
	@echo "Running go lint..."
	golangci-lint run ./...

test:
	@echo "Running go tests..."
	go test -v -race ./...

check: fmt lint test

build:
	@echo "Running go build..."
	go build -o bin/go-application-testing .

image-build:
	@echo "Running go build..."
	docker build . -t $(IMAGE_NAME):$(IMAGE_TAG)