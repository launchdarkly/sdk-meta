---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-the-data-store-5-0-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "5.0 syntax in section \"Understanding changes to the data store\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v5"
  "gopkg.in/launchdarkly/go-server-sdk.v5/ldcomponents"
  ldredis "github.com/launchdarkly/go-server-sdk-redis-redigo"
)

var config ld.Config

// 5.0 model: use Redis, set custom Redis URI and key prefix, set cache TTL to 45 seconds
config.DataStore = ldcomponents.PersistentDataStore(
    ldredis.DataStore().
        URI("redis://my-redis-host").
        Prefix("my-key-prefix"),
).CacheSeconds(45)
```
