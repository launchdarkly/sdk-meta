---
id: dotnet-server-sdk/ai-configs/context
sdk: dotnet-server-sdk
kind: context
lang: csharp
file: dotnet-server-sdk/ai-configs/context.txt
description: Build an evaluation context for dotnet-server-sdk AI Configs.
---

```csharp
var context = Context.Builder("context-key-123abc")
  .Name("Sandy")
  .LastName("Smith")
  .Email("sandy@example.com")
  .Groups(["Google", "Microsoft"])
  .Build();
```
