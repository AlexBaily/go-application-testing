.PHONY: test lint fmt check build

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