---
id: go-server-sdk/sdk-docs/install-the-sdk-go-sdk-v6
sdk: go-server-sdk
kind: reference
lang: go
description: "Go SDK v6 in section \"Install the SDK\""
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    // go-sdk-common/v3/ldcontext defines LaunchDarkly's model for contexts
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"

    // go-server-sdk/v6 is the main SDK package - here we are aliasing it to "ld"
    ld "github.com/launchdarkly/go-server-sdk/v6"

    // go-server-sdk/v6/ldcomponents is for advanced configuration options
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
)
```
