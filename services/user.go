package services

import (
	"context"
	"net/http"

	token_util "octaviusfarrel.dev/latihan_web/lib/token"
	"octaviusfarrel.dev/latihan_web/repositories/pgsql"
	"octaviusfarrel.dev/latihan_web/requests"
	"octaviusfarrel.dev/latihan_web/responses"
)

type IUserUsecase interface {
	GetUserByUsername(c context.Context, requestBody requests.UserRequest) (response responses.UserWithToken, statusCode int, err error)
	InsertUser(c context.Context, requestBody requests.UserRequest) (response responses.User, statusCode int, err error)
	UpdateUser(c context.Context) (response responses.User, statusCode int, err error)
	DeleteUser(c context.Context) (response responses.User, statusCode int, err error)
}

type UserUsecase struct {
	userRepo  pgsql.IUserRepo
	tokenUtil *token_util.TokenUtil
}

func NewUserUseCase(userRepo pgsql.IUserRepo, tokenUtil *token_util.TokenUtil) *UserUsecase {
	return &UserUsecase{userRepo: userRepo, tokenUtil: tokenUtil}
}

func (useCase *UserUsecase) GetUserByUsername(c context.Context, requestBody requests.UserRequest) (response responses.UserWithToken, statusCode int, err error) {

	user, err := useCase.userRepo.GetUserByUsername(requestBody)

	if err != nil {
		statusCode = http.StatusBadRequest
		responses.NewBaseResponseStatusCode(statusCode, &response.BaseResponse, err)

		return
	}

	token, err := useCase.tokenUtil.CreateToken(user)

	if err != nil {
		statusCode = http.StatusInternalServerError
		responses.NewBaseResponseStatusCode(statusCode, &response.BaseResponse, err)

		return
	}

	// token, err = useCase.tokenRepo.InsertTokenByUser(int(user.ID), token)

	// if err != nil {
	// 	statusCode = http.StatusInternalServerError
	// 	responses.NewBaseResponseStatusCode(statusCode, &response.BaseResponse, err)

	// 	return
	// }

	statusCode = http.StatusOK
	response.Token = token
	response.User.User = user
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
