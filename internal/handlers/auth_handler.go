package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gofiber-fullstack/internal/service"
)

type AuthHandlers struct {
	authService service.AuthService
}

func NewAuthHandler(loginService service.AuthService) *AuthHandlers {
	return &AuthHandlers{
		authService: loginService,
	}
}

func (h *AuthHandlers) LoginView(c *fiber.Ctx) error {
	// Render login page
	return c.Render("auth/login", fiber.Map{
		"Title": "Login",
	})
}

func (h *AuthHandlers) LoginApi(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	//check if email is empty or password is empty then return to login page
	if username == "" || password == "" {
		// return to bad request
		return c.Status(400).JSON(fiber.Map{
			"error": "Username or Password is empty",
		})
	}

	// Call the login service
	return c.Status(200).JSON(fiber.Map{
		"data": "Login",
	})
}

func (h *AuthHandlers) RegisterView(c *fiber.Ctx) error {
	return c.Render("auth/register", fiber.Map{
		"Title": "Register",
	})
}

func (h *AuthHandlers) RegisterApi(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	data, err := h.authService.Login(email, password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": data,
	})
}
