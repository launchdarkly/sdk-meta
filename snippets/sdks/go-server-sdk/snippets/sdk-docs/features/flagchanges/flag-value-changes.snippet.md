---
id: go-server-sdk/sdk-docs/features/flagchanges/flag-value-changes
sdk: go-server-sdk
kind: reference
lang: go
description: Flag value change subscription example for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "log"
    ld "github.com/launchdarkly/go-server-sdk/v7"
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
    "github.com/launchdarkly/go-sdk-common/v3/ldvalue"
)

func logWheneverOneFlagChangesForOneUser(client *ld.LDClient, flagKey string, context ldcontext.Context) {
    updateCh := client.GetFlagTracker().AddFlagValueChangeListener(flagKey, context, ldvalue.Null())
    go func() {
        for event := range updateCh {
            log.Printf("Flag %q for context %q has changed from %s to %s", event.Key,
                context.Key(), event.OldValue, event.NewValue)
        }
    }()
}
```
