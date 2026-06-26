---
id: dotnet-server-sdk/sdk-docs/features/storing-data/redis/redis
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Redis feature store configuration example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Integrations;
var config = Configuration.Builder(sdkKey)
    .DataStore(
        Components.PersistentDataStore(
            Redis.DataStore()
                .HostAndPort("my-redis", 6379)
                .Prefix("my-key-prefix")
        ).CacheSeconds(30)
    )
    .Build();
var client = new LdClient(config);
```
