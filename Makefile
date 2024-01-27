test:
	go test -v $$(go list ./... | grep -v /docs)
vet:
	go vet -v $$(go list ./... | grep -v /docs)
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/server ./cmd/server
db_migrate:
	go run ./cmd/migrator/main.go
docs_generate:
	rm -rf docs/* && swag init -d ./cmd/server
