---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-using-the-relay-proxy-5-x-syntax
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "5.x syntax in section \"Using the Relay Proxy\""
---

```csharp
// proxy mode
var relayUri = new Uri("http://my-relay-host:8000");
var config = Configuration.Builder("YOUR_SDK_KEY")
    .StreamUri(relayUri)
    .EventsUri(relayUri) // if you want to proxy events
    .Build();

// daemon mode with a Redis database
var config = Configuration.Builder("YOUR_SDK_KEY")
    .FeatureStore(
        RedisComponents.RedisFeatureStore().WithUri(new Uri("redis://my-redis-host"))
    )
    .UseLdd(true)
    .Build();
```
