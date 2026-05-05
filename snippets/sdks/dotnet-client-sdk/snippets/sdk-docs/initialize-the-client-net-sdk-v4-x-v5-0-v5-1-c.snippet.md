---
id: dotnet-client-sdk/sdk-docs/initialize-the-client-net-sdk-v4-x-v5-0-v5-1-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: ".NET SDK v4.x, v5.0, v5.1 (C#) in section \"Initialize the client\""
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
// You'll need this context later, but you can ignore it for now.
Context context = Context.New("example-context-key");
client = await LdClient.InitAsync("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled, context);
```
