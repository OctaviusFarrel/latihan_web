package services

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/subosito/gotenv"
	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/repositories/pgsql"
	"octaviusfarrel.dev/latihan_web/requests"
	"octaviusfarrel.dev/latihan_web/responses"
)

var secretKey = func() (result string) {
	gotenv.Load()
	result = os.Getenv("TOKEN_SECRET")
	return
}()

type IUserUsecase interface {
	GetUserByUsername(c context.Context, requestBody requests.UserRequest) (response responses.UserWithToken, statusCode int, err error)
	InsertUser(c context.Context, requestBody requests.UserRequest) (response responses.User, statusCode int, err error)
	UpdateUser(c context.Context) (response responses.User, statusCode int, err error)
	DeleteUser(c context.Context) (response responses.User, statusCode int, err error)
}

type UserUsecase struct {
	userRepo  pgsql.IUserRepo
	tokenRepo pgsql.ITokenRepo
}

func NewUserUseCase(userRepo pgsql.IUserRepo, tokenRepo pgsql.ITokenRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo, tokenRepo: tokenRepo}
}

func (useCase *UserUsecase) GetUserByUsername(c context.Context, requestBody requests.UserRequest) (response responses.UserWithToken, statusCode int, err error) {

	user, err := useCase.userRepo.GetUserByUsername(requestBody)

	if err != nil {
		statusCode = http.StatusBadRequest
		responses.NewBaseResponseStatusCode(statusCode, &response.BaseResponse, err)

		return
	}
	token, err := createToken(user)
	if err != nil {
		statusCode = http.StatusInternalServerError
		responses.NewBaseResponseStatusCode(statusCode, &response.BaseResponse, err)

		return
	}

	token, err = useCase.tokenRepo.InsertTokenByUser(user.Id, token)

	if err != nil {
		statusCode = http.StatusInternalServerError
		responses.NewBaseResponseStatusCode(statusCode, &response.BaseResponse, err)

		return
	}

	statusCode = http.StatusOK
	response.Token = token
	responses.NewBaseResponseStatusCode(statusCode, &response.BaseResponse, err)

	return
}

func (useCase *UserUsecase) InsertUser(c context.Context, requestBody requests.UserRequest) (response responses.User, statusCode int, err error) {
	statusCode = http.StatusInternalServerError

	user, err := useCase.userRepo.CreateUser(requestBody)
	if err == nil {
		statusCode = http.StatusOK
		response.User = user
		return
	} else {
		statusCode = http.StatusBadRequest
	}
	responses.NewBaseResponseStatusCode(statusCode, &response.BaseResponse, err)

	return
}

func (useCase *UserUsecase) UpdateUser(c context.Context) (response responses.User, statusCode int, err error) {
	return
}

func (useCase *UserUsecase) DeleteUser(c context.Context) (response responses.User, statusCode int, err error) {
	return
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
