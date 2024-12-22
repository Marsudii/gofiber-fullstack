package handlers

import "github.com/gofiber/fiber/v2"

// HomeClass handles the Home API logic
type HomeClass struct{}

// HomeHandler handles the `/home` API route
func (h *HomeClass) HomeHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to the Home API",
	})
}
