package handlers

import (
	"net/http"

	"github.com/conflux-tech/fiber-rest-boilerplate/configs"
	"github.com/gofiber/fiber/v2"
)

// GetStatus shows status of application
func GetStatus(c *fiber.Ctx) error {
	cfg := configs.Get()
	return c.JSON(fiber.Map{
		"env":    cfg.Application.Env,
		"status": http.StatusText(http.StatusOK),
	})
}
