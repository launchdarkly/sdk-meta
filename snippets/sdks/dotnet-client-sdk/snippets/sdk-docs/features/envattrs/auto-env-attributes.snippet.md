---
id: dotnet-client-sdk/sdk-docs/features/envattrs/auto-env-attributes
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Automatic environment attributes configuration for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only

---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .Build();
```
