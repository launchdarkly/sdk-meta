---
id: go-server-sdk/sdk-docs/features/migrations/read-write-scopedclient
sdk: go-server-sdk
kind: reference
lang: go
description: Migration read and write example for Go SDK v7.13.4+ (LDScopedClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
context := ldcontext.New("example-context-key")

// this is the migration stage to use if the flag's migration stage
// is not available from LaunchDarkly
defaultStage := ldmigration.Off

// There is not an AllFlagsState method in the LDScopedClient.
// If you are using scoped clients, pass in the scoped client's current context
// LDScopedClient is in beta and may change without notice.

readResult := migrator.Read("example-migration-flag-key", scopedClient.CurrentContext(), defaultStage, nil)

writeResult := migrator.Write("example-migration-flag-key", scopedClient.CurrentContext(), defaultStage, nil)
```
