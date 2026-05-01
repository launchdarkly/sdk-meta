---
id: go-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-private-attributes-6-0-syntax-all-attributes-marked-private
sdk: go-server-sdk
kind: reference
lang: go
description: "6.0 syntax, all attributes marked private in section \"Understanding changes to private attributes\""
---

```go
import (
    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
)
var config ld.Config
config.Events = ldcomponents.SendEvents().AllAttributesPrivate(true)
```
