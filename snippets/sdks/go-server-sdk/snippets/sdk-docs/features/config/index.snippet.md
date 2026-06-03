---
id: go-server-sdk/sdk-docs/features/config/index
sdk: go-server-sdk
kind: reference
lang: go
description: SDK configuration example for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "time"

    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
)

var config ld.Config

config.Events = ldcomponents.SendEvents().FlushInterval(10*time.Second)

client, _ := ld.MakeCustomClient("YOUR_SDK_KEY", config, 5*time.Second)
```
