package monitoring

type Metric map[string]float64

type Metrics struct {
	metrics Metric
}

func CreateMetrics() Metrics {
	return Metrics{
		metrics: make(Metric),
	}
}

func (m *Metrics) Add(name string, value float64) {
	m.metrics[name] = value
}

func (m *Metrics) All() Metric {
	return m.metrics
}
