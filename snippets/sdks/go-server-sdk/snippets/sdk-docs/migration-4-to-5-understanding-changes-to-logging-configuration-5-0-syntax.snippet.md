---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-logging-configuration-5-0-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "5.0 syntax in section \"Understanding changes to logging configuration\""
---

```go
import (
  "gopkg.in/launchdarkly/go-sdk-common.v2/ldlog"
  ld "gopkg.in/launchdarkly/go-server-sdk.v5"
  "gopkg.in/launchdarkly/go-server-sdk.v5/ldcomponents"
)

var config ld.Config

// 5.0 model: disabling logging
config.Logging = ldcomponents.NoLogging()

// 5.0 model: setting log level to Warn
config.Logging = ldcomponents.Logging().MinLevel(ldlog.Warn)

// 5.0 model: specifying that evaluation errors should be logged
config.Logging = ldcomponents.Logging().LogEvaluationErrors(true)
```
