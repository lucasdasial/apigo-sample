package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/luccasalves/apigo-sample/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func CreateServer() Server {
	return Server{
		port:   "9000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("Server is running on http://localhost:" + s.port)
	log.Fatal(router.Run(":" + s.port))
}
