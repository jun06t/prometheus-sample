package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"code", "method"},
	)
)

func init() {
	prometheus.MustRegister(httpReqs)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	count()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func count() {
	httpReqs.WithLabelValues("404", "POST").Add(42)

	m := httpReqs.WithLabelValues("200", "GET")
	for i := 0; i < 1000000; i++ {
		m.Inc()
	}
}
