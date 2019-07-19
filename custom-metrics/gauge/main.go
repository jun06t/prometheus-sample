package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsQueued = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "ops_queued",
		Help:      "Number of blob storage operations waiting to be processed.",
	})
)

func init() {
	prometheus.MustRegister(opsQueued)

}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	snapshot()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func snapshot() {
	opsQueued.Set(50)

	opsQueued.Add(10)

	opsQueued.Dec()
	opsQueued.Dec()
}
