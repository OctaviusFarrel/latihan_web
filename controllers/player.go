package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"octaviusfarrel.dev/latihan_web/requests"
	"octaviusfarrel.dev/latihan_web/responses"
	"octaviusfarrel.dev/latihan_web/services"
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

	var playerRequest requests.PlayerRequest
	err := c.BindJSON(&playerRequest)

	if err != nil {

		res := responses.BaseResponse{}

		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

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

	playerRequest := requests.PlayerRequest{Name: res.Player.Name, Age: res.Player.Age}
	err = c.BindJSON(&playerRequest)

	if err != nil {

		res := responses.BaseResponse{}

		responses.NewBaseResponseStatusCode(http.StatusBadRequest, &res, err)

		c.JSON(http.StatusBadRequest, res)
		return
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
