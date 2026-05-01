---
id: dotnet-client-sdk/sdk-docs/initialize-the-client-net-sdk-v5-2-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: ".NET SDK v5.2+ (C#) in section \"Initialize the client\""
---

```csharp
// You'll need this context later, but you can ignore it for now.
Context context = Context.New("example-context-key");
var timeSpan = TimeSpan.FromSeconds(5);
client = await LdClient.InitAsync("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled, context, timeSpan);
```
