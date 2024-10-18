# otelx
Custom SpanProcessor and SpanExporter

# usage

````go
import (
    "github.com/karlpokus/otelx"
    "go.opentelemetry.io/otel/sdk/trace"
)

tp := trace.NewTracerProvider(otelx.WithAll())
````
