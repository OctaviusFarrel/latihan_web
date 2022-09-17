package services

import (
	"context"
	"net/http"

	"octaviusfarrel.dev/latihan_web/repositories/pgsql"
	"octaviusfarrel.dev/latihan_web/requests"
	"octaviusfarrel.dev/latihan_web/responses"
)

type IPlayerUseCase interface {
	AllPlayers(c context.Context) (response responses.AllPlayers, code int, err error)
	GetPlayer(c context.Context, index string) (response responses.Player, code int, err error)
	InsertPlayer(c context.Context, player requests.PlayerRequest) (response responses.Player, code int, err error)
	UpdatePlayer(c context.Context, index string, player requests.PlayerRequest) (response responses.Player, code int, err error)
	DeletePlayer(c context.Context, index string) (response responses.Player, code int, err error)
}

type PlayerUseCase struct {
	playerRepo pgsql.IPlayerRepo
}

func NewPlayerUseCase(playerRepo pgsql.IPlayerRepo) IPlayerUseCase {
	return &PlayerUseCase{}
}

func (playerUseCase *PlayerUseCase) AllPlayers(c context.Context) (response responses.AllPlayers, code int, err error) {
	code = http.StatusInternalServerError

	result, err := playerUseCase.playerRepo.AllPlayers()

	if err != nil {
		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	code = http.StatusOK
	response.Players = result
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
	return
}

func (playerUseCase *PlayerUseCase) GetPlayer(c context.Context, index string) (response responses.Player, code int, err error) {
	code = http.StatusInternalServerError

	result, err := playerUseCase.playerRepo.GetPlayer(index)

	if err != nil {
		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	code = http.StatusOK
	response.Player = result
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
	return
}

func (playerUseCase *PlayerUseCase) InsertPlayer(c context.Context, player requests.PlayerRequest) (response responses.Player, code int, err error) {
	code = http.StatusInternalServerError

	result, err := playerUseCase.playerRepo.InsertPlayer(player)

	if err != nil {
		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	code = http.StatusOK
	response.Player = result
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
	return
}

func (playerUseCase *PlayerUseCase) UpdatePlayer(c context.Context, index string, player requests.PlayerRequest) (response responses.Player, code int, err error) {
	code = http.StatusInternalServerError

	result, err := playerUseCase.playerRepo.UpdatePlayer(index, player)

	if err != nil {
		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	code = http.StatusOK
	response.Player = result
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
	return
}

func (playerUseCase *PlayerUseCase) DeletePlayer(c context.Context, index string) (response responses.Player, code int, err error) {
	code = http.StatusInternalServerError

	result, err := playerUseCase.playerRepo.DeletePlayer(index)

	if err != nil {
		responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
		return
	}

	code = http.StatusOK
	response.Player = result
	responses.NewBaseResponseStatusCode(code, &response.BaseResponse, err)
	return
}
