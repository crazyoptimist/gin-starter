db_migrate:
	go run ./cmd/api/migration/main.go
build_linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/ ./...
