package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	aliveChain := genInstrumentChain("alive", alive)
	helloChain := genInstrumentChain("hello", hello)

	http.Handle("/metrics", promhttp.Handler())

	http.Handle("/", aliveChain)
	http.Handle("/hello", helloChain)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
