---
id: go-server-sdk/sdk-docs/migration-4-to-5-using-the-relay-proxy-5-0-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "5.0 syntax in section \"Using the Relay Proxy\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v5"
  "gopkg.in/launchdarkly/go-server-sdk.v5/ldcomponents"
  ldredis "github.com/launchdarkly/go-server-sdk-redis-redigo"
)

var config ld.Config
relayURI := "http://my-relay-host:8000"

// 5.0 model: proxy mode - this example requires version 5.8.0 or higher
config.ServiceEndpoints = ldcomponents.RelayProxyEndpoints(relayURI)

// 5.0 model: daemon mode with a Redis database
config.DataSource = ldcomponents.ExternalUpdatesOnly()
config.DataStore = ldcomponents.PersistentDataStore(
    ldredis.DataStore().
        URI("redis://my-redis-host").
        Prefix("my-key-prefix"))
```
