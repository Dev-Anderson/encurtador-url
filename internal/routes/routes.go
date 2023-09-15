package routes

import (
	"encurtador-url/internal/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		encurtador := main.Group("encurtador")
		{
			encurtador.POST("/:url", controllers.PostEncurtador)
			encurtador.GET("/:codigo", controllers.GetIDUrlEncurtada)
		}
	}

	return router
}
