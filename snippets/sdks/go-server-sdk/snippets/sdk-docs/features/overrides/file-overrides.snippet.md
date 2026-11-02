---
id: go-server-sdk/sdk-docs/features/overrides/file-overrides
sdk: go-server-sdk
kind: reference
lang: go
description: File-based flag override source configuration example for Go SDK v7.18+.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    "time"

    ld "github.com/launchdarkly/go-server-sdk/v7"
    "github.com/launchdarkly/go-server-sdk/v7/ldcomponents"
    "github.com/launchdarkly/go-server-sdk/v7/ldoverrides"
)

var config ld.Config

config.DataSystem = ldcomponents.DataSystem().Default().
    Overrides(ldoverrides.FileSource().FilePaths("/etc/launchdarkly/overrides.json"))

client, _ := ld.MakeCustomClient("YOUR_SDK_KEY", config, 5*time.Second)
```
