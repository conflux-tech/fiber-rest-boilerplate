package handlers

import (
	"net/http"

	"github.com/conflux-tech/fiber-rest-boilerplate/app/models"
	"github.com/conflux-tech/fiber-rest-boilerplate/database"
	"github.com/gofiber/fiber"
)

// GetAllUsers returns all users from db
func GetAllUsers(c *fiber.Ctx) {
	db := database.Instance()
	var Users []models.User
	if res := db.Find(&Users); res.Error != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(fiber.Map{
			"error": fiber.Map{
				"message": "An error occured while retrieiving data",
				"details": res.Error,
			},
		})
	}
	c.JSON(Users)
}

// GetUser returns information about a user
func GetUser(c *fiber.Ctx) {
	db := database.Instance()
	User := new(models.User)
	id := c.Params("id")
	if res := db.Find(&User, id); res.Error != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(fiber.Map{
			"error": fiber.Map{
				"message": "An error occured while retrieiving data",
				"details": res.Error,
			},
		})
	}
	if User.ID == 0 {
		c.Status(http.StatusNotFound)
		err := c.JSON(fiber.Map{
			"ID": id,
		})
		if err != nil {
			panic("An error occurred when returning JSON")
		}
		return
	}
	c.JSON(User)
}

// AddUser creates new user record
func AddUser(c *fiber.Ctx) {
	db := database.Instance()
	User := new(models.User)
	if err := c.BodyParser(User); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(fiber.Map{
			"error": fiber.Map{
				"message": "An error occurred while parsing the request",
				"details": err,
			},
		})
	}
	if res := db.Create(&User); res.Error != nil {
		c.Send("An error occured while creating record", res.Error)
	}
	c.JSON(User)
}

// UpdateUser updates a user record
func  UpdateUser(c *fiber.Ctx) {
	db := database.Instance()
	id := c.Params("id")
	EditUser := new(models.User)
	User := new(models.User)

	if err := c.BodyParser(EditUser); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(fiber.Map{
			"error": fiber.Map{
				"message": "An error occurred while parsing the request",
				"details": err,
			},
		})
	}
	if res := db.Find(&User, id); res.Error != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(fiber.Map{
			"error": fiber.Map{
				"message": "An error occured while retrieiving data",
				"details": res.Error,
			},
		})
	}

	if User.ID == 0 {
		c.SendStatus(404)
		return
	}

	User.Name = EditUser.Name

	db.Save(&User)

	c.JSON(User)
}

// DeleteUser deletes a use
func DeleteUser(c *fiber.Ctx) {
	db := database.Instance()
	id := c.Params("id")

	var User models.User
	db.Find(&User, id)
	if res := db.Find(&User); res.Error != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(fiber.Map{
			"error": fiber.Map{
				"message": "An error occured while retrieiving data",
				"details": res.Error,
			},
		})
	}
	db.Delete(&User)

	c.JSON(fiber.Map{
		"ID": id,
		"Delete": true,
	})
}
