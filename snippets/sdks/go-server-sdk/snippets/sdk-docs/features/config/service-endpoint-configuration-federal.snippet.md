---
id: go-server-sdk/sdk-docs/features/config/service-endpoint-configuration-federal
sdk: go-server-sdk
kind: reference
lang: go
description: Service endpoint configuration example for Go.
---

```go

config := ld.Config{
    ServiceEndpoints: interfaces.ServiceEndpoints{
      Streaming: "https://stream.launchdarkly.us",
      Polling: "https://sdk.launchdarkly.us",
      Events: "https://events.launchdarkly.us",
    },
}
```
