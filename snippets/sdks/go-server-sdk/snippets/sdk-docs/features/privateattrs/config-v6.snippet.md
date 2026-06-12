---
id: go-server-sdk/sdk-docs/features/privateattrs/config-v6
sdk: go-server-sdk
kind: reference
lang: go
description: Private attribute configuration for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
)

var config ld.Config

// Make all attributes private for all contexts
config.Events = ldcomponents.SendEvents().AllAttributesPrivate(true)

// Or, make just the email and address attributes private for all contexts
config.Events = ldcomponents.SendEvents().
    PrivateAttributes("name", "email")
```
