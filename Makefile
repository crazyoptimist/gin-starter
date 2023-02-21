db_migrate:
	go run ./cmd/api/migration/main.go
docs_generate:
	rm -rf docs/* && swag init
docs_format:
	swag fmt -d cmd/api internal
build_linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/ ./...
