---
id: dotnet-server-sdk/sdk-docs/features/storing-data/consul/consul
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Consul feature store configuration example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Integrations;
var config = Configuration.Builder(sdkKey)
    .DataStore(
        Components.PersistentDataStore(
            LaunchDarkly.Sdk.Server.Integrations.Consul.DataStore().Address("http://my-consul:8100")
        ).CacheSeconds(30)
    )
    .Build();
var client = new LdClient(config);
```
