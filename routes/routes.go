package routes

import (
	"github.com/conflux-tech/fiber-rest-boilerplate/app/handlers"
	"github.com/gofiber/fiber"
)

// RegisterRoutes registers all available routes
func RegisterRoutes(route *fiber.Group) {
	registerRoot(route)
	registerUsers(route)
}

func registerRoot(route *fiber.Group) {
	route.Get("/", handlers.GetStatus)
}

func registerUsers(route *fiber.Group) {
	users := route.Group("/users")
	users.Get("/", handlers.GetAllUsers)
}
