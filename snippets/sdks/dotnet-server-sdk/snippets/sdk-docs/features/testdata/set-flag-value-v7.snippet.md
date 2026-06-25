---
id: dotnet-server-sdk/sdk-docs/features/testdata/set-flag-value-v7
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Setting a test data flag to a specific value for .NET (server-side) SDK v7.0.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
td.Update(td.Flag("example-flag-key").VariationForAll(false));
```
