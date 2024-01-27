test:
	go test -v ./...
vet:
	go vet -v ./...
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/ ./...
db_migrate:
	go run ./cmd/migrator/main.go
docs_generate:
	rm -rf docs/* && swag init -d ./cmd/server
