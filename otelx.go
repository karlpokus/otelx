package otelx

import (
	exporter "github.com/karlpokus/otelx/sdk/exporter/dummy"
	processor "github.com/karlpokus/otelx/sdk/processor/dummy"
	"go.opentelemetry.io/otel/sdk/trace"
)

// WithAll is a tracer provider option that sets our SpanProcessor
// and SpanExporter.
func WithAll() trace.TracerProviderOption {
	return trace.WithSpanProcessor(processor.New(exporter.New()))
}
