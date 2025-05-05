package main

import (
	"github.com/vikas/car-model-app/backend/config"
	"github.com/vikas/car-model-app/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"fmt"
)

func main() {
	fmt.Println("Go backend is running 🎯")
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	config.Connect()
	routes.RegisterRoutes(r)
	r.Run(":5000")
}

