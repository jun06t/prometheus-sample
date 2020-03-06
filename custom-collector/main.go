package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(newCollector())

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type collector struct {
	goroutinesDesc *prometheus.Desc
	threadsDesc    *prometheus.Desc
}

func newCollector() *collector {
	return &collector{
		goroutinesDesc: prometheus.NewDesc(
			"goroutines",
			"Number of goroutines that currently exist.",
			nil, nil),
		threadsDesc: prometheus.NewDesc(
			"threads",
			"Number of OS threads created.",
			nil, nil),
	}
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.goroutinesDesc
	ch <- c.threadsDesc
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.goroutinesDesc, prometheus.GaugeValue, float64(runtime.NumGoroutine()))
	n, _ := runtime.ThreadCreateProfile(nil)
	ch <- prometheus.MustNewConstMetric(c.threadsDesc, prometheus.GaugeValue, float64(n))
}
