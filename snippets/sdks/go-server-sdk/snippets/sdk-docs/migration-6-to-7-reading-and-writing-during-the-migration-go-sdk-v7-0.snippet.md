---
id: go-server-sdk/sdk-docs/migration-6-to-7-reading-and-writing-during-the-migration-go-sdk-v7-0
sdk: go-server-sdk
kind: reference
lang: go
description: "Go SDK v7.0 in section \"Reading and writing during the migration\""
---

```go
context := ldcontext.New("example-context-key")

// this is the migration stage to use if the flag's migration stage
// is not available from LaunchDarkly
defaultStage := ldmigration.Off

readResult := migrator.Read("example-migration-flag-key", context, defaultStage, nil)

writeResult := migrator.Write("example-migration-flag-key", context, defaultStage, nil)

```
