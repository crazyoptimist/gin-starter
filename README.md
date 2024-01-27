# Gin Starter

## DB Migration

```bash
make db_migrate
```

## Development

Create `.env` file.

```bash
cp .env.example .env
```

Use [air](https://github.com/cosmtrek/air) for live reloading.


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

## API Documentation

[gin-swagger](https://github.com/swaggo/gin-swagger) is used for API documentation.

To browse the API documentation, open `BASE_URL/swagger/index.html`.

Generate/update docs:

```bash
go install github.com/swaggo/swag/cmd/swag@latest

make docs_generate
```

## License

MIT
