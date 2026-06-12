---
id: dotnet-server-sdk/sdk-docs/features/datasaving/relay-proxy-fallback
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Data saving mode with Relay Proxy and LaunchDarkly API fallback for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Integrations;

var relayUri = new Uri("http://my-relay-proxy:8030");
var relayEndpoints = Components.ServiceEndpoints().RelayProxy(relayUri);

var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSystem(
        Components.DataSystem().Custom()
            .Initializers(
                DataSystemComponents.Polling()
                    .ServiceEndpointsOverride(relayEndpoints),
                DataSystemComponents.Polling()
            )
            .Synchronizers(
                DataSystemComponents.Streaming()
                    .ServiceEndpointsOverride(relayEndpoints),
                DataSystemComponents.Streaming(),
                DataSystemComponents.Polling()
            )
    )
    .Build();

var client = new LdClient(config);
```
