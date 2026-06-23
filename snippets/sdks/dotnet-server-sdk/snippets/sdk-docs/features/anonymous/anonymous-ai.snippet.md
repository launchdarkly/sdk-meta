---
id: dotnet-server-sdk/sdk-docs/features/anonymous/anonymous-ai
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Anonymous context example for the .NET AI SDK.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var context = Context.Builder("example-context-key")
    .Anonymous(true)
    .Build();
```
