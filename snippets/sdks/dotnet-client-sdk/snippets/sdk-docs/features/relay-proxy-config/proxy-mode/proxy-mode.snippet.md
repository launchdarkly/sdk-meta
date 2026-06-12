---
id: dotnet-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Proxy mode configuration example for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .ServiceEndpoints(Components.ServiceEndpoints()
      .RelayProxy("https://your-relay-proxy.com:8030"))
    .Build();
```
