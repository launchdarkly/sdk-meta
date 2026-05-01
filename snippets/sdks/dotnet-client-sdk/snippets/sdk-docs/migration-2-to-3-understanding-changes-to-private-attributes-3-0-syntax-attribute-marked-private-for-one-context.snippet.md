---
id: dotnet-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-private-attributes-3-0-syntax-attribute-marked-private-for-one-context
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "3.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```csharp
var context = Context.Builder("context-key-123-abc")
    .Set("email", "sandy@example.com")
    .Private("email")
    .Build();
```
