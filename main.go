package main

import (
	"github.com/conflux-tech/fiber-rest-boilerplate/app/providers"
	"github.com/conflux-tech/fiber-rest-boilerplate/config"
	"github.com/conflux-tech/fiber-rest-boilerplate/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	conf, _ := config.LoadConfigs()

	app := fiber.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.Compress())
	app.Use(middleware.RequestID())

	routes.RegisterRoutes(app.Group("/"))

	providers.SetConfig(&conf)

	app.Listen(conf.Application.Port)
}
