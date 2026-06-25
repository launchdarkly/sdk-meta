---
id: go-server-sdk/sdk-docs/features/filedata/flags-from-files
sdk: go-server-sdk
kind: reference
lang: go
description: File data source configuration example for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "time"

    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
    "github.com/launchdarkly/go-server-sdk/v6/ldfiledata"
    "github.com/launchdarkly/go-server-sdk/v6/ldfilewatch"
)

var config ld.Config
config.DataSource = ldfiledata.DataSource().
    FilePaths("file1.json", "file2.json").
    Reloader(ldfilewatch.WatchFiles)
config.Events = ldcomponents.NoEvents()

client, _ := ld.MakeCustomClient("sdk key", config, 5*time.Second)
```
