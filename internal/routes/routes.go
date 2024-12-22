package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiber-fullstack/internal/config"
)

func Routes(app *fiber.App, config *config.Config) {

	// Views Routes
	app.Get("/", config.AuthHandler.LoginView)
	app.Get("/register", config.AuthHandler.RegisterView)

	// API Routes
	apiV1Group := app.Group("/api/v1")
	apiV1Group.Post("/login", config.AuthHandler.LoginApi)
	apiV1Group.Post("/register", config.AuthHandler.RegisterApi)
	// API MIDDLEWARE CHECK TOKEN

}
