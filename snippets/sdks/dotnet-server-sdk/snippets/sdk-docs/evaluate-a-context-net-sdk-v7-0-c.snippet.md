---
id: dotnet-server-sdk/sdk-docs/evaluate-a-context-net-sdk-v7-0-c
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: ".NET SDK v7.0+ (C#) in section \"Evaluate a context\""
---

```csharp
var context = Context.Builder("example-context-key")
  .Name("Sandy")
  .Build();

var flagValue = client.BoolVariation("example-flag-key", context, false);

if (flagValue) {
    // application code to show the feature
}
else {
    // the code to run if the feature is off
}
```
