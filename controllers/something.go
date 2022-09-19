package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSomething(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pesan": "halo",
	})
}

func GetSomethingWithName(c *gin.Context) {
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"pesan": "halo " + name,
	})
}
