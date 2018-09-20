package totalTraffic

import (
	"errors"

	"github.com/lwlcom/fastnetmon_exporter/rpc"

	"github.com/lwlcom/fastnetmon_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "fastnetmon_total_traffic_"

var (
	trafficDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "counter_name", "unit"}
	trafficDesc = prometheus.NewDesc(prefix+"counter", "Traffic counter", l, nil)
}

type totalTrafficCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &totalTrafficCollector{}
}

// Describe describes the metrics
func (*totalTrafficCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- trafficDesc
}

// Collect collects metrics from FastNetMon API
func (c *totalTrafficCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = Response{}
	err := client.RunCommandAndParse("/total_traffic_counters", &x)
	if err != nil {
		return err
	}

	if x.Success == false {
		return errors.New(x.ErrorText)
	}

	for _, counter := range x.Values {
		l := append(labelValues, counter.CounterName, counter.Unit)

		ch <- prometheus.MustNewConstMetric(trafficDesc, prometheus.GaugeValue, float64(counter.Value), l...)
	}

	return nil
}
