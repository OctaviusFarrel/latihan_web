package pgsql

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbPool *gorm.DB = func() (connection *gorm.DB) {
		gotenv.Load()
		dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOSTNAME"), os.Getenv("PORT_DATABASE"), os.Getenv("DATABASE"))
		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		return
	}()
)

func AutoMigrate(model any) {
	dbPool.AutoMigrate(&model)
}
