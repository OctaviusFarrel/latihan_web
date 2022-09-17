package pgsql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/subosito/gotenv"
)

var (
	config *pgxpool.Config = func() (config *pgxpool.Config) {
		gotenv.Load()
		config, err := pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOSTNAME"), os.Getenv("PORT_DATABASE"), os.Getenv("DATABASE")))

		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		return
	}()

	dbPool *pgxpool.Pool = func(config *pgxpool.Config) (connection *pgxpool.Pool) {
		gotenv.Load()
		connection, err := pgxpool.ConnectConfig(context.Background(), config)

		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		return
	}(config)
)
