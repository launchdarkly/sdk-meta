---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-the-data-store-6-0-syntax
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "6.0 syntax in section \"Understanding changes to the data store\""
---

```csharp
// 6.0 model: use Redis, set custom Redis URI and key prefix, set cache TTL to 45 seconds
var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataStore(
        Components.PersistentDataStore(
            Redis.DataStore()
                .Uri(new Uri("redis://my-redis-host"))
                .Prefix("my-prefix")
        ).CacheSeconds(45)
    )
    .Build();
```
