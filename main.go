package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hecharles/doctortest/config"
	"github.com/hecharles/doctortest/helpers"
	"github.com/hecharles/doctortest/router"
)

func main() {

	//initial data
	helpers.InitialSampleData()
	app := fiber.New()

	//setup router
	router.SetupRoutes(app)

	app.Listen(":" + config.ServerPort)
}
