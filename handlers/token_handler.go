package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString([]byte("test"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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

func ParseToken(c *gin.Context) {
	token, err := jwt.Parse(c.GetHeader("Token"), func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("test"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"data":    claims,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
			"data":    err.Error(),
		})
	}
}
