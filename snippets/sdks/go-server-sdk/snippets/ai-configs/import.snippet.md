---
id: go-server-sdk/ai-configs/import
sdk: go-server-sdk
kind: import
lang: go
file: go-server-sdk/ai-configs/import.txt
description: Import block for go-server-sdk AI Configs.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
  "time"

  "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
  "github.com/launchdarkly/go-sdk-common/v3/ldvalue"
  ld "github.com/launchdarkly/go-server-sdk/v7"
  "github.com/launchdarkly/go-server-sdk/ldai"
  "github.com/launchdarkly/go-server-sdk/ldai/datamodel"
)
```
