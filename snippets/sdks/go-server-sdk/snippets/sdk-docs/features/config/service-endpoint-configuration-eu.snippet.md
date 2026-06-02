---
id: go-server-sdk/sdk-docs/features/config/service-endpoint-configuration-eu
sdk: go-server-sdk
kind: reference
lang: go
description: Service endpoint configuration example for Go.
---

```go

config := ld.Config{
    ServiceEndpoints: interfaces.ServiceEndpoints{
      Streaming: "https://stream.eu.launchdarkly.com",
      Polling: "https://sdk.eu.launchdarkly.com",
      Events: "https://events.eu.launchdarkly.com",
    },
}
```
