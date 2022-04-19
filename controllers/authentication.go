package controllers

import (
	database "github.com/weldonla/OneCauseGolangLogin/databases"
	"github.com/weldonla/OneCauseGolangLogin/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	token := data["token"]

	// I would like to put other validation of the token here, but there would be a lot of edge cases so it's a little out of scope.
	if len(token) != 4 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Token",
		})
	}

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	return c.JSON(fiber.Map{
		"token":    token,
		"id":       user.Id,
		"name":     user.Name,
		"username": user.UserName,
		"email":    user.Email,
		"password": user.Password,
	})
}
