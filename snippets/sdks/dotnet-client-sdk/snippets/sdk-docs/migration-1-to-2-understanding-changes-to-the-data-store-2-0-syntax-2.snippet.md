---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-the-data-store-2-0-syntax-2
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 syntax in section \"Understanding changes to the data store\""
---

```csharp
var config = Configuration.Builder("example-mobile-key")
    .Persistence(Components.NoPersistence)
    .Build();
```
