---
id: dotnet-client-sdk/sdk-docs/features/anonymous/anonymous
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Anonymous context example for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
Context context = Context.Builder("example-context-key")
    .Anonymous(true)
    .Build();
```
