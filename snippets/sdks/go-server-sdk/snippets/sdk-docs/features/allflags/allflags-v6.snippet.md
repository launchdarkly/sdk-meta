---
id: go-server-sdk/sdk-docs/features/allflags/allflags-v6
sdk: go-server-sdk
kind: reference
lang: go
description: All flags example for Go SDK v6+ (LDClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "github.com/launchdarkly/go-server-sdk/v6/interfaces/flagstate"
)

state := client.AllFlagsState(context, flagstate.OptionClientSideOnly())
```
