---
id: dotnet-server-sdk/sdk-docs/features/anonymous/anonymous
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Anonymous context example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var context = Context.Builder("example-context-key")
    .Anonymous(true)
    .Build();
```
