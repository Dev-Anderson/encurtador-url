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
	}

	return router
}
