package collector

import (
	"github.com/lwlcom/fastnetmon_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

// RPCCollector collects metrics from FastNetMon API using rpc.Client
type RPCCollector interface {

	// Describe describes the metrics
	Describe(ch chan<- *prometheus.Desc)

	// Collect collects metrics from FastNetMon API
	Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error
}
