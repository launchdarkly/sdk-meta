---
id: dotnet-server-sdk/sdk-docs/features/migrations/stage-switch
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Per-stage migration structure for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
// define the combination of reads and writes from the new and old systems
// that should occur at each migration stage

switch (stage)
    {
        case MigrationStage.Off:
        case MigrationStage.DualWrite:
        case MigrationStage.Shadow:
        case MigrationStage.Live:
        case MigrationStage.RampDown:
        case MigrationStage.Complete:
        default:
            // throw an error
            break;
    }
```
