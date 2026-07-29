---
id: dotnet-client-sdk/sdk-info/flagEval-string
sdk: dotnet-client-sdk
kind: flag-eval
lang: csharp
file: dotnet-client-sdk/flagEval-string.txt
description: Flag evaluation example for dotnet-client-sdk (string flag).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
// Evaluate the feature flag.
var flagValue = client.StringVariation("featureKey", "fallback-value");

if (flagValue == "example-variation-value")
{

    // TODO: Put your feature here

}
else
{

    // TODO: Put your fallback behavior here

}
```
