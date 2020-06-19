package handlers

import (
	"github.com/conflux-tech/fiber-rest-boilerplate/app/providers"
	"github.com/gofiber/fiber"
)

// GetStatus shows status of application
func GetStatus(c *fiber.Ctx) {
	conf := providers.GetConfig()
	c.JSON(fiber.Map{
		"env": conf.Application.Env,
		"status": "OK",
	})
}
