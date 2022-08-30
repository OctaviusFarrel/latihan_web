package utils

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/subosito/gotenv"
	. "octaviusfarrel.dev/latihan_web/models"
)

var (
	configErr, dbErr error
	config           *pgxpool.Config
	Dbpool           *pgxpool.Pool
)

func init() {
	gotenv.Load()
	config, configErr = pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOSTNAME"), os.Getenv("PORT_DATABASE"), os.Getenv("DATABASE")))
	if configErr != nil {
		println(configErr.Error())
		os.Exit(1)
	}

	Dbpool, dbErr = pgxpool.ConnectConfig(context.Background(), config)
	if dbErr != nil {
		println(dbErr.Error())
		defer Dbpool.Close()
		os.Exit(1)
	}

}

func GetAllData() []Player {
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

func GetOneData(index string) (Player, error) {
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

func PostOneData(player Player) bool {
	_, err := Dbpool.Query(context.Background(), "INSERT INTO players (name,age) VALUES ($1,$2)", player.Name, player.Age)

	if err != nil {
		return false
	}

	return true
}

func PutOneData(index string, player Player) bool {
	_, err := Dbpool.Query(context.Background(), "UPDATE players SET name = $1, age = $2 WHERE id = $3", player.Name, player.Age, index)

	if err != nil {
		return false
	}

	return true
}

func DeleteOneData(index string) error {
	_, err := Dbpool.Query(context.Background(), "DELETE FROM players WHERE id = $1", index)

	if err != nil {
		return err
	}

	return nil
}
