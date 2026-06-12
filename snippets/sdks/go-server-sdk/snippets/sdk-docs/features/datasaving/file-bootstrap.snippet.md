---
id: go-server-sdk/sdk-docs/features/datasaving/file-bootstrap
sdk: go-server-sdk
kind: reference
lang: go
description: Data saving mode with file-based bootstrap and live updates for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    ld "github.com/launchdarkly/go-server-sdk/v7"
    "github.com/launchdarkly/go-server-sdk/v7/ldcomponents"
    "github.com/launchdarkly/go-server-sdk/v7/ldfiledatav2"
)

var config ld.Config

config.DataSystem = ldcomponents.DataSystem().Custom().
    Initializers(
        ldfiledatav2.DataSource().FilePaths("flags.json").AsInitializer(),
        ldcomponents.PollingDataSourceV2().AsInitializer(),
    ).
    Synchronizers(
        ldcomponents.StreamingDataSourceV2(),
        ldcomponents.PollingDataSourceV2(),
    )

client, _ := ld.MakeCustomClient("YOUR_SDK_KEY", config, 5*time.Second)
```
