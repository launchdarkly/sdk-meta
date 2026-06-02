---
id: dotnet-client-sdk/sdk-docs/using-the-relay-proxy-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "C# in section \"Using the Relay Proxy\""
# Bucket C: pinned to a deprecated dotnet-client SDK API surface
# (LdClient.Init(string, …) overload removed in v4.0; v3.x async
# variant; v3.x relay-proxy via ConfigurationBuilder shape that
# changed in v4). The csharp-client-syntax-only scaffold compiles
# against the latest LaunchDarkly.ClientSdk, so these v3-shape calls
# fail overload resolution. See _sdk-docs-port-notes.md.
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
