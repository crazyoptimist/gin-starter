# Gin Starter 

Gin boilerplate organized in a modular way.

[![Build & Test](https://github.com/crazyoptimist/gin-starter/actions/workflows/test.yml/badge.svg)](https://github.com/crazyoptimist/gin-starter/actions/workflows/test.yml)

## Table Of Contents

## DB Migration

```bash
make db_migrate
```

## Development

Create a dotenv file:

```bash
cp .env.example .env
```

Install [air](https://github.com/cosmtrek/air) for live reloading. Air config file is already inside the repo, so simply run:

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

To browse docs, open `BASE_URL/swagger/index.html`.

Generate/update docs:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
make docs_generate
```

## TODO

- [ ] Containerize
- [ ] Configure logger using zap

## License

MIT
