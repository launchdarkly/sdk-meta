---
id: dotnet-client-sdk/sdk-docs/evaluate-a-flag-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "C# in section \"Evaluate a flag\""
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
bool showFeature = client.BoolVariation("example-flag-key", false);
if (showFeature) {
    // Application code to show the feature
}
else {
    // The code to run if the feature is off
}
```
