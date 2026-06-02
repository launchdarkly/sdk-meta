---
id: dotnet-client-sdk/sdk-docs/using-the-relay-proxy-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "C# in section \"Using the Relay Proxy\""
# TODO(snippet-bug): body mixes APIs across SDK eras —
# `Configuration.Builder(..., AutoEnvAttributes.Enabled)` is a
# 3.1+ signature, but `LdClient.Init(config)` (1-arg sync) is only
# available in 3.0.x. No single LaunchDarkly.ClientSdk version
# accepts both shapes. Fix in the snippet-bugs PR: align the body
# on one version (either drop AutoEnvAttributes and pin 3.0.x, or
# upgrade Init to InitAsync(config, context) and float to latest).
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
