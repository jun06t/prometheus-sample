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
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	pb "github.com/jun06t/prometheus-sample/grpc-gateway/proto"
)

var (
	endpoint   = "localhost:9090"
	listenAddr = ":3000"
	promAddr   = ":9100"
)

func init() {
	ep := os.Getenv("ENDPOINT")
	if ep != "" {
		endpoint = ep
	}
	aa := os.Getenv("API_ADDR")
	if aa != "" {
		listenAddr = aa
	}
	pa := os.Getenv("PROMETHEUS_METRICS_ADDR")
	if pa != "" {
		promAddr = pa
	}
}

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
	}

	fmt.Println("Backend endpoint: ", endpoint)
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

// run starts a HTTP server and blocks forever if successful.
func run(address string, opts ...runtime.ServeMuxOption) error {
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
	runPrometheus()

	fmt.Println("Gateway bind address:", listenAddr)
	if err := run(listenAddr); err != nil {
		panic(err)
	}
}

// runPrometheus runs prometheus metrics server. This is non-blocking function.
func runPrometheus() {
	mux := http.NewServeMux()
	// Enable histogram
	grpc_prometheus.EnableClientHandlingTimeHistogram()
	mux.Handle("/metrics", promhttp.Handler())
	go func() {
		fmt.Println("Prometheus metrics bind address:", promAddr)
		log.Fatal(http.ListenAndServe(promAddr, mux))
	}()
}
