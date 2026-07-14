package main

import (
	"github.com/durgaprasad97005/bikeFinance/config"
	"github.com/durgaprasad97005/bikeFinance/database"
	"github.com/durgaprasad97005/bikeFinance/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// Setting up app
func SetupApp(cfg *config.Config) *fiber.App {
	// Creating a new app
	app := fiber.New()

	// Public middleware
	app.Use(logger.New())
	app.Use(cors.New())
	
	// Connect to Database
	database.ConnectDatabase(cfg)

	// Initialize Routes
	routes.SetupRoutes(app, cfg)

	return app
}