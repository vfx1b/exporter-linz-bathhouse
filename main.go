package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting")
	reg := prometheus.NewRegistry()
	reg.MustRegister(defaultCollector)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Info("Ready")
	http.ListenAndServe(":9442", nil)
}
