---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-data-source-methods-4-x-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "4.x syntax in section \"Understanding changes to data source methods\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
)

config := ld.DefaultConfig

// 4.x model: setting custom options for streaming mode
config.Stream = true
config.StreamInitialReconnectDelay = 500*time.Millisecond

// 4.x model: specifying polling mode and setting custom polling options
config.Stream = false
config.PollInterval = 60*time.Second
```
