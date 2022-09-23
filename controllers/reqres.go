package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"octaviusfarrel.dev/latihan_web/requests"
	"octaviusfarrel.dev/latihan_web/responses"
	"octaviusfarrel.dev/latihan_web/services"
)

type ReqresHandler struct {
	reqresUseCase services.IReqresUseCase
}

func NewReqresHandler(reqresUseCase *services.ReqresUseCase) *ReqresHandler {
	return &ReqresHandler{reqresUseCase: reqresUseCase}
}

func (reqresHandler *ReqresHandler) AllUsers(c *gin.Context) {
	code, response, _ := reqresHandler.reqresUseCase.AllUsers(c.Request.Context())

	c.JSON(code, response)
}

func (reqresHandler *ReqresHandler) GetUser(c *gin.Context) {
	index, err := strconv.Atoi(c.Param("index"))

	if err != nil {
		response := responses.BaseResponse{}

		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &responses.BaseResponse{}, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code, response, _ := reqresHandler.reqresUseCase.GetUser(c.Request.Context(), index)

	c.JSON(code, response)
}

func (reqresHandler *ReqresHandler) InsertUser(c *gin.Context) {

	var userRequest requests.ReqresUserRequest

	err := c.Bind(&userRequest)

	if err != nil {
		response := responses.BaseResponse{}

		responses.NewBaseResponseStatusCode(http.StatusInternalServerError, &responses.BaseResponse{}, err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(userRequest.Name) == 0 || len(userRequest.Job) == 0 {
		response := responses.BaseResponse{}

		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &responses.BaseResponse{}, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code, response, _ := reqresHandler.reqresUseCase.InsertUser(c.Request.Context(), userRequest)

	c.JSON(code, response)
}

func (reqresHandler *ReqresHandler) UpdateUser(c *gin.Context) {
	index, err := strconv.Atoi(c.Param("index"))

	if err != nil {
		response := responses.BaseResponse{}

		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &responses.BaseResponse{}, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var userRequest requests.ReqresUserRequest

	err = c.Bind(&userRequest)

	if err != nil {
		response := responses.BaseResponse{}

		responses.NewBaseResponseStatusCode(http.StatusInternalServerError, &responses.BaseResponse{}, err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(userRequest.Name) == 0 || len(userRequest.Job) == 0 {
		response := responses.BaseResponse{}

		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &responses.BaseResponse{}, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code, response, _ := reqresHandler.reqresUseCase.UpdateUser(c.Request.Context(), index, userRequest)

	c.JSON(code, response)
}

func (reqresHandler *ReqresHandler) DeleteUser(c *gin.Context) {
	index, err := strconv.Atoi(c.Param("index"))

	if err != nil {
		response := responses.BaseResponse{}

		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &responses.BaseResponse{}, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code, response, _ := reqresHandler.reqresUseCase.DeleteUser(c.Request.Context(), index)

	c.JSON(code, response)
}
