---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-networking-configuration-4-x-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "4.x syntax in section \"Understanding changes to networking configuration\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
  "gopkg.in/launchdarkly/go-server-sdk.v4/ldhttp"
)

config := ld.DefaultConfig

// 4.x model: setting socket connection timeout
config.Timeout = 3*time.Second

// 4.x model: specifying a secure HTTPS proxy with a custom CA certificate
config.HTTPClientFactory = ld.NewHTTPClientFactory(
    ldhttp.ProxyOption(proxyURL),
    ldhttp.CACertOption("mycert.crt"),
)
```
