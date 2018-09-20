package main

import (
	"time"

	"github.com/lwlcom/fastnetmon_exporter/collector"
	"github.com/lwlcom/fastnetmon_exporter/host"
	"github.com/lwlcom/fastnetmon_exporter/network"
	"github.com/lwlcom/fastnetmon_exporter/rpc"
	"github.com/lwlcom/fastnetmon_exporter/totaltraffic"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const prefix = "fastnetmon_"

var (
	scrapeDurationDesc *prometheus.Desc
	upDesc             *prometheus.Desc
)

func init() {
	upDesc = prometheus.NewDesc(prefix+"up", "Scrape of target was successful", []string{"target"}, nil)
	scrapeDurationDesc = prometheus.NewDesc(prefix+"collector_duration_seconds", "Duration of a collector scrape for one target", []string{"target"}, nil)
}

type fastnetmonCollector struct {
	collectors map[string]collector.RPCCollector
}

func newFastnetmonCollector() *fastnetmonCollector {
	collectors := collectors()
	return &fastnetmonCollector{collectors}
}

func collectors() map[string]collector.RPCCollector {
	m := map[string]collector.RPCCollector{}

	if *totalTrafficEnabled == true {
		m["totalTraffic"] = totalTraffic.NewCollector()
	}

	if *networkEnabled == true {
		m["network"] = network.NewCollector()
	}

	if *hostEnabled == true {
		m["host"] = host.NewCollector()
	}

	return m
}

// Describe implements prometheus.Collector interface
func (c *fastnetmonCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- scrapeDurationDesc

	for _, col := range c.collectors {
		col.Describe(ch)
	}
}

// Collect implements prometheus.Collector interface
func (c *fastnetmonCollector) Collect(ch chan<- prometheus.Metric) {
	t := time.Now()
	defer func() {
		ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, time.Since(t).Seconds(), *apiHost)
	}()

	rpc := rpc.NewClient(*apiHost, *apiUsername, *apiPassword, *debug)
	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 1, *apiHost)

	for k, col := range c.collectors {
		err := col.Collect(rpc, ch, []string{*apiHost})
		if err != nil && err.Error() != "EOF" {
			log.Errorln(k + ": " + err.Error())
		}
	}
}
