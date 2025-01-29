# ---- build, run, and test
run:
	go run ./cmd/main

dev:
	air server

build:
	go build -o ./app ./cmd/main


# ---- code quality
lint:
	golangci-lint run ./...

format:
	go fmt ./...

test:
	go test ./internal/... ./pkg/...

coverage:
	mkdir -p coverage
	go test -coverprofile=coverage/coverage.out "./internal/..." "./pkg/..."
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html


# ---- generators
mock:
	mockery

docs:
	swag init -g ./cmd/main/main.go -o ./api

# ---- dependencies
tidy:
	go mod tidy

install:
	go install github.com/air-verse/air@v1.61.7
	go install github.com/vektra/mockery/v2@v2.44.1
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4
	go install github.com/swaggo/swag/cmd/swag@v1.8.11
	go mod download

.PHONY: run dev build bump lint format test coverage mock swagger docs tidy install
