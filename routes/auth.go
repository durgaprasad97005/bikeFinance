package routes

import (
	"github.com/durgaprasad97005/bikeFinance/config"
	"github.com/durgaprasad97005/bikeFinance/database"
	"github.com/durgaprasad97005/bikeFinance/handlers"
	"github.com/durgaprasad97005/bikeFinance/repository"
	"github.com/durgaprasad97005/bikeFinance/services"
	"github.com/gofiber/fiber/v3"
)

// Function to create authentication routes
func AuthRoutes(app *fiber.App, cfg *config.Config) {
	// Initialize configured repository, services, and handlers
	userRepo := repository.NewUserRepository(database.DB)
	authService := services.NewAuthService(userRepo, cfg)
	authHandler := handlers.NewAuthHandler(authService)

	// Grouping of routes
	auth := app.Group("/api/auth")

	// Auth routes
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	// Pending - Need to add two more routes - Logout, GetProfile
}