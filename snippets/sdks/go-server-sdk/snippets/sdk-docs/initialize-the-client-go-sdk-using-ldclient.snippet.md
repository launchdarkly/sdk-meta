---
id: go-server-sdk/sdk-docs/initialize-the-client-go-sdk-using-ldclient
sdk: go-server-sdk
kind: reference
lang: go
description: "Go SDK, using LDClient in section \"Initialize the client\""
---

```go
  client, _ := ld.MakeCustomClient("YOUR_SDK_KEY",
    ld.Config{
      // optional observability plugin, requires Go SDK v7.11+
      Plugins: []ldplugins.Plugin{
        ldobserve.NewObservabilityPlugin()
      },
    }, 5*time.Second)
```
