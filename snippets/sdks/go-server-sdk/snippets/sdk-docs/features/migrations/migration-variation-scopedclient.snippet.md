---
id: go-server-sdk/sdk-docs/features/migrations/migration-variation-scopedclient
sdk: go-server-sdk
kind: reference
lang: go
description: Migration stage evaluation (MigrationVariation) for Go SDK v7.13.4+ (LDScopedClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
context := ldcontext.New("example-context-key")
scopedClient := ld.NewScopedClient(client, context)
// LDScopedClient is in beta and may change without notice.

stage, tracker, err := scopedClient.MigrationVariation("example-migration-flag-key", ldmigration.Off)
```
