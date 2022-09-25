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

type ReqresExternal struct {
	goRequest *gorequest.SuperAgent
}

func NewReqresExternal() *ReqresExternal {
	return &ReqresExternal{
		goRequest: gorequest.New(),
	}
}

func (reqresExternal *ReqresExternal) AllUsers(jsonBody any) (code int, err error) {

	res, _, errs := reqresExternal.goRequest.Get("https://reqres.in/api/users").EndStruct(&jsonBody)

	if len(errs) != 0 {
		err = errs[0]
	}

	code = res.StatusCode
	return
}

func (reqresExternal *ReqresExternal) GetUser(index int, jsonBody any) (code int, err error) {
	res, _, errs := reqresExternal.goRequest.Get(fmt.Sprintf("https://reqres.in/api/users/%d", index)).EndStruct(&jsonBody)

	if len(errs) != 0 {
		err = errs[0]
	}

	code = res.StatusCode
	return
}

func (reqresExternal *ReqresExternal) InsertUser(request requests.ReqresUserRequest, jsonBody any) (code int, err error) {
	res, _, errs := reqresExternal.goRequest.Post("https://reqres.in/api/users/").Send(request).EndStruct(&jsonBody)

	if len(errs) != 0 {
		err = errs[0]
	}

	code = res.StatusCode
	return
}

func (reqresExternal *ReqresExternal) UpdateUser(index int, request requests.ReqresUserRequest, jsonBody any) (code int, err error) {
	res, _, errs := reqresExternal.goRequest.Put(fmt.Sprintf("https://reqres.in/api/users/%d", index)).Send(request).EndStruct(&jsonBody)

	if len(errs) != 0 {
		err = errs[0]
	}

	code = res.StatusCode
	return
}

func (reqresExternal *ReqresExternal) DeleteUser(index int, jsonBody any) (code int, err error) {
	res, _, errs := reqresExternal.goRequest.Delete(fmt.Sprintf("https://reqres.in/api/users/%d", index)).EndStruct(&jsonBody)

	if len(errs) != 0 {
		err = errs[0]
	}

	code = res.StatusCode
	return
}
