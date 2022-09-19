package middlewares

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"octaviusfarrel.dev/latihan_web/utils"
)

func ReadRequiredTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if value := c.GetHeader("Token"); len(value) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"data":    "Token is empty",
			})
			c.Abort()
			return
		} else {
			token, returned := utils.ValidateToken(value, c.Request.Context())
			if !returned {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "invalid token",
				})
				c.Abort()
				return
			}
			value := strings.Split(value, "|")
			if len(value) < 2 {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "invalid token",
				})
				c.Abort()
				return
			}

			if len(regexp.MustCompile(token).FindString("read")) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "insufficient permission",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func WriteRequiredTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if value := c.GetHeader("Token"); len(value) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"data":    "Token is empty",
			})
			c.Abort()
			return
		} else {
			token, returned := utils.ValidateToken(value, c.Request.Context())
			if !returned {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "invalid token",
				})
				c.Abort()
				return
			}
			value := strings.Split(value, "|")
			if len(value) < 2 {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "invalid token",
				})
				c.Abort()
				return
			}

			if len(regexp.MustCompile(token).FindString("write")) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Failed",
					"data":    "insufficient permission",
				})
				c.Abort()
				return
			}
		}
		c.Next()
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
