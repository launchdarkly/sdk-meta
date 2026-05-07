---
id: go-server-sdk/sdk-docs/use-go-contexts-with-ldscopedclient-retrieve-ldscopedclient-from-go-context
sdk: go-server-sdk
kind: reference
lang: go
description: "Retrieve LDScopedClient from Go context in section \"Use Go contexts with LDScopedClient\""
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
  // GetScopedClient returns the scoped client from the Go context
  // If a scoped client is not present, it returns nil and false

  func logicWithFeatureFlag(ctx context.Context) {
    scopedClient, ok := ld.GetScopedClient(ctx)
    isFeatureEnabled := false // default value if scoped client is not available
    if ok {
      isFeatureEnabled, err = scopedClient.BoolVariation("example-flag-key", false)
      // handle err as appropriate...
    }
  }

  // MustGetScopedClient also returns the scoped client from the Go context
  // If a scoped client is not present, it panics

  func logicWithFeatureFlag(ctx context.Context) {
    scopedClient := ld.MustGetScopedClient(ctx)
    isFeatureEnabled, err := scopedClient.BoolVariation("example-flag-key", false)
      // handle err as appropriate...
  }

```
