---
id: dotnet-client-sdk/sdk-docs/migration-2-to-3-working-with-built-in-and-custom-attributes-3-x-context-with-attributes
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "3.x, context with attributes in section \"Working with built-in and custom attributes\""
---

```csharp
var context = Context.Builder("example-context-key")
  .Name("Sandy")
  .Set("email", "sandy@example.com")
  .Set("group", "microsoft")
  .Build();
```
