---
id: go-server-sdk/sdk-docs/features/datasaving/relay-proxy-fallback
sdk: go-server-sdk
kind: reference
lang: go
description: Data saving mode with Relay Proxy and LaunchDarkly API fallback for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    ld "github.com/launchdarkly/go-server-sdk/v7"
    "github.com/launchdarkly/go-server-sdk/v7/ldcomponents"
)

var config ld.Config

relayURI := "http://my-relay-proxy:8030"

config.DataSystem = ldcomponents.DataSystem().Custom().
    Initializers(
        ldcomponents.PollingDataSourceV2().BaseURI(relayURI).AsInitializer(),
        ldcomponents.PollingDataSourceV2().AsInitializer(),
    ).
    Synchronizers(
        ldcomponents.StreamingDataSourceV2().BaseURI(relayURI),
        ldcomponents.StreamingDataSourceV2(),
        ldcomponents.PollingDataSourceV2(),
    )

client, _ := ld.MakeCustomClient("YOUR_SDK_KEY", config, 5*time.Second)
```
