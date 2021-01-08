package main

import (
	"github.com/conflux-tech/fiber-rest-boilerplate/configs"
	"github.com/conflux-tech/fiber-rest-boilerplate/database"
	"github.com/conflux-tech/fiber-rest-boilerplate/routes"
	"github.com/conflux-tech/fiber-rest-boilerplate/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	conf, _ := configs.Load()

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(requestid.New())

	if conf.Database.Enabled {
		database.Connect(&conf.Database)
		defer database.Close()
	}

	userPGRepo := users.NewPGRepo(database.Instance())

	routes.RegisterRoot(app.Group("/"))
	routes.RegisterUsers(app.Group("/users"), userPGRepo)

	app.Listen(conf.Application.Port)
}
