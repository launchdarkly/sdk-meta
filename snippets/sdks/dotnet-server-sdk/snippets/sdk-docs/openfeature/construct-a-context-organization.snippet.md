---
id: dotnet-server-sdk/sdk-docs/openfeature/construct-a-context-organization
sdk: dotnet-server-sdk
kind: reference
lang: csharp
file: dotnet-server-sdk/sdk-docs/openfeature/construct-a-context-organization.cs
description: ".NET (server-side) OpenFeature provider in section \"Construct a context\" (organization)"
validation:
  scaffold: dotnet-server-sdk/scaffolds/openfeature-csharp-context-runner
---

```csharp
var context = EvaluationContext.Builder()
  .Set("kind", "organization")
  .Set("targetingKey", "example-organization-key") // Could also use "key" instead of "targetingKey".
  .Build();
```
