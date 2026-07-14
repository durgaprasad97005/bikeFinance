package main

import (
	"log"

	"github.com/durgaprasad97005/bikeFinance/config"
)

// Main function
func main() {
	// Accessing environment variables
	cfg := config.Load()

	// Setting up app
	app := SetupApp(cfg)

	// Port safe check
	if cfg.Port == "" {
		cfg.Port = "3000"
	}

	// Listen at the given port
	log.Println("The server is running at port:" + cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}