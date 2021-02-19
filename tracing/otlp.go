package tracing

import (
	"go.opentelemetry.io/otel/exporters/otlp"
	export "go.opentelemetry.io/otel/sdk/export/trace"
	"google.golang.org/grpc/credentials"
)

// OTLP service to export traces to
type OTLP struct {
	Address string            `yaml:"address"`
	Headers map[string]string `yaml:"header"`
	UseTLS  bool              `yaml:"use_tls"`
}

// IsConfigured identifies if an Address has been set
func (s OTLP) IsConfigured() bool {
	return s.Address != ""
}

func (s OTLP) security() otlp.ExporterOption {
	if s.UseTLS {
		return otlp.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	}

	return otlp.WithInsecure()
}

// Exporter returns a SpanExporter to sync spans to OTLP
func (s OTLP) Exporter() (export.SpanSyncer, error) {
	options := []otlp.ExporterOption{
		otlp.WithAddress(s.Address),
		otlp.WithHeaders(s.Headers),
		s.security(),
	}

	exporter, err := otlp.NewExporter(options...)

	if err != nil {
		return nil, err
	}

	return exporter, nil
}
