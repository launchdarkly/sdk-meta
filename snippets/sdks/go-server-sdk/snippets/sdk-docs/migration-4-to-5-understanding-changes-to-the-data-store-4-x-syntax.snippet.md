---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-the-data-store-4-x-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "4.x syntax in section \"Understanding changes to the data store\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
  ldredis "gopkg.in/launchdarkly/go-server-sdk.v4/redis"
)

config := ld.DefaultConfig

// 4.x model: use Redis, set custom Redis URI and key prefix, set cache TTL to 45 seconds
config.FeatureStoreFactory = redis.NewRedisFeatureStoreWithDefaults(
    ldredis.URI("redis://my-redis-host"),
    ldredis.Prefix("my-key-prefix"),
    ldredis.CacheTTL(45*time.Second))
```
