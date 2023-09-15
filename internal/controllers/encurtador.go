package controllers

import (
	"encurtador-url/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostEncurtador(c *gin.Context) {
	url := c.Param("url")
	codUrl, err := models.InsertUrl(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"URL Encurtada": codUrl})
}

func GetIDUrlEncurtada(c *gin.Context) {
	codigo := c.Params.ByName("codigo")
	url, err := models.GetIDUrl(codigo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, url)
}
