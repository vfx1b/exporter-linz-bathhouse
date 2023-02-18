package main

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	defaultCollector = NewCollector()
)

type Collector struct {
	metrics map[string]*prometheus.Desc
}

func NewCollector() *Collector {
	c := &Collector{}
	c.metrics = make(map[string]*prometheus.Desc)
	c.metrics["bathhouse_cap"] = prometheus.NewDesc("bathhouse_cap", "bathhouse_cap", []string{"name"}, nil)

	return c
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	log.Info("Fetching metrics")
	m := fetchMetrics()
	if m != nil {
		for _, payload := range *m {
			ch <- prometheus.MustNewConstMetric(c.metrics["bathhouse_cap"], prometheus.GaugeValue, float64(payload.RelativeCurrCapacity), findLocationNameForServerSideId(payload.Id).Name)
		}

	} else {
		log.Error("Error fetching metrics")
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	log.Info("Registering metrics")

	ch <- c.metrics["bathhouse_cap"]
}
