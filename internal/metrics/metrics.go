package metrics

type Metric struct {
	Name  string
	Value float64
}

type Metrics map[string][]*Metric

type Provider interface {
	GetMetrics() (Metrics, error)
}
