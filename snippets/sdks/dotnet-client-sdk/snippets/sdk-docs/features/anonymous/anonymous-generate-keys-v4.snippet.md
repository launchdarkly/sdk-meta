---
id: dotnet-client-sdk/sdk-docs/features/anonymous/anonymous-generate-keys-v4
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Anonymous key generation configuration for .NET (client-side), SDK v4.0.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .GenerateAnonymousKeys(true)
    .Build();
```
