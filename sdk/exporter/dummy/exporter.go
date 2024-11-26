package dummy

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/sdk/trace"
)

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

func New() *Exporter {
	return &Exporter{}
}
