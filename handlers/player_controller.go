package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"okutajager.dev/gindile/models"
	"okutajager.dev/gindile/utils"
)

func GetAllPlayers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
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

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": player,
	})
}

func InsertPlayer(c *gin.Context) {
	var formData models.Player
	if err := c.ShouldBindWith(&formData, binding.JSON); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "\"name\" or \"age\" is required",
		})
		return
	}
	c.Bind(&formData)

	if utils.PostOneData(formData) {
		c.JSON(http.StatusOK, formData)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
	}

}

func UpdatePlayer(c *gin.Context) {
	index := c.Param("index")

	var formData models.Player
	// if err := c.ShouldBindWith(&formData, binding.JSON); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "\"name\" or \"age\" is required",
	// 	})
	// 	return
	// }
	c.Bind(&formData)

	if utils.PutOneData(index, formData) {
		c.JSON(http.StatusOK, formData)
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
