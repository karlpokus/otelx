# otelx
Custom otel SDK-, and otel collector components.

````go
import (
    "github.com/karlpokus/otelx"
    "go.opentelemetry.io/otel/sdk/trace"
)

tp := trace.NewTracerProvider(otelx.WithAll())
````

# Components

````sh
# SDK
* trace provider
├── global config
├── default resource
├── sampler (unsampled spans are created in noop mode)
├── id generator
└── * span processor
    ├── handlers for start and end span
    ├── forceFlush calls exporter.ExportSpans immediately
    ├── Shutdown calls exporter.Shutdown
    └── * span exporter
        └── exports sampled spans
````
