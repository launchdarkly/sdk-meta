---
id: go-server-sdk/sdk-docs/features/storing-data/redis/redis
sdk: go-server-sdk
kind: reference
lang: go
description: Redis feature store configuration example for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "time"

    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
    ldredis "github.com/launchdarkly/go-server-sdk-redis-redigo"
)

var config ld.Config
config.DataStore = ldcomponents.PersistentDataStore(
    ldredis.DataStore().
        HostAndPort("my-redis", 6379).
        Prefix("my-key-prefix"),
).CacheSeconds(30)
client, _ := ld.MakeCustomClient(sdkKey, config, 5*time.Second)
```
