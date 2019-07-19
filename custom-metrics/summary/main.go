package main

import (
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	temps = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "pond_temperature_celsius",
		Help: "The temperature of the frog pond.",
	})
)

func init() {
	prometheus.MustRegister(temps)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)
	simulate()
}

func simulate() {
	var i int
	for {
		i++
		time.Sleep(1 * time.Second)
		val := 30 + math.Floor(120*math.Sin(float64(i)*0.1))/10
		temps.Observe(val)
	}
}
