package routes

import (
	"gofiber-fullstack/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	// Views Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{
			"Title": "Home",
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"Title": "Home",
		})
	})

	app.Get("/blogs", func(c *fiber.Ctx) error {
		return c.Render("blogs", fiber.Map{
			"Title": "Home",
		})
	})

	// API Routes
	apiV1Group := app.Group("/api/v1")

	// Home API Init
	homeHandler := handlers.HomeClass{}
	// Home API Route
	apiV1Group.Get("/home", homeHandler.HomeHandler)

}
