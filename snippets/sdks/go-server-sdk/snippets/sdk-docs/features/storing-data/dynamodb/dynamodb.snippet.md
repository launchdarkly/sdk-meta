---
id: go-server-sdk/sdk-docs/features/storing-data/dynamodb/dynamodb
sdk: go-server-sdk
kind: reference
lang: go
description: DynamoDB feature store configuration example for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "time"

    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
    lddynamodb "github.com/launchdarkly/go-server-sdk-dynamodb"
)

var config ld.Config
config.DataStore = ldcomponents.PersistentDataStore(
    lddynamodb.DataStore("my-table"),
).CacheSeconds(30)
client, _ := ld.MakeCustomClient(sdkKey, config, 5*time.Second)
```
