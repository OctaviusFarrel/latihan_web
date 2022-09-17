package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"octaviusfarrel.dev/latihan_web/requests"
	"octaviusfarrel.dev/latihan_web/responses"
	"octaviusfarrel.dev/latihan_web/services"
	"octaviusfarrel.dev/latihan_web/utils"
)

type PlayerHandler struct {
	playerUseCase services.IPlayerUseCase
}

func NewPlayerHandler(playerUseCase services.IPlayerUseCase) *PlayerHandler {
	return &PlayerHandler{playerUseCase}
}

func (playerHandler *PlayerHandler) AllPlayers(c *gin.Context) {
	ctx := c.Request.Context()

	res, code, err := playerHandler.playerUseCase.AllPlayers(ctx)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &res.BaseResponse, err)

		c.JSON(code, res)
		return
	}

	c.JSON(code, res)
}

func (playerHandler *PlayerHandler) GetPlayer(c *gin.Context) {
	ctx := c.Request.Context()
	index := c.Param("index")

	res, code, err := playerHandler.playerUseCase.GetPlayer(ctx, index)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &res.BaseResponse, err)

		c.JSON(code, res)
		return
	}

	c.JSON(code, res)
}

func (playerHandler *PlayerHandler) InsertPlayer(c *gin.Context) {
	ctx := c.Request.Context()

	var formData map[string]interface{}
	c.Bind(&formData)
	if formData["name"] == nil {
		res := responses.BaseResponse{}
		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, fmt.Errorf("name is undefined"))

		c.JSON(http.StatusBadRequest, res)
		return
	}
	if formData["age"] == nil {
		res := responses.BaseResponse{}
		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, fmt.Errorf("age is undefined"))

		c.JSON(http.StatusBadRequest, res)
		return
	}
	if !utils.IsTypeCorrect[string](formData["name"], false) {
		res := responses.BaseResponse{}
		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, fmt.Errorf("name is not a string"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if !utils.IsTypeCorrect[float64](formData["age"], false) {
		res := responses.BaseResponse{}
		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, fmt.Errorf("age is not a number"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formData["age"] = int(formData["age"].(float64))

	playerRequest := requests.PlayerRequest{Name: formData["name"].(string), Age: int8(formData["age"].(int))}

	res, code, err := playerHandler.playerUseCase.InsertPlayer(ctx, playerRequest)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &res.BaseResponse, err)

		c.JSON(code, res)
		return
	}

	c.JSON(code, res)
}

func (playerHandler *PlayerHandler) UpdatePlayer(c *gin.Context) {
	ctx := c.Request.Context()
	index := c.Param("index")

	res, code, err := playerHandler.playerUseCase.GetPlayer(ctx, index)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &res.BaseResponse, err)

		c.JSON(code, res)
		return
	}

	var formData map[string]interface{}
	c.Bind(&formData)
	if formData["name"] == nil {
		res := responses.BaseResponse{}
		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, fmt.Errorf("name is undefined"))

		c.JSON(http.StatusBadRequest, res)
		return
	}
	if formData["age"] == nil {
		res := responses.BaseResponse{}
		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, fmt.Errorf("age is undefined"))

		c.JSON(http.StatusBadRequest, res)
		return
	}
	if !utils.IsTypeCorrect[string](formData["name"], false) {
		res := responses.BaseResponse{}
		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, fmt.Errorf("name is not a string"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if !utils.IsTypeCorrect[float64](formData["age"], false) {
		res := responses.BaseResponse{}
		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, fmt.Errorf("age is not a number"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formData["age"] = int(formData["age"].(float64))

	playerRequest := requests.PlayerRequest{Name: res.Player.Name, Age: res.Player.Age}
	{
		playerRequest.Name = formData["name"].(string)
		playerRequest.Age = int8(formData["age"].(int))
	}

	res, code, err = playerHandler.playerUseCase.UpdatePlayer(ctx, index, playerRequest)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &res.BaseResponse, err)

		c.JSON(code, res)
		return
	}

	c.JSON(code, res)
}

func (playerHandler *PlayerHandler) DeletePlayer(c *gin.Context) {
	ctx := c.Request.Context()
	index := c.Param("index")

	res, code, err := playerHandler.playerUseCase.DeletePlayer(ctx, index)

	if err != nil {

		responses.NewBaseResponseStatusCode(code, &res.BaseResponse, err)

		c.JSON(code, res)
		return
	}

	c.JSON(code, res)
}
