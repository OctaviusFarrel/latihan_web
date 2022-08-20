package main

import (
	"github.com/gin-gonic/gin"

	"octaviusfarrel.dev/gindile/handlers"
	"octaviusfarrel.dev/gindile/utils"
)

func main() {
	router := gin.Default()

	router.GET("/", handlers.GetSomething)

	halo := router.Group("/halo")
	{
		halo.GET("/", handlers.GetSomething)
		halo.GET("/:name", handlers.GetSomethingWithName)
	}

	utils.CreateResource(router, "players").AddStoreRoute(handlers.InsertPlayer).AddIndexRoute(handlers.GetAllPlayers).AddShowRoute(handlers.GetPlayer).AddDestroyRoute(handlers.DeletePlayer).AddUpdateRoute(handlers.UpdatePlayer).Build()

	router.Run()
}
