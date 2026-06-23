---
id: dotnet-client-sdk/sdk-docs/features/anonymous/anonymous-placeholder-key
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Placeholder-key anonymous context example for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
Context context = Context.Builder("placeholder-key")
    .Anonymous(true)
    .Build();
```
