---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-events-4-x-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "4.x syntax in section \"Understanding changes to events\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
)

config := ld.DefaultConfig

// 4.x model: disabling events
config.SendEvents = false

// 4.x model: customizing event behavior
config.Capacity = 20000
config.FlushInterval = 10*time.Second
config.PrivateAttributeNames = []string{
    "email",
    "name",
    "myCustomAttribute",
}
```
