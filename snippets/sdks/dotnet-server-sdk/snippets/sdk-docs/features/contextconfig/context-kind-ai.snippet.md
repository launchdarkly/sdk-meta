---
id: dotnet-server-sdk/sdk-docs/features/contextconfig/context-kind-ai
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Context with a non-user kind for .NET AI SDK.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
var context2 = Context.New(ContextKind.Of("organization"), "example-organization-key");
```
