---
id: dotnet-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-anonymous-users-3-0-syntax-building-an-anonymous-context
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "3.0 syntax, building an anonymous context in section \"Understanding changes to anonymous users\""
---

```csharp
Context.Builder("unknown-context-key")
    .Anonymous(true)
    .Build();
```
