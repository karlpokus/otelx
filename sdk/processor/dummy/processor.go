package dummy

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

func New(exp trace.SpanExporter) *Processor {
	return &Processor{
		e: exp,
	}
}
