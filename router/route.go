package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hecharles/doctortest/controllers"
	"github.com/hecharles/doctortest/middleware"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// routes root
	app.Get("/", controllers.HomeController)
	app.Post("/api/login", controllers.LoginController)
	app.Post("/api/register", controllers.RegisterController)

	// Group Public Api
	publicAPI := app.Group("/api/public")
	publicAPI.Get("/doctor", controllers.PublicDoctorController)

	// Group Customer Api
	customerAPI := app.Group("/api/customer")

	// JWT check middleware
	customerAPI.Use(middleware.AuthRequired())

	customerAPI.Get("/user", controllers.UserController)

	customerAPI.Get("/doctor", controllers.CustomerDoctorController)

}
