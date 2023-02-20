# Gin Starter 

This is a Gin boilerplate organized in a modular way.

## Table Of Contents

## Environment Variables

Following variables are required.

```
export APP_DSN="host=localhost user=admin password=password dbname=test port=5432 sslmode=disable TimeZone=America/Chicago"
```

## API Documentation

[gin-swagger](https://github.com/swaggo/gin-swagger) is used for API documentation.

To browse docs, open `BASE_URL/swagger/index.html`.

Generate/update docs:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

## TODO

- [ ] Implement Auth
- [ ] Containerize
- [ ] Cleanup Tests
- [ ] Cleanup Lint, Vet
- [ ] Setup CI/CD
