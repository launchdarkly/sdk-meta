---
id: dotnet-client-sdk/sdk-info/flagEval-json
sdk: dotnet-client-sdk
kind: flag-eval
lang: csharp
file: dotnet-client-sdk/flagEval-json.txt
description: Flag evaluation example for dotnet-client-sdk (JSON flag).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
// Evaluate the feature flag.
var flagValue = client.JsonVariation("featureKey", LdValue.Null);

// Use the flag value to configure your feature
// TODO: Put your feature here
```
