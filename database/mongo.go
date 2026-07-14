package database

import (
	"context"
	"log"
	"time"

	"github.com/durgaprasad97005/bikeFinance/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	UsersCollection = "users"
)

var DB *mongo.Database

// function to connect to database
func ConnectDatabase(cfg *config.Config) {
	// Connect to Database
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Println("Error connecting MongoDB: " + err.Error())
	}

	// Ping database for connection check
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Println("Error while executing Ping(): " + err.Error())
	}

	// Getting Database
	DB = client.Database(cfg.DatabaseName)
}

// function to get collection from database
// func GetCollection(name string) *mongo.Collection {
// 	return DB.Collection(name)
// }