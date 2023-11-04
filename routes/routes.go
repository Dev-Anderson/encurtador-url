package routes

import (
	"encurtador-url/controllers"

	"github.com/gin-gonic/gin"
)

func configRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		url := main.Group("url")
		{
			url.GET("/", controllers.GetAllUrls)
			url.GET("/:cod", controllers.GetUrl)
			url.POST("/", controllers.PostUrl)
		}
	}

	return router
}
