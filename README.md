# otelx
Custom otel SDK-, and otel collector components.

````go
import (
    "github.com/karlpokus/otelx"
    "go.opentelemetry.io/otel/sdk/trace"
)

tp := trace.NewTracerProvider(otelx.WithAll())
````
