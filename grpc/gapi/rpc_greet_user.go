package gapi

import (
	"context"
	"strconv"

	"github.com/roonglit/gin_new/grpc/pb"
)

func (s *Server) GreetUser(ctx context.Context, req *pb.GreetUserRequest) (*pb.GreetUserResponse, error) {
	name := req.Name
	age := req.Age

	res := &pb.GreetUserResponse{
		Message: "Hello " + name + "! You are " + strconv.Itoa(int(age)) + " years old.",
		User: &pb.User{
			Name: name,
			Age:  age,
		},
	}
	return res, nil
}
