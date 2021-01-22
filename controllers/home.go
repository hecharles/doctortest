package controllers

import (
	"github.com/gofiber/fiber/v2"
)

//HomeController handler
func HomeController(c *fiber.Ctx) error {

	return c.Status(200).SendString("Welcome, It's Work")

}
