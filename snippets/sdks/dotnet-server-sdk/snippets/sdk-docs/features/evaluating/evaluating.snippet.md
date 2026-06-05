---
id: dotnet-server-sdk/sdk-docs/features/evaluating/evaluating
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Flag evaluation example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
var value = client.BoolVariation("your.feature.key", context, false);
```
