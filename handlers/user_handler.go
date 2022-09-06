package handlers

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/utils"
)

func Register(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.Bind(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": err.Error()})
		return
	}

	data := struct {
		Username   string
		Permission string
	}{
		Username:   requestBody.Username,
		Permission: "",
	}

	utils.CreateUser(data, requestBody.Password)

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "User Created"})
}

func Login(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.Bind(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": err.Error()})
		return
	}

	user, err := utils.GetUserWithPassword(requestBody.Username, requestBody.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	if token, err := createToken(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Failed", "message": err.Error()})
	} else {
		if token, err := utils.InsertTokenByUser(user.Id, token); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed", "message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "Success", "user": user, "token": token})
		}
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func createToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"created_at": time.Now().Unix(),
		"salt":       randomString(24),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
