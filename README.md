# E-commerce API (Go + Gin + GORM)

A simple e-commerce REST API built with Go, Gin, and GORM. It supports authentication, product management, cart operations, and order processing. This project is organized with handlers, middleware, models, services, and routes.

## Features
- Authentication: register, login, JWT-based auth
- Product catalog: list, view, create, update, delete
- Cart: add/remove items, view user cart
- Orders: create order, list user orders, get order by id
- Admin-only product management
- CORS, rate limiting, and basic error middleware

## Getting Started

### Prerequisites
- Go 1.20+
- A database supported by GORM (Postgres/MySQL/SQLite). SQLite works out of the box.

### Configuration
- Copy `env-example` to `.env` and adjust values.
- Database connection is configured in `config/database.go`.

### Run Locally
```bash
# From project root
go run ./cmd/main.go
```

### Run Tests
```bash
# Runs unit tests (uses sqlite in-memory)
go test ./...
```

## Project Structure
```
cmd/            # entrypoint
config/         # config and DB setup
database/       # migrations
handlers/       # HTTP handlers
middleware/     # auth, admin, cors, rate limit, error
models/         # data models
routes/         # route wiring
services/       # business logic layer (new)
tests/          # unit tests
utils/          # helpers (jwt, encryption, validation)
```

## Services Layer
- Product service is implemented in `services/product_service.go`.
- Handlers delegate to services to keep business logic centralized and testable.

## API Docs
See `docs/API.md` for endpoints, parameters, and response examples.

## Notes
- CORS is enabled via `middleware.CORSMiddleware()` in `cmd/main.go`.
- Admin routes are protected by `AuthMiddleware` and `AdminMiddleware`.
- For production, configure a persistent database and migrations accordingly.