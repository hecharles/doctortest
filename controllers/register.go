package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	"github.com/hecharles/doctortest/helpers"
	"github.com/hecharles/doctortest/models"
	"golang.org/x/crypto/bcrypt"
)

//LoginController handler
func RegisterController(c *fiber.Ctx) error {

	// parse body
	var registerForm models.RegisterForm

	err := json.Unmarshal([]byte(c.Body()), &registerForm)
	if err != nil {
		return c.Status(500).SendString("Parsing json Error :" + err.Error())
	}

	if len(registerForm.Email) < 3 || !helpers.IsEmailValid(registerForm.Email) {
		return c.Status(500).SendString("unsupport email")
	}
	if len(registerForm.Password) < 6 {
		return c.Status(500).SendString("A minimum password length greater than 5")
	}
	if registerForm.Password != registerForm.RePassword {
		return c.Status(500).SendString("Password and Confirm Password must equal")
	}

	if len(registerForm.Name) < 1 {
		return c.Status(500).SendString("Name cannot be empty")
	}

	if len(registerForm.Mobile) < 1 {
		return c.Status(500).SendString("Mobile cannot be empty")
	}

	if registerForm.Gender != "Male" && registerForm.Gender != "Female" {
		return c.Status(500).SendString("Gender not valid")
	}

	customer, err := models.GetCustomerByEmail(registerForm.Email)

	if err == nil && customer != nil {
		return c.Status(500).SendString("Email already registered")
	}

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerForm.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).SendString("hash password ERROR:" + err.Error())

	}

	newCustomer := models.Customer{
		ID:       guuid.New().String(),
		Name:     registerForm.Name,
		Gender:   registerForm.Gender,
		Mobile:   registerForm.Mobile,
		Email:    registerForm.Email,
		Password: string(hashedPassword),
	}

	err = models.CreateNewCustomer(&newCustomer)
	if err != nil {
		return c.Status(500).SendString("create customer ERROR:" + err.Error())

	}

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"message":    "Success register",
	})

}
