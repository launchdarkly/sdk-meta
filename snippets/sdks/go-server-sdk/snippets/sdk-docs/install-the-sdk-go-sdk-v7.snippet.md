---
id: go-server-sdk/sdk-docs/install-the-sdk-go-sdk-v7
sdk: go-server-sdk
kind: reference
lang: go
description: "Go SDK v7 in section \"Install the SDK\""
---

```go
import (
    // go-sdk-common/v3/ldcontext defines LaunchDarkly's model for contexts
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"

    // go-sdk-common/v3/ldmigration defines LaunchDarkly's model for migration feature flags
    // (only needed if you are working with migration flags)
    "github.com/launchdarkly/go-sdk-common/v3/ldmigration"

    // go-server-sdk/v7 is the main SDK package - here we are aliasing it to "ld"
    ld "github.com/launchdarkly/go-server-sdk/v7"

    // go-server-sdk/v7/ldcomponents is for advanced configuration options
    "github.com/launchdarkly/go-server-sdk/v7/ldcomponents"

    // go-server-sdk/v7/ldplugins allows you to add plugins to the main SDK
    "github.com/launchdarkly/go-server-sdk/v7/ldplugins"

    // observability-sdk/go is the observability plugin - here we are aliasing it to "ldobserve"
    // this package requires go-server-sdk/v7 version 7.11 or later
    ldobserve "github.com/launchdarkly/observability-sdk/go"
)
```
