---
id: dotnet-server-sdk/sdk-docs/migration-6-to-7-working-with-built-in-and-custom-attributes-7-0-syntax-context-with-attributes
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "7.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```csharp
var context = Context.Builder("example-context-key")
    .Name("Sandy")
    .Set("email", "sandy@example.com")
    .Set("groups", LdValue.ArrayOf(LdValue.Of("Acme"), LdValue.Of("Global Health Services")))
    .Build();
```
