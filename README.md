# Gin Starter

This is a REST API boilerplate.

## Overview

This project implements the Clean Architecture proposed by Uncle Bob. I do not claim this to be a perfect implementation, but it makes sense to me and hopefully to you. It's important to note that Clean Architecture is not about directory structures, however, it significantly influences them to some extent.

The project has been organized adhering to several principles: Single Responsibility Principle (SRP), Common Closure Principle (CCP), Dependency Inversion Principle (DIP), and Acyclic Dependency Principle (ADP).

- `internal/domain/model`: This directory contains the entities, also known as the core business rules.
- Other directories within `internal/domain`: These hold the application-specific business rules.
- `internal/infrastructure`: This directory houses the low level details of the system.

In this project, components such as the web server, router, and controller are details, obviously. The `repository` is a database access gateway, thus a detail. Any component that interacts with external services, like a Redis server or a third-party API, is also a detail.

Details can depend on application business rules, but not vice versa.
Application business rules can depend on core business rules, but not vice versa.

We may need to utilize some external service in application business rules; Dependency Inversion comes to the rescue!

That concludes the overview for now. I believe this implementation is truly opinionated, though your perspective may differ.

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
