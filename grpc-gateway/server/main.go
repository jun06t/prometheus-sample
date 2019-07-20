package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	pb "github.com/jun06t/prometheus-sample/grpc-gateway/proto"
)

var (
	endpoint = ":8080"
	promAddr = ":9100"
)

func init() {
	ep := os.Getenv("ENDPOINT")
	if ep != "" {
		endpoint = ep
	}
	pa := os.Getenv("PROMETHEUS_METRICS_ADDR")
	if pa != "" {
		promAddr = pa
	}
}

func main() {
	fmt.Println("Listen gRPC Address:", endpoint)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	// Register your gRPC service implementations.
	pb.RegisterAliveServiceServer(s, new(aliveService))
	pb.RegisterUserServiceServer(s, new(userService))

	// After all your registrations, make sure all of the Prometheus metrics are initialized.
	grpc_prometheus.Register(s)

	runPrometheus()

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

// runPrometheus runs prometheus metrics server. This is non-blocking function.
func runPrometheus() {
	mux := http.NewServeMux()
	// Enable histogram
	grpc_prometheus.EnableHandlingTimeHistogram()
	mux.Handle("/metrics", promhttp.Handler())
	go func() {
		fmt.Println("Prometheus metrics bind address", promAddr)
		log.Fatal(http.ListenAndServe(promAddr, mux))
	}()
}
