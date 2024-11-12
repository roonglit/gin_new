package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/roonglit/credentials/pkg/credentials"
	"github.com/roonglit/gin_new/config"
	"github.com/roonglit/gin_new/grpc/gapi"
	"github.com/roonglit/gin_new/grpc/pb"
	"github.com/roonglit/gin_new/http/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Read configuration with gin mode
	reader := credentials.NewConfigReader()
	var config config.Config
	if err := reader.Read(gin.Mode(), &config); err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}

	// runHttpServer(config)
	runGrpcServer(config)
}

func runHttpServer(config config.Config) {
	// Start the server
	server, err := api.NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	err = server.Start(config.HttpServerAddress)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func runGrpcServer(config config.Config) {
	server, err := gapi.NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("gRPC server is running at %s", config.GrpcServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start gRPC server")
	}
}
