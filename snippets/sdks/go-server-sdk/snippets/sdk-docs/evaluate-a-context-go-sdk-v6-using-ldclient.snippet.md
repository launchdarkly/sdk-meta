---
id: go-server-sdk/sdk-docs/evaluate-a-context-go-sdk-v6-using-ldclient
sdk: go-server-sdk
kind: reference
lang: go
description: "Go SDK v6+, using LDClient in section \"Evaluate a context\""
---

```go
import (
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
)

flagKey := "example-flag-key"
context := ldcontext.NewBuilder("example-context-key").
    Name("Sandy").
    Build()

showFeature, _ := client.BoolVariation(flagKey, context, false)
if showFeature {
    // Application code to show the feature
} else {
    // The code to run if the feature is off
}
```
