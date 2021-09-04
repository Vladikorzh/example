package yaml

import (
	"example/internal/metrics"
	"gopkg.in/yaml.v2"
	"io"
)

func Decode(from io.Reader) (metrics.Metrics, error) {
	m := make(metrics.Metrics)
	if err := yaml.NewDecoder(from).Decode(m); err != nil {
		return nil, err
	}

	return m, nil
}
