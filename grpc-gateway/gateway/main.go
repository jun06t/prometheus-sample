package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/jun06t/prometheus-sample/grpc-gateway/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

var (
	endpoint   = "localhost:9090"
	listenAddr = ":3000"
	promAddr   = ":8081"
)

func init() {
	ep := os.Getenv("ENDPOINT")
	if ep != "" {
		endpoint = ep
	}
}

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
	}

	fmt.Println("Endpoint: ", endpoint)
	conn, err := grpc.Dial(endpoint, dialOpts...)
	if err != nil {
		return nil, err
	}
	err = pb.RegisterAliveServiceHandler(ctx, mux, conn)
	if err != nil {
		return nil, err
	}
	err = pb.RegisterUserServiceHandler(ctx, mux, conn)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

// Run starts a HTTP server and blocks forever if successful.
func Run(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}

	return http.ListenAndServe(address, gw)
}

func main() {
	mux := http.NewServeMux()
	grpc_prometheus.EnableClientHandlingTimeHistogram()
	mux.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(promAddr, mux))
	}()

	fmt.Println("Listen Address:", listenAddr)
	if err := Run(listenAddr); err != nil {
		panic(err)
	}
}
