---
id: dotnet-client-sdk/sdk-docs/features/contextconfig/context-kind
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Context with a non-user kind for .NET (client-side) SDK v3.0+.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only

---

```csharp
var context = Context.New(ContextKind.Of("organization"), "example-organization-key");
```
