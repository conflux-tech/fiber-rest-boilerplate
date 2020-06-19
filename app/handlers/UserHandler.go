package handlers

import (
	"github.com/conflux-tech/fiber-rest-boilerplate/app/models"
	"github.com/conflux-tech/fiber-rest-boilerplate/database"
	"github.com/gofiber/fiber"
)

// GetAllUsers returns all users from db
func GetAllUsers(c *fiber.Ctx) {
	db := database.Instance()
	var Users []models.User
	if res := db.Find(&Users); res.Error != nil {
		c.Send("Error occured while retrieving users from the datbase", res.Error)
	}
	err := c.JSON(Users)
	if err != nil {
		panic("Error occurred when returning JSON of users")
	}
}

// GetUser returns information about a user
func GetUser(c *fiber.Ctx) {
	db := database.Instance()
	User := new(models.User)
	id := c.Params("id")
	if res := db.Find(&User, id); res.Error != nil {
		c.Send("Error occured while retrieving the user", res.Error)
	}
	if User.ID == 0 {
		c.SendStatus(404)
		err := c.JSON(fiber.Map{
			"ID": id,
		})
		if err != nil {
			panic("Error occurred when returning JSON")
		}
		return
	}
	err := c.JSON(User)
	if err != nil {
		panic("Error occured when returning JSON")
	}
}
