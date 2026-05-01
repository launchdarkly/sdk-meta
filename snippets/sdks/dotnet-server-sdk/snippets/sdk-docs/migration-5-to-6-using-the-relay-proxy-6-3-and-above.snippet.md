---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-using-the-relay-proxy-6-3-and-above
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "6.3 and above in section \"Using the Relay Proxy\""
---

```csharp
// proxy mode - this example requires version 6.3.0 or higher
var relayUri = new Uri("http://my-relay-host:8000");
var config = Configuration.Builder("YOUR_SDK_KEY")
    .ServiceEndpoints(Components.ServiceEndpoints().RelayProxy(relayUri))
    .Build();

// daemon mode with a Redis database
var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSource(Components.ExternalUpdatesOnly) // replaces "UseLdd"
    .DataStore(
        Components.PersistentDataStore(
            Redis.DataStore().Uri(new Uri("redis://my-redis-host"))
        )
    )
    .Build();
```
