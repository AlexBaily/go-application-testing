.PHONY: test lint fmt check run build image-build docker-up docker-down

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

run:
	@echo "Running server..."
	go run cmd/server/main.go

build:
	@echo "Running go build..."
	go build -o bin/go-application-testing ./cmd/server

image-build:
	@echo "Running image build..."
	docker build . -t $(IMAGE_NAME):$(IMAGE_TAG)

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down