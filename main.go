package main

import (
	"github.com/gin-gonic/gin"
	"octaviusfarrel.dev/latihan_web/controllers"
	token_util "octaviusfarrel.dev/latihan_web/lib/token"
	"octaviusfarrel.dev/latihan_web/middlewares"
	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/repositories/pgsql"
	"octaviusfarrel.dev/latihan_web/services"
)

func main() {
	app := gin.New()

	// {
	// 	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, sarama.NewConfig())

	// 	if err != nil {
	// 		return
	// 	}

	// 	app.Use(func(ctx *gin.Context) {
	// 		log := logger_local.NewLogger()
	// 		log.Log(ctx.HandlerName(), 0)

	// 		message := &sarama.ProducerMessage{Topic: "logs", Value: sarama.StringEncoder(log.String(ctx.HandlerName(), 0))}
	// 		producer.Input() <- message
	// 	})
	// }

	app.Use(middlewares.NewLoggerMiddleware().GetLogger())
	userRepo := pgsql.NewUserRepo()
	playerRepo := pgsql.NewPlayerRepo()
	tokenUtil := token_util.NewTokenUtil()

	userUseCase := services.NewUserUseCase(userRepo, tokenUtil)
	playerUseCase := services.NewPlayerUseCase(playerRepo)
	reqresUseCase := services.NewReqresUseCase()

	userController := controllers.NewUserHandler(userUseCase)
	playerController := controllers.NewPlayerHandler(playerUseCase)
	reqresController := controllers.NewReqresHandler(reqresUseCase)

	{
		pgsql.AutoMigrate(&models.PlayerModel{})
		pgsql.AutoMigrate(&models.UserModel{})
		pgsql.AutoMigrate(&models.TokenModel{})
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
		tokenMiddleware := middlewares.NewTokenMiddleware(tokenUtil)
		readToken := players.Group("")
		readToken.Use(tokenMiddleware.ReadToken())
		{
			readToken.GET("/", playerController.AllPlayers)
			readToken.GET("/:index", playerController.GetPlayer)
		}

		writeToken := players.Group("")
		writeToken.Use(tokenMiddleware.WriteToken())
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
