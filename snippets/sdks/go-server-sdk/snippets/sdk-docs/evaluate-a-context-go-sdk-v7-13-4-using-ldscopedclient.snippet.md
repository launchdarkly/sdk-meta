---
id: go-server-sdk/sdk-docs/evaluate-a-context-go-sdk-v7-13-4-using-ldscopedclient
sdk: go-server-sdk
kind: reference
lang: go
description: "Go SDK v7.13.4+, using LDScopedClient in section \"Evaluate a context\""
---

```go
import (
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
)

flagKey := "example-flag-key"
context := ldcontext.NewBuilder("example-context-key").
    Name("Sandy").
    Build()

scopedClient := ld.NewScopedClient(client, context)
// LDScopedClient is in beta and may change without notice.

showFeature, _ := scopedClient.BoolVariation(flagKey, false)
if showFeature {
    // Application code to show the feature
} else {
    // The code to run if the feature is off
}
```
