package routes

import (
	"github.com/conflux-tech/fiber-rest-boilerplate/handlers"
	"github.com/conflux-tech/fiber-rest-boilerplate/users"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoot registers all available routes for root
func RegisterRoot(route fiber.Router) {
	route.Get("/", handlers.GetStatus)
}

// RegisterUsers registers all available routes for /user
func RegisterUsers(route fiber.Router, pgRepo users.Repository) {
	route.Get("/", handlers.ListUsers(pgRepo))
	route.Get("/:id", handlers.GetUser(pgRepo))
	route.Post("/", handlers.CreateUser(pgRepo))
	route.Patch("/:id", handlers.UpdateUser(pgRepo))
	route.Delete("/:id", handlers.DeleteUser(pgRepo))
}
