---
id: dotnet-client-sdk/sdk-docs/features/privateattrs/context
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Marking context attributes private with the context builder for .NET client SDK v3.0+.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only

---

```csharp
var context = Context.Builder("example-context-key")
    .Set("email", "sandy@example.com")
    .Private("email")
    .Build();
```
