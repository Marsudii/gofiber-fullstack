package config

import (
	"gofiber-fullstack/internal/handlers"
	"gofiber-fullstack/internal/repository"
	"gofiber-fullstack/internal/service"
)

type Config struct {
	AuthHandler *handlers.AuthHandlers
}

func NewConfig() *Config {
	// SETUP LOGIN DEPENDENCY INJECTION
	authRepo := repository.NewAuthRepo()
	authService := service.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// SETUP USER DEPENDENCY INJECTION

	return &Config{
		AuthHandler: authHandler,
	}
}
