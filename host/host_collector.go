package host

import (
	"errors"

	"github.com/lwlcom/fastnetmon_exporter/rpc"

	"github.com/lwlcom/fastnetmon_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "fastnetmon_host_"

var (
	incomingPacketsDesc *prometheus.Desc
	incomingBytesDesc   *prometheus.Desc
	incomingFlowsDesc   *prometheus.Desc
	outgoingPacketsDesc *prometheus.Desc
	outgoingBytesDesc   *prometheus.Desc
	outgoingFlowsDesc   *prometheus.Desc
)

func init() {
	l := []string{"target", "host"}
	incomingPacketsDesc = prometheus.NewDesc(prefix+"incoming_packets", "Counter for incoming packets per host", l, nil)
	incomingBytesDesc = prometheus.NewDesc(prefix+"incoming_bytes", "Counter for incoming bytes per host", l, nil)
	incomingFlowsDesc = prometheus.NewDesc(prefix+"incoming_flows", "Counter for incoming flows per host", l, nil)

	outgoingPacketsDesc = prometheus.NewDesc(prefix+"outgoing_packets", "Counter for outgoing packets per host", l, nil)
	outgoingBytesDesc = prometheus.NewDesc(prefix+"outgoing_bytes", "Counter for outgoing bytes per host", l, nil)
	outgoingFlowsDesc = prometheus.NewDesc(prefix+"outgoing_flows", "Counter for outgoing flows per host", l, nil)
}

type hostCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &hostCollector{}
}

// Describe describes the metrics
func (*hostCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- incomingPacketsDesc
	ch <- incomingBytesDesc
	ch <- incomingFlowsDesc
	ch <- outgoingPacketsDesc
	ch <- outgoingBytesDesc
	ch <- outgoingFlowsDesc
}

// Collect collects metrics from FastNetMon API
func (c *hostCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var inIf = IncomingResponse{}
	err := client.RunCommandAndParse("/host_counters/bytes/incoming", &inIf)
	if err != nil {
		return err
	}

	if inIf.Success == false {
		return errors.New(inIf.ErrorText)
	}

	for _, counter := range inIf.Values {
		l := append(labelValues, counter.Host)

		ch <- prometheus.MustNewConstMetric(incomingPacketsDesc, prometheus.GaugeValue, float64(counter.IncomingPackets), l...)
		ch <- prometheus.MustNewConstMetric(incomingBytesDesc, prometheus.GaugeValue, float64(counter.IncomingBytes), l...)
		ch <- prometheus.MustNewConstMetric(incomingFlowsDesc, prometheus.GaugeValue, float64(counter.IncomingFlows), l...)
	}

	var outIf = OutgoingResponse{}
	err = client.RunCommandAndParse("/host_counters/bytes/outgoing", &outIf)
	if err != nil {
		return err
	}

	if outIf.Success == false {
		return errors.New(outIf.ErrorText)
	}

	for _, counter := range outIf.Values {
		l := append(labelValues, counter.Host)

		ch <- prometheus.MustNewConstMetric(outgoingPacketsDesc, prometheus.GaugeValue, float64(counter.OutgoingPackets), l...)
		ch <- prometheus.MustNewConstMetric(outgoingBytesDesc, prometheus.GaugeValue, float64(counter.OutgoingBytes), l...)
		ch <- prometheus.MustNewConstMetric(outgoingFlowsDesc, prometheus.GaugeValue, float64(counter.OutgoingFlows), l...)
	}

	return nil
}
