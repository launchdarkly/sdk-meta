---
id: go-server-sdk/sdk-docs/features/webproxy/web-proxy-config
sdk: go-server-sdk
kind: reference
lang: go
description: Programmatic web proxy configuration for the Go SDK.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
)

var config ld.Config
config.HTTP = ldcomponents.HTTPConfiguration().
    ProxyURL("https://my-proxy-host:8080")
```
