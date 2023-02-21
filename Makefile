db_migrate:
	go run ./cmd/api/migration/main.go
docs:
	swag init && swag fmt -d internal
build_linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/ ./...
