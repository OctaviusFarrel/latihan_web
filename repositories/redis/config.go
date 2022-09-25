package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/subosito/gotenv"
)

var (
	dbPool = func() *redis.Client {
		gotenv.Load()
		client := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOSTNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})

		_, err := client.Ping().Result()

		if err != nil {
			fmt.Println("cannot connect to redis database")
			os.Exit(1)
		}
		return client
	}()
)
