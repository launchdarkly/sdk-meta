---
id: dotnet-server-sdk/sdk-docs/features/overrides/override-affected
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Check whether a flag override affected an evaluation in .NET (server-side) SDK v9.0+.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var detail = client.BoolVariationDetail("my-boolean-flag-key", context, false);
if (detail.Reason.OverrideAffected)
{
    Console.WriteLine("An override affected this evaluation");
}
```
