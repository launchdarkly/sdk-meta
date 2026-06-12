---
id: go-server-sdk/sdk-docs/features/storing-data/index/persistent-store
sdk: go-server-sdk
kind: reference
lang: go
description: Persistent feature store configuration example for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "time"

    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
    examplepackage "github.com/launchdarkly/go-server-sdk-some-example-database"
)

var config ld.Config
config.DataStore = ldcomponents.PersistentDataStore(
    examplepackage.DataStore().SomeStoreOptions(),
)
client, _ := ld.MakeCustomClient(sdkKey, config, 5*time.Second)
```
