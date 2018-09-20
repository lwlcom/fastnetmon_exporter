package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

const version string = "0.1"

var (
	showVersion   = flag.Bool("version", false, "Print version information.")
	listenAddress = flag.String("web.listen-address", ":9368", "Address on which to expose metrics and web interface.")
	metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	apiHost       = flag.String("api.address", "127.0.0.1:10007", "Address where the FastNetMon API listens")
	apiUsername   = flag.String("api.username", "admin", "Username to access FastNetMon API")
	apiPassword   = flag.String("api.password", "your_password_replace_it", "Password to access FastNetMon API")
	debug         = flag.Bool("debug", false, "Show verbose debug output in log")

	totalTrafficEnabled = flag.Bool("traffic.enabled", true, "Scrape total traffic counters metrics")
	networkEnabled      = flag.Bool("network.enabled", true, "Scrape network counters metrics")
	hostEnabled         = flag.Bool("host.enabled", true, "Scrape host counters metrics")
)

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: fastnetmon_exporter [ ... ]\n\nParameters:")
		fmt.Println()
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *showVersion {
		printVersion()
		os.Exit(0)
	}

	startServer()
}

func printVersion() {
	fmt.Println("fastnetmon_exporter")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("Author(s): Martin Poppen")
	fmt.Println("Metric exporter FastNetMon")
}

func startServer() {
	log.Infof("Starting FastNetMon exporter (Version: %s)\n", version)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>FastNetMon Exporter (Version ` + version + `)</title></head>
			<body>
			<h1>FastNetMon Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			<h2>More information:</h2>
			<p><a href="https://github.com/lwlcom/fastnetmon_exporter">github.com/lwlcom/fastnetmon_exporter</a></p>
			</body>
			</html>`))
	})
	http.HandleFunc(*metricsPath, handleMetricsRequest)

	log.Infof("Listening for %s on %s\n", *metricsPath, *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}

func handleMetricsRequest(w http.ResponseWriter, r *http.Request) {
	reg := prometheus.NewRegistry()

	c := newFastnetmonCollector()
	reg.MustRegister(c)

	promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog:      log.NewErrorLogger(),
		ErrorHandling: promhttp.ContinueOnError}).ServeHTTP(w, r)
}
