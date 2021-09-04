package file

import (
	"example/internal/metrics"
	"example/internal/metrics/encoding"
	"os"
)

type Provider struct {
	path    string
	decoder encoding.Decoder
}

func NewProvider(path string, decoder encoding.Decoder) metrics.Provider {
	return &Provider{
		path:    path,
		decoder: decoder,
	}
}

func (p *Provider) GetMetrics() (metrics.Metrics, error) {
	f, err := os.Open(p.path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	return p.decoder.Decode(f)
}
