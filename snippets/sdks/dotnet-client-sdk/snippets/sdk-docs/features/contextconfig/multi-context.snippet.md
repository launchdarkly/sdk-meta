---
id: dotnet-client-sdk/sdk-docs/features/contextconfig/multi-context
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Multi-context example for .NET (client-side) SDK v3.0+.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only

---

```csharp
var userContext = Context.New("example-user-key");

var orgContext = Context.New(ContextKind.Of("organization"), "example-organization-key");

var multiContext = Context.NewMulti(userContext, orgContext);
```
