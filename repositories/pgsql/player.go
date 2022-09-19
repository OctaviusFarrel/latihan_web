package pgsql

import (
	"context"
	"fmt"
	"strconv"

	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/requests"
)

type IPlayerRepo interface {
	AllPlayers() (players []models.PlayerModel, err error)
	GetPlayer(index string) (player models.PlayerModel, err error)
	InsertPlayer(playerRequest requests.PlayerRequest) (player models.PlayerModel, err error)
	UpdatePlayer(index string, playerRequest requests.PlayerRequest) (player models.PlayerModel, err error)
	DeletePlayer(index string) (player models.PlayerModel, err error)
}

type PlayerRepo struct{}

func NewPlayerRepo() *PlayerRepo {
	return &PlayerRepo{}
}

func (playerRepo *PlayerRepo) AllPlayers() (players []models.PlayerModel, err error) {

	rows, err := dbPool.Query(context.Background(), "SELECT * FROM players ORDER BY id")

	if err != nil {
		return
	}

	for rows.Next() {
		player := models.PlayerModel{}
		rows.Scan(&player.Id, &player.Name, &player.Age)
		players = append(players, player)
	}

	return
}

func (playerRepo *PlayerRepo) GetPlayer(index string) (player models.PlayerModel, err error) {
	rows, err := dbPool.Query(context.Background(), "SELECT * FROM players WHERE id = $1", index)

	if err != nil {
		return
	}

	if !rows.Next() {
		err = fmt.Errorf("data not found")
		return
	}
	rows.Scan(&player.Id, &player.Name, &player.Age)

	return
}

func (playerRepo *PlayerRepo) InsertPlayer(playerRequest requests.PlayerRequest) (player models.PlayerModel, err error) {
	data, err := dbPool.Query(context.Background(), "SELECT NEXTVAL('users_id_seq')")

	if err != nil {
		return
	}

	if !data.Next() {
		err = fmt.Errorf("cannot insert a new data")
		return
	}

	var i int
	data.Scan(&i)

	_, err = dbPool.Query(context.Background(), "INSERT INTO players (name,age) VALUES ($1,$2)", player.Name, player.Age)

	if err != nil {
		return
	}

	player = models.PlayerModel{Id: i, Name: playerRequest.Name, Age: playerRequest.Age}
	return
}

func (playerRepo *PlayerRepo) UpdatePlayer(index string, playerRequest requests.PlayerRequest) (player models.PlayerModel, err error) {
	_, err = dbPool.Query(context.Background(), "UPDATE players SET name = $1, age = $2 WHERE id = $3", player.Name, player.Age, index)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(index)

	if err != nil {
		return
	}

	player = models.PlayerModel{Id: id, Name: playerRequest.Name, Age: playerRequest.Age}
	return
}

func (playerRepo *PlayerRepo) DeletePlayer(index string) (player models.PlayerModel, err error) {
	rows, err := dbPool.Query(context.Background(), "SELECT * FROM players WHERE id = $1", index)

	if err != nil {
		return
	}

	if !rows.Next() {
		err = fmt.Errorf("data not found")
		return
	}
	rows.Scan(&player.Id, &player.Name, &player.Age)

	_, err = dbPool.Query(context.Background(), "DELETE FROM players WHERE id = $1", index)

	if err != nil {
		return
	}

	return
}
