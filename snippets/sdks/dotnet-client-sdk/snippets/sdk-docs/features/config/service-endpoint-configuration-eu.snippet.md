---
id: dotnet-client-sdk/sdk-docs/features/config/service-endpoint-configuration-eu
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Service endpoint configuration example for .NET (client-side).
---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .ServiceEndpoints(Components.ServiceEndpoints()
      .Streaming("https://clientstream.eu.launchdarkly.com")
      .Polling("https://clientsdk.eu.launchdarkly.com")
      .Events("https://events.eu.launchdarkly.com"))
    .Build();
```
