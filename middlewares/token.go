package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	token "octaviusfarrel.dev/latihan_web/lib/token"
)

type ITokenMiddleware interface {
	ReadToken() gin.HandlerFunc
	WriteToken() gin.HandlerFunc
}

type TokenMiddleware struct {
	tokenUtil *token.TokenUtil
}

func NewTokenMiddleware(tokenUtil *token.TokenUtil) *TokenMiddleware {
	return &TokenMiddleware{tokenUtil: tokenUtil}
}

func (tokenMiddleware *TokenMiddleware) ReadToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("Authorization")
		if len(value) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"data":    "Token is empty",
			})
			c.Abort()
			return
		}

		valueArray := strings.Split(value, " ")
		if len(valueArray) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"data":    "invalid token format",
			})
			c.Abort()
			return
		}
		value = valueArray[1]

		err := tokenMiddleware.tokenUtil.ValidateToken(value, "read")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"data":    "invalid token permission",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (tokenMiddleware *TokenMiddleware) WriteToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("Authorization")
		if len(value) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"data":    "Token is empty",
			})
			c.Abort()
			return
		}

		valueArray := strings.Split(value, " ")
		if len(valueArray) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"data":    "invalid token format",
			})
			c.Abort()
			return
		}
		value = valueArray[1]

		err := tokenMiddleware.tokenUtil.ValidateToken(value, "write")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"data":    "invalid token permission",
			})
			c.Abort()
			return
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
