# Fiber REST Boilerplate

A boilerplate for creating REST APIs using [Fiber](https://gofiber.io/)

## Running the application

To run the application locally, simply run the following command:

```bash
make dev
```

## Configurations

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

## Docker

You can run the application using Docker. To build and run the Docker image, you can use the following commands:

```bash
docker build -t fiber-rest-boilerplate .
docker run --name fiber-rest-boilerplate -p 3000:3000 fiber-reset-boilerplate
```

## Live Reloading

We used [Air](https://github.com/cosmtrek/air) to enable live reloading during development. This allows you to live reload your Go application when you make some changes on your Go files.
