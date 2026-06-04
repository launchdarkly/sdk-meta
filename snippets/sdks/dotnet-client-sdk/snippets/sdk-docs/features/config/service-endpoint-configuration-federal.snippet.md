---
id: dotnet-client-sdk/sdk-docs/features/config/service-endpoint-configuration-federal
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Service endpoint configuration example for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only

---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .ServiceEndpoints(Components.ServiceEndpoints()
      .Streaming("https://clientstream.launchdarkly.us")
      .Polling("https://clientsdk.launchdarkly.us")
      .Events("https://events.launchdarkly.us"))
    .Build();
```
