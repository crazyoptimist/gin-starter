# Gin Starter 

This is a Gin boilerplate organized in a modular way.

## Table Of Contents

## Environmental Variables

Following variables are required.
```
export APP_DSN="host=localhost user=admin password=password dbname=test port=5432 sslmode=disable TimeZone=America/Chicago"
export APP_API_KEY="dummy_key"
```

## API Documentation

- [gin-swagger](https://github.com/swaggo/gin-swagger) is used.
- To generate/update docs use `swag init` (install swag prior: `go install github.com/swaggo/swag/cmd/swag@latest`)
- To view docs, navigate to <http://localhost:8080/swagger/index.html>

## TODO

- [ ] Implement Auth
- [ ] Containerise
- [ ] Cleanup Tests
- [ ] Cleanup Lint, Vet
- [ ] Setup CI/CD
