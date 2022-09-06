package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/utils"
)

func GetAllPlayers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"results": utils.GetAllPlayers(),
		"status":  http.StatusOK,
	})
}

func GetPlayer(c *gin.Context) {
	index := c.Param("index")

	player, err := utils.GetPlayer(index)
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
	var formData map[string]interface{}
	c.Bind(&formData)
	if formData["name"] == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "\"name\" is empty",
		})
		return
	}
	if formData["age"] == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "\"age\" is empty",
		})
		return
	}
	if !utils.IsTypeCorrect[string](formData["name"], false) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "\"name\" is not a string",
		})
		return
	}

	if !utils.IsTypeCorrect[float64](formData["age"], false) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "\"age\" is not a number",
		})
		return
	}

	formData["age"] = int(formData["age"].(float64))

	if utils.InsertPlayer(models.Player{Name: formData["name"].(string), Age: int8(formData["age"].(int))}) {
		c.JSON(http.StatusOK, formData)
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
	}

}

func UpdatePlayer(c *gin.Context) {
	index := c.Param("index")

	formData, err := utils.GetPlayer(index)
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

	formUpdate := map[string]interface{}{}
	c.Bind(&formUpdate)

	if formUpdate["name"] != nil {
		if !utils.IsTypeCorrect[string](formUpdate["name"], false) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "\"name\" is not a string",
			})
			return
		}
		formData.Name = formUpdate["name"].(string)
	}

	if formUpdate["age"] != nil {
		if !utils.IsTypeCorrect[float64](formUpdate["age"], false) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "\"age\" is not a number",
			})
			return
		}
		formData.Age = int8(formUpdate["age"].(float64))
	}

	if utils.UpdatePlayer(index, formData) {
		c.JSON(http.StatusOK, formData)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
	}

}

func DeletePlayer(c *gin.Context) {
	index := c.Param("index")

	err := utils.DeletePlayer(index)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Player deleted",
	})

}
