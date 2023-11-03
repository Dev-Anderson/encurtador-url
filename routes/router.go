package routes

import (
	"encurtador-url/config"

	"github.com/gin-gonic/gin"
)

func Inicialize() {
	router := gin.Default()
	configRoutes(router)
	port := config.LoadEnv().PortServer
	router.Run("localhost:" + port)
}
