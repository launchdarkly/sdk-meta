---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-the-data-store-5-x-syntax
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "5.x syntax in section \"Understanding changes to the data store\""
---

```csharp
// 5.x model: use Redis, set custom Redis URI and key prefix, set cache TTL to 45 seconds
var config = Configuration.Builder("YOUR_SDK_KEY")
    .FeatureStoreFactory(
        RedisComponents.RedisFeatureStore()
            .WithRedisUri(new Uri("redis://my-redis-host"))
            .WithPrefix("my-prefix")
            .WithCacheExpiration(TimeSpan.FromSeconds(45))
    )
    .Build();
```
