---
id: dotnet-client-sdk/sdk-docs/using-the-relay-proxy-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "C# in section \"Using the Relay Proxy\""
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .ServiceEndpoints(
        Components.ServiceEndpoints().RelayProxy("YOUR_RELAY_URI")
    )
    .Build();
var context = Context.Builder("example-context-key").Build();
LdClient client = LdClient.Init(config, context, TimeSpan.FromSeconds(10));
```
