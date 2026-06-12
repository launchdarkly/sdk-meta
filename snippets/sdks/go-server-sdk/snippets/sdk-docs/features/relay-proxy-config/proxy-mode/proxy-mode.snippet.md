---
id: go-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode
sdk: go-server-sdk
kind: reference
lang: go
description: Proxy mode configuration example for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
// To use the Replay Proxy in proxy mode for both streaming and events:

config := ld.Config{
    ServiceEndpoints: ldcomponents.RelayProxyEndpoints(
      "https://your-relay-proxy.com:8030"),
}

// Alternatively, to use the Relay Proxy in proxy mode for streaming,
// but send events directly to LaunchDarkly, use:
config := ld.Config{
    ServiceEndpoints: ldcomponents.RelayProxyEndpointsWithoutEvents(
        "https://your-relay-proxy.com:8030"),
}
```
