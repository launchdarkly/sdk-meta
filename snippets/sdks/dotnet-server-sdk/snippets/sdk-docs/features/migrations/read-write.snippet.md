---
id: dotnet-server-sdk/sdk-docs/features/migrations/read-write
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Migration read and write example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
Context context = Context.Builder("example-context-key")
    .Build();

// this is the migration stage to use if the flag's migration stage
// is not available from LaunchDarkly
var defaultStage = MigrationStage.Off;

var readResult = migration.Read("example-migration-flag-key", context, defaultStage, payload);

var writeResult = migration.Write("example-migration-flag-key", context, defaultStage, payload);
```
