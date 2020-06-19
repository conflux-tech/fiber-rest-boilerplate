# Fiber REST Boilerplate

A boilerplate for creating REST APIs using Fiber

## Configuration

All configuration for the application such as HTTP Port, Database connection, etc. are set through environment variables. For development purpose, these environment variables can be set in a `.env` file. A sample environment variables can be found in `.env.example` file.

## Routing

Routing examples can be found in `/routes/route.go` file.

## Handlers

Example handlers can be found in `/app/handlers` directory.

## Database

We use GORM as an ORM to provide useful features to your models.

## Models

Models can be located in `/app/models` directory.

## Providers

Providers (custom helpers) can be found in `/app/providers`.
