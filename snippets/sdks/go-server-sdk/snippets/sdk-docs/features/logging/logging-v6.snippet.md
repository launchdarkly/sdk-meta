---
id: go-server-sdk/sdk-docs/features/logging/logging-v6
sdk: go-server-sdk
kind: reference
lang: go
description: Logging destination, level, and disable configuration example for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "log"
    "os"
    ldlog "github.com/launchdarkly/go-sdk-common/v3/ldlog"
    ld "github.com/launchdarkly/go-server-sdk/v6"
    "github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
)

var config ld.Config

loggers := ldlog.NewDefaultLoggers()

// Send output to a file
file, _ := os.Create("app.log")
loggers.SetBaseLogger(log.New(file, "", log.LstdFlags))

config.Logging = ldcomponents.Logging().
    Loggers(loggers).
    MinLevel(ldlog.Warn) // Change minimum level to Warn (Debug and Info are disabled)

// Or, disable logging
config.Logging = ldcomponents.NoLogging()
```
