---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-events-5-0-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "5.0 syntax in section \"Understanding changes to events\""
---

```go
import (
  "gopkg.in/launchdarkly/go-sdk-common.v2/lduser"
  ld "gopkg.in/launchdarkly/go-server-sdk.v5"
  "gopkg.in/launchdarkly/go-server-sdk.v5/ldcomponents"
)

var config ld.Config

// 5.0 model: disabling events
config.Events = ldcomponents.NoEvents()

// 5.0 model: customizing event behavior
config.Events = ldcomponents.SendEvents().
    Capacity(20000).
    FlushInterval(10*time.Second).
    PrivateAttributeNames(
        lduser.EmailAttribute,
        lduser.NameAttribute,
        lduser.UserAttribute("myCustomAttribute"),
    )
```
