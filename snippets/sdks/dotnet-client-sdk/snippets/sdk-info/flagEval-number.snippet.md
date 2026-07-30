---
id: dotnet-client-sdk/sdk-info/flagEval-number
sdk: dotnet-client-sdk
kind: flag-eval
lang: csharp
file: dotnet-client-sdk/flagEval-number.txt
description: Flag evaluation example for dotnet-client-sdk (number flag).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
// Evaluate the feature flag.
var flagValue = client.DoubleVariation("featureKey", 0);

if (flagValue == 1)
{

    // TODO: Put your feature here

}
else
{

    // TODO: Put your fallback behavior here

}
```
