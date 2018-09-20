# fastnetmon_exporter
Exporter for metrics from FastNetMon API https://prometheus.io/

# flags
Name     | Description | Default
---------|-------------|---------
version | Print version information. |
web.listen-address | Address on which to expose metrics and web interface. | :9368
web.telemetry-path | Path under which to expose metrics. | /metrics
api.address | Address where the FastNetMon API listens | 127.0.0.1:10007
api.username | Username to access FastNetMon API | admin
api.password | Password to access FastNetMon API | your_password_replace_it
debug | Show verbose debug output | false

# metrics

All metrics are enabled by default. To disable something pass a flag `--<name>.enabled=false`, where `<name>` is the name of the metric.

Name     | Description
---------|------------
traffic | show total_traffic_counters
network | show network_counters
host | show host_counters bytes (incoming|outgoing)

## Install
```bash
go get -u github.com/lwlcom/fastnetmon_exporter
```

## Usage
```bash
./fastnetmon_exporter -api.address="127.0.0.1:10007" -api.username="admin" -api.password="your_password_replace_it"
```

## Third Party Components
This software uses components of the following projects
* Prometheus Go client library (https://github.com/prometheus/client_golang)

## License
(c) Martin Poppen, 2018. Licensed under [MIT](LICENSE) license.

## Prometheus
see https://prometheus.io/
