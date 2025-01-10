// package dumper is an exporter that dumps spans to stdout and keeps no state
package dumper

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/sdk/trace"
)

type dumper struct{}

func (e dumper) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
	log.Println("Dumping spans")
	for i, span := range spans {
		log.Printf("%d name:%s sid:%s tid:%s status:%v",
			i,
			span.Name(),
			span.SpanContext().SpanID().String(),
			span.SpanContext().TraceID().String(),
			span.Status(),
		)
	}
	return nil
}

func (e dumper) Shutdown(ctx context.Context) error {
	return nil
}

func New() trace.SpanExporter {
	return dumper{}
}
