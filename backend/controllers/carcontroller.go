package controllers

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/vikas/car-model-app/backend/config"
    "github.com/vikas/car-model-app/backend/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateCar(c *gin.Context) {
	userID := c.Query("userId")
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userCollection := config.GetCollection("users")
	carCollection := config.GetCollection("cars")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if user exists
	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var car models.Car
	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	car.ID = primitive.NewObjectID()
	car.UserID = oid

	_, err = carCollection.InsertOne(ctx, car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add car"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car added successfully", "car": car})
}


func GetCarsByUser(c *gin.Context) {
    userID := c.Query("userId")
    oid, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    collection := config.GetCollection("cars")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var cars []models.Car
    cursor, err := collection.Find(ctx, bson.M{"userId": oid})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching cars"})
        return
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var car models.Car
        if err := cursor.Decode(&car); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Decode error"})
            return
        }
        cars = append(cars, car)
    }

    c.JSON(http.StatusOK, cars)
}
