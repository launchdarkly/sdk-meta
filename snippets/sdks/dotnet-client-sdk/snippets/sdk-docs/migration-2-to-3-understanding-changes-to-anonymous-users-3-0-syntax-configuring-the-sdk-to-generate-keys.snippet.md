---
id: dotnet-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-anonymous-users-3-0-syntax-configuring-the-sdk-to-generate-keys
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "3.0 syntax, configuring the SDK to generate keys in section \"Understanding changes to anonymous users\""
---

```csharp
var config = Configuration.Builder("example-mobile-key")
    .GenerateAnonymousKeys(true)
    .Build();
```
