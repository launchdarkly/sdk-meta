---
id: go-server-sdk/sdk-docs/features/otel/resource-attributes
sdk: go-server-sdk
kind: reference
lang: go
description: Programmatic OpenTelemetry resource attribute configuration (Go).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
)

res := resource.NewWithAttributes(
    semconv.SchemaURL,
    attribute.String("launchdarkly.project_id", "YOUR_SDK_KEY"),
)

tp := trace.NewTracerProvider(
    trace.WithResource(res),
)
```
