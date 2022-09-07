package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"octaviusfarrel.dev/latihan_web/utils"
)

var secretKey string

func init() {
	gotenv.Load()
	secretKey = os.Getenv("TOKEN_SECRET")
}

func ReadRequiredTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if value := c.GetHeader("Token"); len(value) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"data":    "Token is empty",
			})
			return
		} else {
			token, returned := utils.ValidateToken(value)
			if !returned {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "invalid token",
				})
				return
			}
			value := strings.Split(value, "|")
			if len(value) < 2 {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "invalid token",
				})
				return
			}
			fmt.Println(token)

			if len(regexp.MustCompile(token).FindString("read")) == 0 {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "insufficient permission",
				})
				return
			}
		}
	}
}

func WriteRequiredTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if value := c.GetHeader("Token"); len(value) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"data":    "Token is empty",
			})
			return
		} else {
			token, returned := utils.ValidateToken(value)
			if !returned {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "invalid token",
				})
				return
			}
			value := strings.Split(value, "|")
			if len(value) < 2 {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "invalid token",
				})
				return
			}

			if len(regexp.MustCompile(token).FindString("write")) == 0 {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "insufficient permission",
				})
				return
			}
		}
	}
}

// func getToken(inputToken string) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return []byte(secretKey), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	claims, okClaim := token.Claims.(jwt.MapClaims)

// 	if !okClaim || !token.Valid {
// 		return nil, errors.New("invalid token type")
// 	}

// 	return claims, nil
// }
