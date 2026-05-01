---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-networking-configuration-5-0-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "5.0 syntax in section \"Understanding changes to networking configuration\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v5"
  "gopkg.in/launchdarkly/go-server-sdk.v5/ldcomponents"
)

var config ld.Config

// 5.0 model: setting socket connection timeout
config.HTTP = ldcomponents.HTTPConfiguration().
    ConnectTimeout(3*time.Second)

// 5.0 model: specifying a secure HTTPS proxy with a custom CA certificate
config.HTTP = ldcomponents.HTTPConfiguration().
    ProxyURL(proxyURL).
    CACertFile("mycert.crt")
```
