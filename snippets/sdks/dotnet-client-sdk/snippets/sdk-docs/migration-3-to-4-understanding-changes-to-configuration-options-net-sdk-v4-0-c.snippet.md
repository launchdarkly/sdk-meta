---
id: dotnet-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-configuration-options-net-sdk-v4-0-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: ".NET SDK v4.0 (C#) in section \"Understanding changes to configuration options\""
---

```csharp
var config = Configuration
  .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
  .Build();
```
