---
id: go-server-sdk/sdk-docs/features/migrations/migration-variation
sdk: go-server-sdk
kind: reference
lang: go
description: Migration stage evaluation (MigrationVariation) for Go SDK v7+ (LDClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
context := ldcontext.New("example-context-key")
stage, tracker, err := client.MigrationVariation("example-migration-flag-key", context, ldmigration.Off)
```
