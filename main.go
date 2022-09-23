package main

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"octaviusfarrel.dev/latihan_web/controllers"
	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/repositories/pgsql"
	"octaviusfarrel.dev/latihan_web/services"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger())

	{
		producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, sarama.NewConfig())

		if err != nil {
			return
		}

		app.Use(func(ctx *gin.Context) {
			message := &sarama.ProducerMessage{Topic: "logs", Value: sarama.StringEncoder(fmt.Sprintf("Handler : %s", ctx.HandlerName()))}
			producer.Input() <- message
		})
	}

	userRepo := pgsql.NewUserRepo()
	playerRepo := pgsql.NewPlayerRepo()
	tokenRepo := pgsql.NewTokenRepo()

	userUseCase := services.NewUserUseCase(userRepo, tokenRepo)
	playerUseCase := services.NewPlayerUseCase(playerRepo)
	reqresUseCase := services.NewReqresUseCase()

	userController := controllers.NewUserHandler(userUseCase)
	playerController := controllers.NewPlayerHandler(playerUseCase)
	reqresController := controllers.NewReqresHandler(reqresUseCase)

	{
		pgsql.AutoMigrate(&models.PlayerModel{})
		pgsql.AutoMigrate(&models.UserModel{})
	}

	halo := app.Group("/halo")
	{
		halo.GET("/", controllers.GetSomething)
		halo.GET("/:name", controllers.GetSomethingWithName)
	}

	reqres := app.Group("/reqres")
	{
		reqres.GET("/", reqresController.AllUsers)
		reqres.GET("/:index", reqresController.GetUser)
		reqres.POST("/", reqresController.InsertUser)
		reqres.PUT("/:index", reqresController.UpdateUser)
		reqres.DELETE("/:index", reqresController.DeleteUser)
	}

	players := app.Group("/players")
	{
		readToken := players.Group("")
		// readToken.Use(middlewares.ReadRequiredTokenMiddleware())
		{
			readToken.GET("/", playerController.AllPlayers)
			readToken.GET("/:index", playerController.GetPlayer)
		}

		writeToken := players.Group("")
		// writeToken.Use(middlewares.WriteRequiredTokenMiddleware())
		{
			writeToken.POST("/", playerController.InsertPlayer)
			writeToken.PUT("/:index", playerController.UpdatePlayer)
			writeToken.DELETE("/:index", playerController.DeletePlayer)
		}
	}

	app.POST("/login", userController.GetUserByUsername)
	app.POST("/register", userController.InsertUser)
	app.Run()
}
