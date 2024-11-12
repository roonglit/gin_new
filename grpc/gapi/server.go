package gapi

import "github.com/roonglit/gin_new/grpc/pb"

type Server struct {
	pb.UnimplementedUserServiceServer
}

// NewServer creates a new gRPC server
func NewServer() (*Server, error) {

	server := &Server{}

	return server, nil
}
