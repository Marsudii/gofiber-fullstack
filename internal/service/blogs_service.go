package service

import "github.com/gofiber/fiber/v2"

// BlogsService handles the Blogs API logic

type BlogsService struct{}

// BlogsHandler handles the `/blogs` API route
func (h *BlogsService) BlogsHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to the Blogs API",
	})
}
