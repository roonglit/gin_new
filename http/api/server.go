package api

import (
	"github.com/gin-gonic/gin"
)

// Server represents the API server
type Server struct {
	router *gin.Engine
}

// NewServer creates a new server and setup routing
func NewServer() (*Server, error) {
	router := gin.Default()

	server := &Server{
		router: router,
	}

	server.SetUpRoutes()

	return server, nil
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
