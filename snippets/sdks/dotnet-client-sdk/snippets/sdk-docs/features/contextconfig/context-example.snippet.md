---
id: dotnet-client-sdk/sdk-docs/features/contextconfig/context-example
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Context example for .NET (client-side) SDK v3.0+.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only

---

```csharp
Context context = Context.Builder("example-context-key")
    .Set("firstName", "Sandy")
    .Set("lastName", "Smith")
    .Set("email", "sandy@example.com")
    .Set("group", "microsoft")
    .Build();
```
