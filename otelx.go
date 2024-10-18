package otelx

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/sdk/trace"
)

type Processor struct {
	e trace.SpanExporter
}

// span start
func (p *Processor) OnStart(parent context.Context, s trace.ReadWriteSpan) {
	log.Printf("Processor.OnStart called on span: %s", s.Name())
}

// span end
func (p *Processor) OnEnd(s trace.ReadOnlySpan) {
	log.Printf("Processor.OnEnd called on span: %s", s.Name())
	// if p.exporter != nil && s.SpanContext().TraceFlags().IsSampled() {
	err := p.e.ExportSpans(context.Background(), []trace.ReadOnlySpan{s})
	if err != nil {
		log.Printf("Processor.Exporter.ExportSpans err: %v", err)
	}
}

// Shutdown is initiated from the parent (trace provider) and is supposed to
// shutdown the exporter. Lookup trace.simpleSpanProcessor for reference.
func (p *Processor) Shutdown(ctx context.Context) error {
	log.Println("Processor.Shutdown called")
	return p.e.Shutdown(ctx)
}

// noop
func (p *Processor) ForceFlush(ctx context.Context) error {
	log.Println("Processor.ForceFlush called")
	return nil
}

type Exporter struct{}

// noop
func (e *Exporter) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
	log.Printf("Exporter.ExportSpans called on %d spans", len(spans))
	return nil
}

func (e *Exporter) Shutdown(ctx context.Context) error {
	// Note!
	//
	// Honor the context deadline.
	log.Println("Exporter.Shutdown called")
	return nil
}

// WithAll is a tracer provider option that sets our SpanProcessor
// and SpanExporter.
func WithAll() trace.TracerProviderOption {
	return trace.WithSpanProcessor(&Processor{e: &Exporter{}})
}
