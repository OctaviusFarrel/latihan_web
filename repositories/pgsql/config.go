package pgsql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbPool *gorm.DB = func() (connection *gorm.DB) {
		gotenv.Load()
		dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOSTNAME"), os.Getenv("PORT_DATABASE"), os.Getenv("DATABASE"))
		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second, // Slow SQL threshold
					LogLevel:                  logger.Info, // Log level
					IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
					Colorful:                  true,        // Disable color
				},
			),
		})

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
