---
id: dotnet-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-configuration-options-net-sdk-v4-0-c-using-init
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: ".NET SDK v4.0 (C#), using Init in section \"Understanding changes to configuration options\""
---

```csharp
var context = Context.New("example-context-key");
var timeSpan = TimeSpan.FromSeconds(10);

client = LdClient.Init("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled, context, timeSpan);
```
