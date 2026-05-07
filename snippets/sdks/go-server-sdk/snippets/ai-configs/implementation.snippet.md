---
id: go-server-sdk/ai-configs/implementation
sdk: go-server-sdk
kind: implementation
lang: go
file: go-server-sdk/ai-configs/implementation.txt
description: Resolve an AI Config with a fallback for go-server-sdk.
---

```go
fallbackValue := ldai.NewConfig().
  Enable().
  WithModelName("my-default-model").
  WithModelParam("temperature", ldvalue.Float64(0.8)).
  WithMessage("", datamodel.System).
  WithProviderName("my-default-provider").
  Build()

cfg, tracker := aiClient.Config("{{configKey}}", context, fallbackValue, map[string]interface{}{"exampleCustomVariable": "exampleCustomValue"})
```
