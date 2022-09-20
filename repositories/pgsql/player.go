package pgsql

import (
	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/requests"
)

type IPlayerRepo interface {
	AllPlayers() (players []models.PlayerModel, err error)
	GetPlayer(index int) (player models.PlayerModel, err error)
	InsertPlayer(playerRequest requests.PlayerRequest) (player models.PlayerModel, err error)
	UpdatePlayer(index int, playerRequest requests.PlayerRequest) (player models.PlayerModel, err error)
	DeletePlayer(index int) (player models.PlayerModel, err error)
}

type PlayerRepo struct{}

func NewPlayerRepo() *PlayerRepo {
	return &PlayerRepo{}
}

func (playerRepo *PlayerRepo) AllPlayers() (players []models.PlayerModel, err error) {
	err = dbPool.Find(&players).Error

	return
}

func (playerRepo *PlayerRepo) GetPlayer(index int) (player models.PlayerModel, err error) {
	err = dbPool.First(&player, index).Error

	return
}

func (playerRepo *PlayerRepo) InsertPlayer(playerRequest requests.PlayerRequest) (player models.PlayerModel, err error) {
	player = models.PlayerModel{Name: playerRequest.Name, Age: playerRequest.Age}
	err = dbPool.Create(&player).Error

	return
}

func (playerRepo *PlayerRepo) UpdatePlayer(index int, playerRequest requests.PlayerRequest) (player models.PlayerModel, err error) {
	if err = dbPool.First(&player, index).Error; err != nil {
		return
	}

	player.Name = playerRequest.Name
	player.Age = playerRequest.Age
	if err = dbPool.Updates(&player).Error; err != nil {
		return
	}

	return
}

func (playerRepo *PlayerRepo) DeletePlayer(index int) (player models.PlayerModel, err error) {
	if err = dbPool.Delete(&player, index).Error; err != nil {
		return
	}

	return
}
