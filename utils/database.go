package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/subosito/gotenv"
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
