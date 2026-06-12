---
id: dotnet-server-sdk/sdk-docs/features/migrations/migration-variation
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Migration stage evaluation (MigrationVariation) for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only-typed
---

```csharp
Context context = Context.Builder("example-context-key")
    .Build();

var (stage, tracker) = client.MigrationVariation("example-migration-flag-key", context, MigrationStage.Off);
```
