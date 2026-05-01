---
id: go-server-sdk/sdk-docs/migration-4-to-5-using-the-relay-proxy-4-x-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "4.x syntax in section \"Using the Relay Proxy\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
  ldredis "gopkg.in/launchdarkly/go-server-sdk.v4/redis"
)

config := ld.DefaultConfig
relayURI := "http://my-relay-host:8000"

// 4.x model: proxy mode
config.BaseURI = relayURI
config.StreamURI = relayURI
config.EventsURI = relayURI // if you want to proxy events

// 4.x model: daemon mode with a Redis database
config.UseLdd = true
config.FeatureStoreFactory = ldredis.NewRedisFeatureStoreWithDefaults(
    ldredis.URI("redis://my-redis-host"),
    ldredis.Prefix("my-key-prefix"))
```
