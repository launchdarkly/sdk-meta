---
id: dotnet-client-sdk/sdk-docs/initialize-the-client-net-sdk-v3-x-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: ".NET SDK v3.x (C#) in section \"Initialize the client\""
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-v3
---

```csharp
// You'll need this context later, but you can ignore it for now.
Context context = Context.New("example-context-key");
client = await LdClient.InitAsync("example-mobile-key", context);
```
