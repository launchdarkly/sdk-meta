---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-logging-configuration-4-x-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "4.x syntax in section \"Understanding changes to logging configuration\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
  "gopkg.in/launchdarkly/go-server-sdk.v4/ldlog"
)

config := ld.DefaultConfig

// 4.x model: disabling logging
config.Loggers = ldlog.NewDisabledLoggers()

// 4.x model: setting log level to Warn
config.Loggers.SetMinLevel(ldlog.Warn)

// 4.x model: specifying that evaluation errors should be logged
config.LogEvaluationErrors = true
```
