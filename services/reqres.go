package services

import (
	"context"
	"fmt"
	"net/http"

	"octaviusfarrel.dev/latihan_web/external"
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

type ReqresUseCase struct {
	reqresExternal external.IReqresExternal
}

func NewReqresUseCase() *ReqresUseCase {
	return &ReqresUseCase{reqresExternal: external.NewReqresExternal()}
}

func (reqresUseCase *ReqresUseCase) AllUsers(context context.Context) (code int, response responses.AllReqresUsers, err error) {
	code = http.StatusInternalServerError

	var jsonBody struct {
		Data []models.ReqresUser `json:"data"`
	}

	code, err = reqresUseCase.reqresExternal.AllUsers(&jsonBody)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	response.ReqresUsers = jsonBody.Data
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)

	return
}

func (reqresUseCase *ReqresUseCase) GetUser(context context.Context, index int) (code int, response responses.ReqresUser, err error) {
	code = http.StatusInternalServerError

	var jsonBody struct {
		Data models.ReqresUser `json:"data"`
	}

	code, err = reqresUseCase.reqresExternal.GetUser(index, &jsonBody)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	response.ReqresUser = jsonBody.Data
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)

	return
}

func (reqresUseCase *ReqresUseCase) InsertUser(context context.Context, request requests.ReqresUserRequest) (code int, response responses.ReqresPostUser, err error) {
	code = http.StatusInternalServerError

	var jsonBody models.ReqresPostUser

	code, err = reqresUseCase.reqresExternal.InsertUser(request, &jsonBody)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	response.ReqresPostUser = jsonBody
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)

	return
}

func (reqresUseCase *ReqresUseCase) UpdateUser(context context.Context, index int, request requests.ReqresUserRequest) (code int, response responses.ReqresPostUser, err error) {
	code = http.StatusInternalServerError

	var jsonBody models.ReqresPostUser

	code, err = reqresUseCase.reqresExternal.UpdateUser(index, request, &jsonBody)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	response.ReqresPostUser = jsonBody
	response.ReqresPostUser.ID = fmt.Sprintf("%d", index)
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)

	return
}

func (reqresUseCase *ReqresUseCase) DeleteUser(context context.Context, index int) (code int, response responses.BaseResponse, err error) {
	code = http.StatusInternalServerError

	code, err = reqresUseCase.reqresExternal.DeleteUser(index, struct{}{})

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &response, err)
		return
	}

	responses.NewBaseResponseStatusCode(code, &response, err)

	return
}
