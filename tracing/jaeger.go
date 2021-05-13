package tracing

import (
	"errors"
	"fmt"

	"github.com/concourse/flag"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
)

// Jaeger service to export traces to
type Jaeger struct {
	Endpoint string              `yaml:"endpoint,omitempty"`
	Tags     flag.StringToString `yaml:"tags,omitempty"`
	Service  string              `yaml:"service,omitempty"`
}

func (j Jaeger) ID() string {
	return "jaeger"
}

// Validate identifies if an endpoint has been set
func (j Jaeger) Validate() error {
	if j.Endpoint == "" {
		return errors.New("endpoint is missing")
	}

	return nil
}

// Exporter returns a SpanExporter to sync spans to Jaeger
func (j Jaeger) Exporter() (sdktrace.SpanExporter, []sdktrace.TracerProviderOption, error) {
	attributes := append([]attribute.KeyValue{semconv.ServiceNameKey.String(j.Service)}, keyValueSlice(map[string]string(j.Tags))...)
	exporter, err := jaeger.NewRawExporter(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(j.Endpoint)),
	)
	extraOptions := []sdktrace.TracerProviderOption{
		sdktrace.WithResource(resource.NewWithAttributes(attributes...)),
	}
	if err != nil {
		err = fmt.Errorf("failed to create jaeger exporter: %w", err)
		return nil, nil, err
	}

	return exporter, extraOptions, nil
}
