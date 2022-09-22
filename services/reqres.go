package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/requests"
	"octaviusfarrel.dev/latihan_web/responses"
)

type IReqresUseCase interface {
	AllUsers(context context.Context) (code int, response responses.AllReqresUsers, err error)
	GetUser(context context.Context, index int) (code int, response responses.ReqresUser, err error)
	InsertUser(context context.Context, request requests.ReqresUserRequest) (code int, response responses.ReqresPostUser, err error)
	UpdateUser(context context.Context, index int, request requests.ReqresUserRequest) (code int, response responses.ReqresPostUser, err error)
	DeleteUser(context context.Context, index int) (code int, response responses.BaseResponse, err error)
}

type ReqresUseCase struct{}

func NewReqresUseCase() *ReqresUseCase {
	return &ReqresUseCase{}
}

func (ReqresUseCase) AllUsers(context context.Context) (code int, response responses.AllReqresUsers, err error) {
	code = http.StatusInternalServerError

	var jsonBody struct {
		Data []models.ReqresUser `json:"data"`
	}

	res, _, errs := gorequest.New().Get("https://reqres.in/api/users").EndStruct(&jsonBody)

	if len(errs) != 0 {
		code = http.StatusBadRequest
		err = errs[0]

		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
	}

	code = res.StatusCode
	response.ReqresUsers = jsonBody.Data
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)

	return
}

func (ReqresUseCase) GetUser(context context.Context, index int) (code int, response responses.ReqresUser, err error) {
	code = http.StatusInternalServerError

	var jsonBody struct {
		Data models.ReqresUser `json:"data"`
	}

	res, _, errs := gorequest.New().Get(fmt.Sprintf("https://reqres.in/api/users/%d", index)).EndStruct(&jsonBody)

	if len(errs) != 0 {
		code = http.StatusBadRequest
		err = errs[0]

		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
	}

	code = res.StatusCode
	response.ReqresUser = jsonBody.Data
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)

	return
}

func (ReqresUseCase) InsertUser(context context.Context, request requests.ReqresUserRequest) (code int, response responses.ReqresPostUser, err error) {
	code = http.StatusInternalServerError

	var jsonBody models.ReqresPostUser

	_, _, errs := gorequest.New().Post("https://reqres.in/api/users/").Send(request).EndStruct(&jsonBody)

	if len(errs) != 0 {
		code = http.StatusBadRequest
		err = errs[0]

		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	code = http.StatusOK
	response.ReqresPostUser = jsonBody
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)

	return
}

func (ReqresUseCase) UpdateUser(context context.Context, index int, request requests.ReqresUserRequest) (code int, response responses.ReqresPostUser, err error) {
	code = http.StatusInternalServerError

	var jsonBody models.ReqresPostUser

	_, _, errs := gorequest.New().Put(fmt.Sprintf("https://reqres.in/api/users/%d", index)).Send(request).EndStruct(&jsonBody)

	if len(errs) != 0 {
		code = http.StatusBadRequest
		err = errs[0]

		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	code = http.StatusOK
	response.ReqresPostUser = jsonBody
	response.ReqresPostUser.ID = fmt.Sprintf("%d", index)
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)

	return
}

func (ReqresUseCase) DeleteUser(context context.Context, index int) (code int, response responses.BaseResponse, err error) {
	code = http.StatusInternalServerError

	res, _, errs := gorequest.New().Delete(fmt.Sprintf("https://reqres.in/api/users/%d", index)).End()

	if len(errs) != 0 {
		code = http.StatusBadRequest
		err = errs[0]

		responses.NewBaseResponseStatusCode(code, &response, err)
	}

	code = res.StatusCode
	responses.NewBaseResponseStatusCode(code, &response, err)

	return
}
