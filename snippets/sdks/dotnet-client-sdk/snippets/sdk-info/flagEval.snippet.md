---
id: dotnet-client-sdk/sdk-info/flagEval
sdk: dotnet-client-sdk
kind: flag-eval
lang: csharp
file: dotnet-client-sdk/flagEval.txt
description: Flag evaluation example for dotnet-client-sdk.
---

```csharp
// Evaluate the feature flag.
var flagValue = client.BoolVariation("featureKey", false);

if (flagValue)
{

    // TODO: Put your feature here

}
else
{

    // TODO: Put your fallback behavior here

}
```
