---
id: dotnet-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-private-attributes-3-0-syntax-two-attributes-marked-private
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "3.0 syntax, two attributes marked private in section \"Understanding changes to private attributes\""
---

```csharp
var config = Configuration.Builder("example-mobile-key")
    .PrivateAttributes("email", "address")
    .Build();
LdClient client = LdClient.Init(config, context);
```
