package server

import (
	"encurtador-url/cmd/config"
	"encurtador-url/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   config.LoadEnv().PortServer,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("Server runner in port %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
