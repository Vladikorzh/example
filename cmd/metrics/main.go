package main

import (
	delivery "example/internal/metrics/delivery/http"
	"example/internal/metrics/encoding"
	"example/internal/metrics/encoding/openmetrics"
	"example/internal/metrics/encoding/yaml"
	provider "example/internal/metrics/provider/file"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	path := flag.String("source", "source.yaml", "source file path")
	port := flag.Uint("port", 8080, "port")

	flag.Parse()

	http.Handle(
		"/metrics",
		delivery.NewHandler(
			provider.NewProvider(*path, encoding.DecodeFunc(yaml.Decode)),
			encoding.EncodeFunc(openmetrics.Encode),
		),
	)

	http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
}
