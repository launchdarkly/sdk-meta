---
id: go-server-sdk/sdk-docs/features/offlinemode/offline-mode-v6
sdk: go-server-sdk
kind: reference
lang: go
description: Offline mode example for Go SDK v6.x+ (LDClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
var config ld.Config
config.Offline = true

client, _ := ld.MakeCustomClient("YOUR_SDK_KEY", config, 5*time.Second)
client.BoolVariation("any.feature.flag", context, false) // will always return the default value (false)
```
