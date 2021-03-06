package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hecharles/doctortest/models"
)

//CustomerDoctorController simple version all doctor...
func CustomerDoctorController(c *fiber.Ctx) error {

	doctors, err := models.GetAllDoctorCustomer()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"data":       doctors,
	})

}
