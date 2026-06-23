---
id: dotnet-client-sdk/sdk-docs/features/anonymous/anonymous-generate-keys-v3
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Anonymous key generation configuration for .NET (client-side), SDK v3.0.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-v3
---

```csharp
var config = Configuration.Builder("example-mobile-key")
    .GenerateAnonymousKeys(true)
    .Build();
```
