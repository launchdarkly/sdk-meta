---
id: dotnet-server-sdk/sdk-docs/openfeature/evaluate-a-context
sdk: dotnet-server-sdk
kind: reference
lang: csharp
file: dotnet-server-sdk/sdk-docs/openfeature/evaluate-a-context.cs
description: ".NET (server-side) OpenFeature provider in section \"Evaluate a context\""
validation:
  scaffold: dotnet-server-sdk/scaffolds/openfeature-csharp-runner
---

```csharp
var flagValue = await client.GetBooleanValueAsync("example-flag-key", false, context);
```
