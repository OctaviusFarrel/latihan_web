package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/subosito/gotenv"
)

var secretKey string

func init() {
	gotenv.Load()
	secretKey = os.Getenv("TOKEN_SECRET")
}

func GetReadToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"read": true,
		"bruh": true,
	})
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   tokenString,
	})
}

func GetWriteToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"read": true,
	})
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   tokenString,
	})
}
