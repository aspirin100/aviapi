POSTGRES_DSN := "postgres://postgres:postgres@localhost:5432/finapi?sslmode=disable"

.PHONY: run
run:
	go run ./cmd/aviapi/main.go

