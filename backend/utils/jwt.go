package utils

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("CarPedia_privateKey")

func GenerateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "userId": userID,
        "exp":    time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (string, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims["userId"].(string), nil
    }

    return "", err
}
