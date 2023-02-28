test:
	go test -v ./...
db_migrate:
	go run ./cmd/api/migrate/default.go
docs_generate:
	rm -rf docs/* && swag init
build_linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/ ./...
