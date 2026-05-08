---
id: dotnet-server-sdk/ai-configs/context
sdk: dotnet-server-sdk
kind: context
lang: csharp
file: dotnet-server-sdk/ai-configs/context.txt
description: Build an evaluation context for dotnet-server-sdk AI Configs.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var context = Context.Builder("context-key-123abc")
  .Name("Sandy")
  .Set("lastName", "Smith")
  .Set("email", "sandy@example.com")
  .Set("groups", LdValue.ArrayOf(LdValue.Of("Google"), LdValue.Of("Microsoft")))
  .Build();
```
