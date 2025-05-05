package middleware

import (
    "car-app-backend/utils"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
        userID, err := utils.ValidateToken(tokenStr)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Set("userId", userID)
        c.Next()
    }
}
