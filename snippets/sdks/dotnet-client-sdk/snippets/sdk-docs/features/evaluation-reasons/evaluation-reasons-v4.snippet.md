---
id: dotnet-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v4
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Flag evaluation reason example for .NET (client-side) v4.0.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only

---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .EvaluationReasons(true)
    .Build();
LdClient client = LdClient.Init(config, context, TimeSpan.FromSeconds(10));

EvaluationDetail<bool> detail =
    client.BoolVariationDetail("example-bool-flag-key", false);
    // or StringVariationDetail for a string-valued flag, and so on.

bool value = detail.Value;
int? index = detail.VariationIndex;
EvaluationReason reason = detail.Reason;
```
