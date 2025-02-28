POSTGRES_DSN := "postgres://postgres:postgres@localhost:5432/aviapi_db?sslmode=disable"
MIGRATION_NAME := ""

.PHONY: run
run:
	go run ./cmd/aviapi/main.go

DEFAULT-GOAL: run

.PHONY: up
up: 
	docker-compose up -d

goose-create:
	go run github.com/pressly/goose/v3/cmd/goose@latest \
	-dir ./internal/repository/migrations create ${MIGRATION_NAME} sql

goose-up:
	go run github.com/pressly/goose/v3/cmd/goose@latest \
	-dir ./internal/repository/migrations postgres $(POSTGRES_DSN) up

goose-down:
	go run github.com/pressly/goose/v3/cmd/goose@latest \
	-dir ./internal/repository/migrations postgres $(POSTGRES_DSN) down

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

build:
	mkdir -p bin && \
	go build -o ./bin/aviapi-server ./cmd/aviapi/main.go

build-img:
	docker build -t aviapi-img .