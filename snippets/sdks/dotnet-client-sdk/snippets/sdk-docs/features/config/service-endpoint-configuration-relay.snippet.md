---
id: dotnet-client-sdk/sdk-docs/features/config/service-endpoint-configuration-relay
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Service endpoint configuration example for .NET (client-side).
---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .ServiceEndpoints(Components.ServiceEndpoints()
      .RelayProxy("https://your-relay-proxy.com:8030"))
    .Build();
```
