---
id: go-server-sdk/sdk-docs/features/migrations/read-write
sdk: go-server-sdk
kind: reference
lang: go
description: Migration read and write example for Go SDK v7+ (LDClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
context := ldcontext.New("example-context-key")

// this is the migration stage to use if the flag's migration stage
// is not available from LaunchDarkly
defaultStage := ldmigration.Off

readResult := migrator.Read("example-migration-flag-key", context, defaultStage, nil)

writeResult := migrator.Write("example-migration-flag-key", context, defaultStage, nil)
```
