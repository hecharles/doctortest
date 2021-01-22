package controllers

import (
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

//UserController handler
func UserController(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return c.JSON(fiber.Map{
		"ID":     claims["ID"].(string),
		"Name":   claims["Name"].(string),
		"Gender": claims["Gender"].(string),
		"Mobile": claims["Mobile"].(string),
		"Email":  claims["Email"].(string),
	})

}
