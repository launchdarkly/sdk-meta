---
id: dotnet-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Flag evaluation reason example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
EvaluationDetail<bool> detail =
    client.BoolVariationDetail("example-flag-key", myContext, false);
    // or StringVariationDetail for a string-valued flag, etc.

bool value = detail.Value;
int? index = detail.VariationIndex;
EvaluationReason reason = detail.Reason;
```
