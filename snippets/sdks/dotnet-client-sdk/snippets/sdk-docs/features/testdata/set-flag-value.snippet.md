---
id: dotnet-client-sdk/sdk-docs/features/testdata/set-flag-value
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Setting a test data flag to a specific value for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
td.Update(td.Flag("example-flag-key").Variation(false));
```
