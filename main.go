package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/roonglit/credentials/pkg/credentials"
	"github.com/roonglit/gin_new/api"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func main() {
	reader := credentials.NewConfigReader()

	var config Config
	// Read configuration with mode "debug" or "production"
	if err := reader.Read(gin.Mode(), &config); err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}

	fmt.Printf("Loaded Configuration: %+v\n", config)

	// Start the server
	server := api.NewServer()

	err := server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
