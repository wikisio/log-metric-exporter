package megric

import "github.com/prometheus/client_golang/prometheus"

type Exporter struct {
}

func NewExporter() *Exporter {
	return nil
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
}
