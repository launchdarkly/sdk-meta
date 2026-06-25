---
id: go-server-sdk/sdk-docs/features/testdata/configure-v6
sdk: go-server-sdk
kind: reference
lang: go
description: Test data source configuration for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/testhelpers/ldtestdata"
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
    "github.com/launchdarkly/go-sdk-common/v3/ldvalue"
)

td := ldtestdata.DataSource()
// You can set any initial flag states here with td.Update

config := ld.Config{
    DataSource: td,
}
client, _ := ld.MakeCustomClient("YOUR_SDK_KEY", config, 0)
```
