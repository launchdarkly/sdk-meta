---
id: go-server-sdk/sdk-docs/https-proxy-go-sdk
sdk: go-server-sdk
kind: reference
lang: go
description: "Go SDK in section \"HTTPS Proxy\""
---

```go
var config ld.Config
config.HTTP = ldcomponents.HTTPConfiguration().
    ProxyURL("https://web-proxy.domain.com:8080")
```
