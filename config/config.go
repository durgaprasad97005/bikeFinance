package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	MongoURI     string
	DatabaseName string
	JwtSecret    string
}

// Load environment variables
func Load() *Config {
	// Loading .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load .env file. Using server environment variables...")
	}

	return &Config{
		Port: os.Getenv("PORT"), 
		MongoURI: os.Getenv("MONGO_URI"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}