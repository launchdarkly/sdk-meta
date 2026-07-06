---
id: dotnet-server-sdk/sdk-docs/openfeature/construct-a-context-user
sdk: dotnet-server-sdk
kind: reference
lang: csharp
file: dotnet-server-sdk/sdk-docs/openfeature/construct-a-context-user.cs
description: ".NET (server-side) OpenFeature provider in section \"Construct a context\" (user)"
validation:
  scaffold: dotnet-server-sdk/scaffolds/openfeature-csharp-context-runner
---

```csharp
var context = EvaluationContext.Builder()
  .Set("targetingKey", "example-user-key") // Could also use "key" instead of "targetingKey".
  .Build();
```
