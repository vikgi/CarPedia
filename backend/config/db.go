package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// Connect initializes the MongoDB connection
func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Mongo client creation error:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	DB = client.Database("car_model_app")
	fmt.Println("✅ Connected to MongoDB!")
}


func GetCollection(collectionName string) *mongo.Collection {
    return DB.Collection(collectionName)
}