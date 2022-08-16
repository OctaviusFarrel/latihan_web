package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"okutajager.dev/gindile/models"
	"okutajager.dev/gindile/utils"
)

func GetAllPlayers(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"results": utils.GetAllData(),
		"status":  http.StatusOK,
	})
}

func GetPlayer(c *gin.Context) {
	index := c.Param("index")

	player, err := utils.GetOneData(index)
	if err != nil {
		if strings.Compare(err.Error(), "404") == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "data not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"result": player,
	})
}

func InsertPlayer(c *gin.Context) {
	player := models.Player{}
	c.Bind(&player)

	if utils.PostOneData(player) {
		c.JSON(http.StatusOK, player)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
	}

}

func UpdatePlayer(c *gin.Context) {
	player := models.Player{}
	c.Bind(&player)

	if utils.PostOneData(player) {
		c.JSON(http.StatusOK, player)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
	}

}

func DeletePlayer(c *gin.Context) {
	index := c.Param("index")

	player, err := utils.DeleteOneData(index)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, player)

}
