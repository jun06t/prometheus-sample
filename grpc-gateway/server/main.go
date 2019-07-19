package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	pb "github.com/jun06t/prometheus-sample/grpc-gateway/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

const (
	port = ":9090"
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
	return &pb.UsersResponse{
		Group: in.Group,
		Users: []*pb.UserResponse{
			{Name: "Alice", Age: 20},
			{Name: "Bob", Age: 24},
		},
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	log.Printf("update body is {id: %s, name: %s, age: %d}\n", in.Id, in.Name, in.Age)
	return &pb.Empty{}, nil
}

func main() {
	fmt.Println("Listen gRPC Address:", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	// Register your gRPC service implementations.
	pb.RegisterAliveServiceServer(s, new(aliveService))
	pb.RegisterUserServiceServer(s, new(userService))

	// Enable histogram
	grpc_prometheus.EnableHandlingTimeHistogram()

	// After all your registrations, make sure all of the Prometheus metrics are initialized.
	grpc_prometheus.Register(s)

	// Register Prometheus metrics handler.
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
