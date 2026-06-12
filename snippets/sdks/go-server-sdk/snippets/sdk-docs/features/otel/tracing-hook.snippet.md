---
id: go-server-sdk/sdk-docs/features/otel/tracing-hook
sdk: go-server-sdk
kind: reference
lang: go
description: OpenTelemetry tracing hook configuration for the Go SDK.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
  ld "github.com/launchdarkly/go-server-sdk/v7"
  "github.com/launchdarkly/go-server-sdk/v7/ldhooks"
  "github.com/launchdarkly/go-server-sdk/ldotel"
)

var config ld.Config

client, _ = ld.MakeCustomClient("YOUR_SDK_KEY",
  ld.Config{
    Hooks: []ldhooks.Hook{ldotel.NewTracingHook()},
  }, 5*time.Second)

```
