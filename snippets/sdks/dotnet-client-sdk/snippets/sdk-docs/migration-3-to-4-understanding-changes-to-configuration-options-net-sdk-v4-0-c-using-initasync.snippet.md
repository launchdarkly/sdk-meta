---
id: dotnet-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-configuration-options-net-sdk-v4-0-c-using-initasync
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: ".NET SDK v4.0 (C#), using InitAsync in section \"Understanding changes to configuration options\""
---

```csharp
var context = Context.New("example-context-key");
client = await LdClient.InitAsync("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled, context);
```
