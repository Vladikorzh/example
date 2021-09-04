package encoding

import (
	"example/internal/metrics"
	"io"
)

type Decoder interface {
	Decode(from io.Reader) (metrics.Metrics, error)
}

type DecodeFunc func(from io.Reader) (metrics.Metrics, error)

func (fn DecodeFunc) Decode(from io.Reader) (metrics.Metrics, error) {
	return fn(from)
}

type Encoder interface {
	Encode(to io.Writer, m metrics.Metrics) error
}

type EncodeFunc func(to io.Writer, m metrics.Metrics) error

func (fn EncodeFunc) Encode(to io.Writer, m metrics.Metrics) error {
	return fn(to, m)
}
