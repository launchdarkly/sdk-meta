---
id: go-server-sdk/sdk-docs/features/hooks/define-hook
sdk: go-server-sdk
kind: reference
lang: go
description: Hook implementation and configuration for the Go SDK.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only-raw
---

```go
import (
  ld "github.com/launchdarkly/go-server-sdk/v7"
  "github.com/launchdarkly/go-server-sdk/v7/ldhooks"
)

type exampleHook struct {
  ldhooks.Unimplemented
  metadata ldhooks.Metadata
}

func (e exampleHook) Metadata() ldhooks.Metadata {
  return e.metadata
}

// Implement at least one of `BeforeEvaluation`, `AfterEvaluation`

// `BeforeEvaluation` is called during the execution of a variation method
// before the flag value has been determined

// `AfterEvaluation` is called during the execution of a variation method
// after the flag value has been determined

func newExampleHook() exampleHook {}

var config ld.Config

client, _ = ld.MakeCustomClient("YOUR_SDK_KEY",
  ld.Config{
    Hooks: []ldhooks.Hook{newExampleHook()},
  }, 5*time.Second)
```
