---
id: dotnet-server-sdk/sdk-docs/migration-7-to-8-reading-and-writing-during-the-migration-net-sdk-v8-0-c
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: ".NET SDK v8.0 (C#) in section \"Reading and writing during the migration\""
---

```csharp
LDContext context = Context.Builder("example-context-key")
    .Build();

// this is the migration stage to use if the flag's migration stage
// is not available from LaunchDarkly
var defaultStage = MigrationStage.Off;

var readResult = migration.Read("example-migration-flag-key", context, defaultStage, payload);

var writeResult = migration.Write("example-migration-flag-key", context, defaultStage, payload);

```
