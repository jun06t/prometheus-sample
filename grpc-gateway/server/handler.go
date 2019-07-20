package main

import (
	"context"
	"errors"
	"log"
	"math/rand"

	pb "github.com/jun06t/prometheus-sample/grpc-gateway/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type aliveService struct{}

func (s *aliveService) GetStatus(ctx context.Context, in *pb.Empty) (*pb.AliveResponse, error) {
	return &pb.AliveResponse{Status: true}, nil
}

type userService struct{}

func (s *userService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Id:   in.Id,
		Name: "Alice",
		Age:  20,
	}, nil
}

func (s *userService) GetUsersByGroup(ctx context.Context, in *pb.UserGroupRequest) (*pb.UsersResponse, error) {
	err := randomError()
	return &pb.UsersResponse{
		Group: in.Group,
		Users: []*pb.UserResponse{
			{Name: "Alice", Age: 20},
			{Name: "Bob", Age: 24},
		},
	}, err
}

func (s *userService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	log.Printf("update body is {id: %s, name: %s, age: %d}\n", in.Id, in.Name, in.Age)
	return &pb.Empty{}, nil
}

func randomError() error {
	n := rand.Intn(5)
	switch n {
	case 0:
		return nil
	case 1:
		return status.Error(codes.NotFound, "Not found")
	case 2:
		return status.Error(codes.InvalidArgument, "Bad Request")
	case 3:
		return status.Error(codes.Internal, "Internal Server Error")
	case 4:
		return errors.New("Unknown error")
	}
	return nil
}
