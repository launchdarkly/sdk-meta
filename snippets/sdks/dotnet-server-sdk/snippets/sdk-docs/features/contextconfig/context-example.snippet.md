---
id: dotnet-server-sdk/sdk-docs/features/contextconfig/context-example
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Context example for .NET (server-side) SDK v7.0.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
Context context = Context.Builder("example-context-key")
    .Set("firstName", "Sandy")
    .Set("lastName", "Smith")
    .Set("email", "sandy@example.com")
    .Set("groups", LdValue.ArrayOf(LdValue.Of("Acme"), LdValue.Of("Global Health Services")))
    .Build();
```
