---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-data-source-methods-5-0-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "5.0 syntax in section \"Understanding changes to data source methods\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v5"
  "gopkg.in/launchdarkly/go-server-sdk.v5/ldcomponents"
)

var config ld.Config

// 5.0 model: setting custom options for streaming mode
config.DataSource = ldcomponents.StreamingDataSource().
    InitialReconnectDelay(500*time.Millisecond)

// 5.0 model: specifying polling mode and setting custom polling options
config.DataSource = ldcomponents.PollingDataSource().
    PollInterval(60*time.Second)
```
