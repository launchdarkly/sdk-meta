---
id: dotnet-server-sdk/sdk-docs/evaluate-a-context-net-sdk-v7-0-c-2
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: ".NET SDK v7.0+ (C#) in section \"Evaluate a context\""
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var context = Context.Builder("example-context-key")
  .Kind("device")
  .Name("Android")
  .Build();
```
