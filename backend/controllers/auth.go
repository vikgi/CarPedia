package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vikas/car-model-app/backend/config"
	"github.com/vikas/car-model-app/backend/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func getUserCollection() *mongo.Collection {
	return config.DB.Collection("users")
}

var jwtKey = []byte("your_secret_key") 

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	log.Printf("debug: User struct: %+v\n", user)

	userCollection := getUserCollection()

	// Check if username already exists
	var existing models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"name": user.Name}).Decode(&existing)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	// Insert new user
	res, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Signup failed"})
		return
	}

	// Generate JWT token
	insertedID := res.InsertedID.(primitive.ObjectID).Hex()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": insertedID,
		"exp":    time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "User registered successfully",
		"username": user.Name,
		"userId":   insertedID,
		"token":    tokenString,
	})
}

func Login(c *gin.Context) {
	var loginData models.User
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	userCollection := getUserCollection()

	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"name": loginData.Name}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	//Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID.Hex(),
		"exp":    time.Now().Add(72 * time.Hour).Unix(), // token expires in 3 days
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Login successful",
		"username": user.Name,
		"userId":   user.ID.Hex(),
		"token":    tokenString,
	})
}

