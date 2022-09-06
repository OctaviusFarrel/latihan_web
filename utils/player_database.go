package utils

import (
	"context"
	"errors"
	"os"

	. "octaviusfarrel.dev/latihan_web/models"
)

func GetAllPlayers() []Player {
	test := []Player{}

	rows, err := Dbpool.Query(context.Background(), "SELECT * FROM players ORDER BY id")

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	for rows.Next() {
		player := Player{}
		rows.Scan(&player.Id, &player.Name, &player.Age)
		test = append(test, player)
	}

	return test
}

func GetPlayer(index string) (Player, error) {
	rows, err := Dbpool.Query(context.Background(), "SELECT * FROM players WHERE id = $1", index)

	if err != nil {
		return Player{}, err
	}

	player := Player{}
	if rows.Next() {
		rows.Scan(&player.Id, &player.Name, &player.Age)
	} else {
		return player, errors.New("404")
	}

	return player, nil
}

func InsertPlayer(player Player) bool {
	_, err := Dbpool.Query(context.Background(), "INSERT INTO players (name,age) VALUES ($1,$2)", player.Name, player.Age)

	return err == nil
}

func UpdatePlayer(index string, player Player) bool {
	_, err := Dbpool.Query(context.Background(), "UPDATE players SET name = $1, age = $2 WHERE id = $3", player.Name, player.Age, index)

	return err == nil
}

func DeletePlayer(index string) error {
	_, err := Dbpool.Query(context.Background(), "DELETE FROM players WHERE id = $1", index)

	if err != nil {
		return err
	}

	return nil
}
