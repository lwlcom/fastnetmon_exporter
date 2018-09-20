package network

import (
	"errors"

	"github.com/lwlcom/fastnetmon_exporter/rpc"

	"github.com/lwlcom/fastnetmon_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "fastnetmon_network_"

var (
	incomingPacketsDesc *prometheus.Desc
	outgoingPacketsDesc *prometheus.Desc
	incomingBytesDesc   *prometheus.Desc
	outgoingBytesDesc   *prometheus.Desc
)

func init() {
	l := []string{"target", "network_name"}
	incomingPacketsDesc = prometheus.NewDesc(prefix+"incoming_packets", "Counter for incoming packets", l, nil)
	outgoingPacketsDesc = prometheus.NewDesc(prefix+"outgoing_packets", "Counter for outgoing packets", l, nil)
	incomingBytesDesc = prometheus.NewDesc(prefix+"incoming_bytes", "Counter for incoming bytes", l, nil)
	outgoingBytesDesc = prometheus.NewDesc(prefix+"outgoing_bytes", "Counter for outgoing bytes", l, nil)
}

type networkCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &networkCollector{}
}

// Describe describes the metrics
func (*networkCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- incomingPacketsDesc
	ch <- outgoingPacketsDesc
	ch <- incomingBytesDesc
	ch <- outgoingBytesDesc
}

// Collect collects metrics from FastNetMon API
func (c *networkCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = Response{}
	err := client.RunCommandAndParse("/network_counters", &x)
	if err != nil {
		return err
	}

	if x.Success == false {
		return errors.New(x.ErrorText)
	}

	for _, counter := range x.Values {
		l := append(labelValues, counter.NetworkName)

		ch <- prometheus.MustNewConstMetric(incomingPacketsDesc, prometheus.GaugeValue, float64(counter.IncomingPackets), l...)
		ch <- prometheus.MustNewConstMetric(outgoingPacketsDesc, prometheus.GaugeValue, float64(counter.OutgoingPackets), l...)
		ch <- prometheus.MustNewConstMetric(incomingBytesDesc, prometheus.GaugeValue, float64(counter.IncomingBytes), l...)
		ch <- prometheus.MustNewConstMetric(outgoingBytesDesc, prometheus.GaugeValue, float64(counter.OutgoingBytes), l...)
	}

	return nil
}
