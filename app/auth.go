package app

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

)

func (a *App) CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user": userID,
		"exp":  time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.Key)
}

func (a *App) Check(c *gin.Context) (int, bool) {
	tokenStr := c.GetHeader("Authorization")

	if tokenStr == "" {
		c.JSON(401, gin.H{"error": "missing token"})
		return 0, false
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return a.Key, nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "invalid token"})
		return 0, false
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["user"].(float64))

	return userID, true
}
