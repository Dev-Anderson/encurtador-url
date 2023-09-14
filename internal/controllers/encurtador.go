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

func GetAllUrlEncurtada(c *gin.Context) {
	UrlEncurtada, err := models.GetAllUrl()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": UrlEncurtada})
}
