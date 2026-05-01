---
id: dotnet-client-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-automatic-environment-attributes-net-client-side-sdk-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: ".NET (client-side) SDK (C#) in section \"Understanding changes to automatic environment attributes\""
---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .Build();
```
