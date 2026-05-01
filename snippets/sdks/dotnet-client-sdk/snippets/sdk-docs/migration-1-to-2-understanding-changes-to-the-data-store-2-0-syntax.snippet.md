---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-the-data-store-2-0-syntax
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 syntax in section \"Understanding changes to the data store\""
---

```csharp
// Keep flags for the most recent 2 users in persistent storage
var config = Configuration.Builder("example-mobile-key")
    .Persistence(
        Components.Persistence().MaxCachedUsers(2)
    )
    .Build();
```
