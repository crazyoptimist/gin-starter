# Gin Starter

This is a REST API boilerplate.

It uses:
- Gin for routing
- Gorm with Postgres database
- Viper for configuration management
- JWT for authentication
- Swagger for API documentation
- Air for live reloading in development
- Redis for caching

## Development

Install [air](https://github.com/cosmtrek/air) for live reloading.

```bash
go install github.com/cosmtrek/air@latest
```

```bash
air
```

## Test

```bash
make test
```

## Build

```bash
make build
```

Binaries will be generated inside `PROJECT_ROOT/bin/`

## DB Migration

```bash
make migrate
```

## API Documentation

[gin-swagger](https://github.com/swaggo/gin-swagger) is used for API documentation.

To browse the API documentation, open `BASE_URL/swagger/index.html`.

Generate/update docs:

```bash
go install github.com/swaggo/swag/cmd/swag@latest

make docs
```

## License

MIT

Made with :heart: by crazyoptimist
