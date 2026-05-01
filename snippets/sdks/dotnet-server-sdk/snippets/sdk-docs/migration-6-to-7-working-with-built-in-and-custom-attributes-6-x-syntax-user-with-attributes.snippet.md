---
id: dotnet-server-sdk/sdk-docs/migration-6-to-7-working-with-built-in-and-custom-attributes-6-x-syntax-user-with-attributes
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "6.x syntax, user with attributes in section \"Working with built-in and custom attributes\""
---

```csharp
var user = User.Builder("example-user-key")
    .Name("Sandy")
    .Email("sandy@example.com")
    .Custom("groups", LdValue.ArrayOf(LdValue.Of("Acme"), LdValue.Of("Global Health Services")))
    .Build();
```
