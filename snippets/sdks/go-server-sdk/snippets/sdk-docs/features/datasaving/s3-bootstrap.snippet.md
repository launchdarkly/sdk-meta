---
id: go-server-sdk/sdk-docs/features/datasaving/s3-bootstrap
sdk: go-server-sdk
kind: reference
lang: go
description: Data saving mode with Amazon S3 bootstrap and live updates for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    ld "github.com/launchdarkly/go-server-sdk/v7"
    "github.com/launchdarkly/go-server-sdk/v7/ldcomponents"
    lds3 "github.com/launchdarkly/go-server-sdk-s3"
)

var config ld.Config

config.DataSystem = ldcomponents.DataSystem().Custom().
    Initializers(
        ldcomponents.PollingDataSourceV2().AsInitializer(),
        lds3.DataInitializerForObject("my-bucket", "flags/snapshot.json"),
    ).
    Synchronizers(
        ldcomponents.StreamingDataSourceV2(),
        ldcomponents.PollingDataSourceV2(),
    )

client, _ := ld.MakeCustomClient("YOUR_SDK_KEY", config, 5*time.Second)
```
