package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	app = &App{}
)

func init() {
	bindAddr := ""
	bindPort := ""

	bindAddr, present := os.LookupEnv("EXPORTER_BIND_ADDR")
	if !present {
		bindAddr = "0.0.0.0"
	}

	bindPort, present = os.LookupEnv("EXPORTER_BIND_PORT")
	if !present {
		bindPort = "9442"
	}

	app.BindAddr = bindAddr
	app.BindPort = bindPort
}

func main() {
	log.Info("Starting")
	reg := prometheus.NewRegistry()
	reg.MustRegister(defaultCollector)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Info("Binding to: " + app.BindAddr + ":" + app.BindPort)
	http.ListenAndServe(app.BindAddr+":"+app.BindPort, nil)
}
