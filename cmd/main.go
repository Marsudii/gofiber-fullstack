package main

import (
	"gofiber-fullstack/internal/config"
	"gofiber-fullstack/internal/database"
	"gofiber-fullstack/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"os"
)

func main() {
	// call connectPostgreSQL
	database.ConnectPostgreSQL()

	// call config
	configDependencyInjection := config.NewConfig()

	// Initialize the HTML views
	htmlViewsInit := html.New("./internal/views", ".html")

	// Initialize the Fiber app
	app := fiber.New(
		fiber.Config{
			Views: htmlViewsInit,

			AppName: os.Getenv("APP_NAME"),
		})

	// Initialize load static files
	app.Static("/assets", "./internal/views/assets")

	// Streamlined logging route
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} - ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	// Set up routes
	routes.Routes(app, configDependencyInjection)

	// Run Listen on port 3000
	app.Listen(":3000")
}
