package external

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
	"octaviusfarrel.dev/latihan_web/requests"
)

type IReqresExternal interface {
	AllUsers(jsonBody any) (code int, err error)
	GetUser(index int, jsonBody any) (code int, err error)
	InsertUser(request requests.ReqresUserRequest, jsonBody any) (code int, err error)
	UpdateUser(index int, request requests.ReqresUserRequest, jsonBody any) (code int, err error)
	DeleteUser(index int, jsonBody any) (code int, err error)
}

type ReqresExternal struct{}

func NewReqresExternal() *ReqresExternal {
	return &ReqresExternal{}
}

func (ReqresExternal) AllUsers(jsonBody any) (code int, err error) {

	res, _, errs := gorequest.New().Get("https://reqres.in/api/users").EndStruct(&jsonBody)

	if len(errs) != 0 {
		err = errs[0]
	}

	code = res.StatusCode
	return
}

func (ReqresExternal) GetUser(index int, jsonBody any) (code int, err error) {
	res, _, errs := gorequest.New().Get(fmt.Sprintf("https://reqres.in/api/users/%d", index)).EndStruct(&jsonBody)

	if len(errs) != 0 {
		err = errs[0]
	}

	code = res.StatusCode
	return
}

func (ReqresExternal) InsertUser(request requests.ReqresUserRequest, jsonBody any) (code int, err error) {
	res, _, errs := gorequest.New().Post("https://reqres.in/api/users/").Send(request).EndStruct(&jsonBody)

	if len(errs) != 0 {
		err = errs[0]
	}

	code = res.StatusCode
	return
}

func (ReqresExternal) UpdateUser(index int, request requests.ReqresUserRequest, jsonBody any) (code int, err error) {
	res, _, errs := gorequest.New().Put(fmt.Sprintf("https://reqres.in/api/users/%d", index)).Send(request).EndStruct(&jsonBody)

	if len(errs) != 0 {
		err = errs[0]
	}

	code = res.StatusCode
	return
}

func (ReqresExternal) DeleteUser(index int, jsonBody any) (code int, err error) {
	res, _, errs := gorequest.New().Delete(fmt.Sprintf("https://reqres.in/api/users/%d", index)).EndStruct(&jsonBody)

	if len(errs) != 0 {
		err = errs[0]
	}

	code = res.StatusCode
	return
}
