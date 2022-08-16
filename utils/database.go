package utils

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	. "okutajager.dev/gindile/models"
)

var (
	configErr, dbErr error
	config           *pgxpool.Config
	dbpool           *pgxpool.Pool
	// config, configErr = pgxpool.ParseConfig(fmt.Sprintf("user=%q host=%q port=%q dbname=%q", "go_role", "localhost", "5432", "go_db"))
)

func init() {
	config, configErr = pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOSTNAME"), os.Getenv("PORT_DATABASE"), os.Getenv("DATABASE")))
	if configErr != nil {
		println(configErr.Error())
		os.Exit(1)
	}

	dbpool, dbErr = pgxpool.ConnectConfig(context.Background(), config)
	if dbErr != nil {
		println(dbErr.Error())
		defer dbpool.Close()
		os.Exit(1)
	}

}

func GetAllData() []Player {
	test := []Player{}

	rows, err := dbpool.Query(context.Background(), "SELECT name,age FROM players")

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	for rows.Next() {
		player := Player{}
		rows.Scan(&player.Name, &player.Age)
		test = append(test, player)
	}

	return test
}

func GetOneData(index string) (Player, error) {
	rows, err := dbpool.Query(context.Background(), "SELECT name,age FROM players WHERE id = $1", index)

	if err != nil {
		return Player{}, err
	}

	player := Player{}
	if rows.Next() {
		rows.Scan(&player.Name, &player.Age)
	} else {
		return player, errors.New("404")
	}

	return player, nil
}

func PostOneData(player Player) bool {
	_, err := dbpool.Query(context.Background(), "INSERT INTO players (name,age) VALUES ($1,$2)", player.Name, player.Age)

	if err != nil {
		return false
	}

	return true
}

func DeleteOneData(index string) (Player, error) {
	rows, err := dbpool.Query(context.Background(), "DELETE FROM players WHERE id = $1", index)

	if err != nil {
		return Player{}, err
	}

	player := Player{Name: "Unknown", Age: 0}
	rows.Next()
	rows.Scan(&player.Name, &player.Age)

	return player, nil
}
