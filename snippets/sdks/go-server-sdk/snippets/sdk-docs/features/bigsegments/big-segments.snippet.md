---
id: go-server-sdk/sdk-docs/features/bigsegments/big-segments
sdk: go-server-sdk
kind: reference
lang: go
description: Big segments Redis store configuration example for Go.
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
config.BigSegments = ldcomponents.BigSegments(
    ldredis.BigSegmentStore().
        HostAndPort("your-redis", 6379).
        Prefix("example-client-side-id"),
    ).
    ContextCacheSize(2000).
    ContextCacheTime(30*time.Second)
client, _ := ld.MakeCustomClient(sdkKey, config, 5*time.Second)
```
