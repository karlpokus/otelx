package inmem

import (
	"context"

	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

// mockExporter stores spans in memory and
// allows for span inspection post provider shutdown.
type mockExporter struct {
	e          *tracetest.InMemoryExporter
	isShutdown bool
}

func (e *mockExporter) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
	return e.e.ExportSpans(ctx, spans)
}

func (e *mockExporter) GetSpans() tracetest.SpanStubs {
	return e.e.GetSpans()
}

func (e *mockExporter) Shutdown(ctx context.Context) error {
	e.isShutdown = true
	return nil
}

func New() *mockExporter {
	return &mockExporter{e: tracetest.NewInMemoryExporter()}
}
