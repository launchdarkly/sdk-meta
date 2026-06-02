---
id: dotnet-server-sdk/sdk-docs/features/config/service-endpoint-configuration-eu
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Service endpoint configuration example for .NET (server-side).
---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .ServiceEndpoints(Components.ServiceEndpoints()
      .Streaming("https://stream.eu.launchdarkly.com")
      .Polling("https://sdk.eu.launchdarkly.com")
      .Events("https://events.eu.launchdarkly.com"))
    .Build();
```
