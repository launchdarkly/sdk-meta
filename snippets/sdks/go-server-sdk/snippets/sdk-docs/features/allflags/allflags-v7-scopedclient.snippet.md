---
id: go-server-sdk/sdk-docs/features/allflags/allflags-v7-scopedclient
sdk: go-server-sdk
kind: reference
lang: go
description: All flags example for Go SDK v7.13.4+ (LDScopedClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "github.com/launchdarkly/go-server-sdk/v7/interfaces/flagstate"
)

// There is not an AllFlagsState method in the LDScopedClient,
// so you need to access the method from the LDClient.
// Then, pass in the scoped client's current context.
// LDScopedClient is in beta and may change without notice.

state := scopedClient.Client().AllFlagsState(scopedClient.CurrentContext(), flagstate.OptionClientSideOnly())
```
