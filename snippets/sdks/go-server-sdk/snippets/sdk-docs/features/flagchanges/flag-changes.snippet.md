---
id: go-server-sdk/sdk-docs/features/flagchanges/flag-changes
sdk: go-server-sdk
kind: reference
lang: go
description: Flag change subscription example for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "log"
    ld "github.com/launchdarkly/go-server-sdk/v7"
)

func logWheneverAnyFlagChanges(client *ld.LDClient) {
    updateCh := client.GetFlagTracker().AddFlagChangeListener()
    go func() {
        for event := range updateCh {
            log.Printf("Flag %q has changed", event.Key)
        }
    }()
}
```
