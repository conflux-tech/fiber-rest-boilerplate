# Fiber REST Boilerplate

A boilerplate for creating REST APIs using [Fiber](https://gofiber.io/)

## Running locally

```bash
> docker-compose up

# ...

fiber_rest_api | running...
fiber_rest_api |
fiber_rest_api |  ┌───────────────────────────────────────────────────┐
fiber_rest_api |  │                    Fiber v2.3.2                   │
fiber_rest_api |  │               http://127.0.0.1:3000               │
fiber_rest_api |  │                                                   │
fiber_rest_api |  │ Handlers ............ 13  Processes ........... 1 │
fiber_rest_api |  │ Prefork ....... Disabled  PID ............... 149 │
fiber_rest_api |  └───────────────────────────────────────────────────┘
fiber_rest_api |
```

## Configurations

All configuration for the application such as HTTP Port, Database connection, etc. are set through environment variables. For development purpose, these environment variables can be set in a `.env` file. A sample environment variables can be found in `.env.example` file.

## Routes

Routes examples can be found in `/api/routes/routes.go` file.

## Handlers

Example handlers can be found in `/api/handlers` directory.

## Models

Models can be located in `/api/models` directory. We use [GORM](https://gorm.io/) as an ORM to provide useful features to your models.

## Database Migration

We use [Goose](https://github.com/pressly/goose) as a database migration tool. Run `./scripts/migration` for more information.

Sample migration files can be found in `/database/migrations`.
