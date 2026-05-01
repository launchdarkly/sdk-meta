---
id: dotnet-client-sdk/sdk-docs/migration-2-to-3-working-with-built-in-and-custom-attributes-2-x-syntax-user-with-attributes
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.x syntax, user with attributes in section \"Working with built-in and custom attributes\""
---

```csharp
var user = User.Builder("example-user-key")
  .Name("Sandy")
  .Email("sandy@example.com")
  .Custom("group", "Global Health Services")
  .Build();
```
