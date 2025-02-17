package app

import (
	"log"
	"os"

	// "weight-tracker/pkg/api"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if no PORT is set
	}

	err := r.Run("0.0.0.0:" + port)
	// err := r.Run("localhost:" + port)
	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
