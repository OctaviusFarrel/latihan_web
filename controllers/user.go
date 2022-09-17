package controllers

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/requests"
	"octaviusfarrel.dev/latihan_web/responses"
	"octaviusfarrel.dev/latihan_web/services"
)

type UserHandler struct {
	userUseCase services.IUserUsecase
}

func NewUserHandler(userUseCase services.IUserUsecase) *UserHandler {
	return &UserHandler{userUseCase}
}

func (userHandler *UserHandler) InsertUser(c *gin.Context) {
	ctx := c.Request.Context()
	var requestBody requests.UserRequest

	if err := c.Bind(&requestBody); err != nil {

		res := responses.BaseResponse{}
		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	res, code, err := userHandler.userUseCase.InsertUser(ctx, requestBody)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &res.BaseResponse, err)

		c.JSON(code, res)
		return
	}

	c.JSON(code, res)
}

func (userHandler *UserHandler) GetUserByUsername(c *gin.Context) {
	ctx := c.Request.Context()

	var requestBody requests.UserRequest

	err := c.Bind(&requestBody)
	if err != nil {
		res := responses.BaseResponse{}
		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	res, code, err := userHandler.userUseCase.GetUserByUsername(ctx, requestBody)

	if err != nil {
		responses.NewBaseResponseStatusCode(code, &res.BaseResponse, err)

		c.JSON(code, res)
		return
	}

	c.JSON(code, res)
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

func createToken(user models.UserModel) (string, error) {
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
