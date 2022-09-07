package main

import (
	"github.com/gin-gonic/gin"

	"octaviusfarrel.dev/latihan_web/handlers"
	"octaviusfarrel.dev/latihan_web/middlewares"
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

	app.GET("/", handlers.GetSomething)

	halo := app.Group("/halo")
	{
		halo.GET("/", handlers.GetSomething)
		halo.GET("/:name", handlers.GetSomethingWithName)
	}

	players := app.Group("/players")
	{
		players.Use(middlewares.ReadRequiredTokenMiddleware()).GET("/", handlers.GetAllPlayers)
		players.Use(middlewares.ReadRequiredTokenMiddleware()).GET("/:index", handlers.GetPlayer)
		players.Use(middlewares.WriteRequiredTokenMiddleware()).POST("/", handlers.InsertPlayer)
		players.Use(middlewares.WriteRequiredTokenMiddleware()).PUT("/:index", handlers.UpdatePlayer)
		players.Use(middlewares.WriteRequiredTokenMiddleware()).DELETE("/:index", handlers.DeletePlayer)
	}

	app.POST("/login", handlers.Login)
	app.POST("/register", handlers.Register)

	app.Run()
}
