package main

import (
	"github.com/gin-gonic/gin"

	"okutajager.dev/gindile/handlers"
	"okutajager.dev/gindile/utils"
)

func main() {
	router := gin.Default()

	router.GET("/", handlers.GetSomething)

	halo := router.Group("/halo")
	{
		halo.GET("/", handlers.GetSomething)
		halo.GET("/:name", handlers.GetSomethingWithName)
	}

	utils.CreateResource(router, "players").AddStoreRoute(handlers.InsertPlayer).AddIndexRoute(handlers.GetAllPlayers).AddShowRoute(handlers.GetPlayer).AddDestroyRoute(handlers.DeletePlayer).Build()

	router.Run()
}
