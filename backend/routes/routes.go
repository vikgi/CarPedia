package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikas/car-model-app/backend/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/signup", controllers.SignUp)
	api.POST("/login", controllers.Login)
	api.GET("/cars/user", controllers.GetCarsByUser)
	api.POST("/cars/user", controllers.CreateCar);

}
