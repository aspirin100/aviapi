POSTGRES_DSN := "postgres://postgres:postgres@localhost:5432/finapi?sslmode=disable"
M := ""

.PHONY: run
run:
	go run ./cmd/aviapi/main.go

.PHONY: commit
commit:
	git add . && \
	git commit -m "${M}"