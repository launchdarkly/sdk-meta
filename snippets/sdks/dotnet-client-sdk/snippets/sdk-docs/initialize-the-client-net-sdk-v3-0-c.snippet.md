---
id: dotnet-client-sdk/sdk-docs/initialize-the-client-net-sdk-v3-0-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: ".NET SDK v3.0 (C#) in section \"Initialize the client\""
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
// You'll need this context later, but you can ignore it for now.
var context = Context.New("example-context-key");
var timeSpan = TimeSpan.FromSeconds(5);
client = LdClient.Init("example-mobile-key", context, timeSpan);
```
