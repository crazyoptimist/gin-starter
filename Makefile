test:
	go test -v ./...
vet:
	go vet -v ./...
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/ ./...
db_migrate:
	go run ./cmd/api/migrate/default.go
docs_generate:
	rm -rf docs/* && swag init
