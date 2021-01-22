package controllers

import (
	"encoding/json"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/hecharles/doctortest/config"
	"github.com/hecharles/doctortest/helpers"
	"github.com/hecharles/doctortest/models"
	"golang.org/x/crypto/bcrypt"
)

//LoginController handler
func LoginController(c *fiber.Ctx) error {

	// parse body
	var loginForm models.LoginForm

	err := json.Unmarshal([]byte(c.Body()), &loginForm)
	if err != nil {
		return c.Status(500).SendString("Parsing json Error :" + err.Error())
	}

	if len(loginForm.Email) < 3 || !helpers.IsEmailValid(loginForm.Email) {
		return c.Status(500).SendString("unsupport email")
	}
	if len(loginForm.Password) < 6 {
		return c.Status(500).SendString("A minimum password length greater than 5")
	}

	customer, err := models.GetCustomerByEmail(loginForm.Email)

	if err != nil {
		return c.Status(500).SendString("Email not registered")
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(loginForm.Password))
	if err != nil {
		return c.Status(500).SendString("Wrong password")

	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = customer.ID
	claims["Name"] = customer.Name
	claims["Gender"] = customer.Gender
	claims["Mobile"] = customer.Mobile
	claims["Email"] = customer.Email
	claims["exp"] = time.Now().Add(config.JwtExpireIn).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return c.Status(500).SendString("JWT sign Error :" + err.Error())
	}

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"token":      t,
	})

}
