package openmetrics

import (
	"example/internal/metrics"
	"github.com/golang/protobuf/proto"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"io"
)

func Encode(to io.Writer, metrics metrics.Metrics) error {
	for key, values := range metrics {
		family := dto.MetricFamily{
			Name:   &key,
			Type:   dto.MetricType_GAUGE.Enum(),
			Metric: make([]*dto.Metric, 0, len(values)),
		}
		for _, v := range values {
			family.Metric = append(family.Metric, &dto.Metric{
				Label: []*dto.LabelPair{
					{
						Name:  proto.String("name"),
						Value: proto.String(v.Name),
					},
				},
				Gauge: &dto.Gauge{
					Value: proto.Float64(v.Value),
				},
			})
		}

		if _, err := expfmt.MetricFamilyToOpenMetrics(to, &family); err != nil {
			return err
		}
	}

	return nil
}
