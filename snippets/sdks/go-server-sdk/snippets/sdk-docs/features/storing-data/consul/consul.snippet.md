---
id: go-server-sdk/sdk-docs/features/storing-data/consul/consul
sdk: go-server-sdk
kind: reference
lang: go
description: Consul feature store configuration example for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "time"

    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
    ldconsul "github.com/launchdarkly/go-server-sdk-consul"
)

var config ld.Config
config.DataStore = ldcomponents.PersistentDataStore(
    ldconsul.DataStore().
        Address("http://my-consul:8100").
        Prefix("my-key-prefix"),
).CacheSeconds(30)
client, _ := ld.MakeCustomClient(sdkKey, config, 5*time.Second)
```
