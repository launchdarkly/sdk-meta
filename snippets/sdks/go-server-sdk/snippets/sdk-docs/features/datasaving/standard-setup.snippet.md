---
id: go-server-sdk/sdk-docs/features/datasaving/standard-setup
sdk: go-server-sdk
kind: reference
lang: go
description: Data saving mode standard setup for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
    ld "github.com/launchdarkly/go-server-sdk/v7"
    "github.com/launchdarkly/go-server-sdk/v7/ldcomponents"
)

var config ld.Config

config.DataSystem = ldcomponents.DataSystem().Default()

client, _ := ld.MakeCustomClient("YOUR_SDK_KEY", config, 5*time.Second)
```
