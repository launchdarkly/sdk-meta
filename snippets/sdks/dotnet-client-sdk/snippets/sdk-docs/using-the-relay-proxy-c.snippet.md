---
id: dotnet-client-sdk/sdk-docs/using-the-relay-proxy-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "C# in section \"Using the Relay Proxy\""
---

```csharp
var config = Configuration
    .Builder("example-mobile-key", AutoEnvAttributes.Enabled)
    .ServiceEndpoints(
        Components.ServiceEndpoints().RelayProxy("YOUR_RELAY_URI")
    )
    .Build();
LdClient client = LdClient.Init(config);
```
