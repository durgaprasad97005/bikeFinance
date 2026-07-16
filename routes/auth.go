package routes

import (
	"github.com/durgaprasad97005/bikeFinance/database"
	"github.com/durgaprasad97005/bikeFinance/handlers"
	"github.com/durgaprasad97005/bikeFinance/repository"
	"github.com/durgaprasad97005/bikeFinance/services"
	"github.com/gofiber/fiber/v3"
)

// Function to create authentication routes
func AuthRoutes(app *fiber.App, authMiddleware fiber.Handler, jwtSecret string) {
	// Initialize configured repository, services, and handlers
	userRepo := repository.NewUserRepository(database.DB)
	authService := services.NewAuthService(userRepo, jwtSecret)
	authHandler := handlers.NewAuthHandler(authService)

	// Grouping of routes
	auth := app.Group("/api/auth")

	// Auth routes
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Get("/profile", authMiddleware, authHandler.Profile)
	// Pending - Need to add two more routes - Logout, GetProfile
}