---
id: go-server-sdk/sdk-docs/features/offlinemode/offline-mode-v7-scopedclient
sdk: go-server-sdk
kind: reference
lang: go
description: Offline mode example for Go SDK v7.13.4+ (LDScopedClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
var config ld.Config
config.Offline = true

client, _ := ld.MakeCustomClient("YOUR_SDK_KEY", config, 5*time.Second)
scopedClient := ld.NewScopedClient(client, context)
// LDScopedClient is in beta and may change without notice.

scopedClient.BoolVariation("example-flag-key", false) // will always return the default value (false)
```
