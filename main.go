package main

import (
	"github.com/gin-gonic/gin"
	"octaviusfarrel.dev/latihan_web/controllers"
	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/repositories/pgsql"
	"octaviusfarrel.dev/latihan_web/services"
)

// var authConfig *oauth2.Config

// func init() {
// 	gotenv.Load()
// 	authConfig = &oauth2.Config{
// 		ClientID:     os.Getenv("CLIENT_ID"),
// 		ClientSecret: os.Getenv("CLIENT_SECRET"),
// 		RedirectURL:  "http://localhost:8080/auth",
// 		Scopes: []string{
// 			"https://www.googleapis.com/auth/userinfo.profile",
// 		},
// 		Endpoint: google.Endpoint,
// 	}
// }

func main() {
	app := gin.New()
	app.Use(gin.Logger())

	userRepo := pgsql.NewUserRepo()
	playerRepo := pgsql.NewPlayerRepo()
	tokenRepo := pgsql.NewTokenRepo()

	userUseCase := services.NewUserUseCase(userRepo, tokenRepo)
	playerUseCase := services.NewPlayerUseCase(playerRepo)
	reqresUseCase := services.NewReqresUseCase()

	userController := controllers.NewUserHandler(userUseCase)
	playerController := controllers.NewPlayerHandler(playerUseCase)
	reqresController := controllers.NewReqresHandler(reqresUseCase)

	// app.GET("/", controllers.GetSomething)

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
