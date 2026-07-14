package routes

import (
	"github.com/durgaprasad97005/bikeFinance/config"
	"github.com/gofiber/fiber/v3"
)

// Setup routes
func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// Authentication routes
	AuthRoutes(app, cfg)
}