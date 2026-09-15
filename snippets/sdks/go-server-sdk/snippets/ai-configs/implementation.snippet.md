---
id: go-server-sdk/ai-configs/implementation
sdk: go-server-sdk
kind: implementation
lang: go
file: go-server-sdk/ai-configs/implementation.txt
description: Resolve an AI Config with a fallback for go-server-sdk.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
defaultValue := ldai.NewAICompletionConfigDefault().
  WithEnabled(true).
  WithModelName("my-default-model").
  WithModelParam("temperature", ldvalue.Float64(0.8)).
  WithMessage("", datamodel.System).
  WithProviderName("my-default-provider")

cfg := aiClient.CompletionConfig("{{configKey}}", context, defaultValue, map[string]interface{}{"exampleCustomVariable": "exampleCustomValue"})
tracker := cfg.CreateTracker()
```
