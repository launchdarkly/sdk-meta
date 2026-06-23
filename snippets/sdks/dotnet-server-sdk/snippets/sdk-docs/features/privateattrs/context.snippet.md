---
id: dotnet-server-sdk/sdk-docs/features/privateattrs/context
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Marking context attributes private with the context builder for .NET server SDK v7.0.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
var context = Context.Builder("example-context-key")
    .Set("email", "sandy@example.com")
    .Private("email")
    .Build();
```
