---
id: dotnet-server-sdk/sdk-docs/evaluate-a-context-net-sdk-v6-x-c
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: ".NET SDK v6.x (C#) in section \"Evaluate a context\""
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var user = User.Builder("example-user-key")
  .Name("Sandy")
  .Build();

var flagValue = client.BoolVariation("example-flag-key", user, false);

if (flagValue) {
    // application code to show the feature
}
else {
    // the code to run if the feature is off
}
```
