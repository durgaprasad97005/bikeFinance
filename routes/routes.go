package routes

import (
	"github.com/durgaprasad97005/bikeFinance/middleware"
	"github.com/gofiber/fiber/v3"
)

// Setup routes
func SetupRoutes(app *fiber.App, jwtSecret string) {
	// Auth middleware
	authMiddleware := middleware.NewAuthMiddleware(jwtSecret)

	// Registering Routes
	AuthRoutes(app, authMiddleware, jwtSecret)
}