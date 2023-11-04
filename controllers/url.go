package controllers

import (
	"encurtador-url/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUrls(c *gin.Context) {
	url, err := models.GetAllUrls()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, url)

}

func GetUrl(c *gin.Context) {
	cod := c.Params.ByName("cod")
	url, err := models.GetUrl(cod)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, url)
}

func PostUrl(c *gin.Context) {
	url := c.Param("teste")
	codUrl, err := models.InsertUrl(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Codigo da URL invalido" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"URL encurtada": codUrl})
}
