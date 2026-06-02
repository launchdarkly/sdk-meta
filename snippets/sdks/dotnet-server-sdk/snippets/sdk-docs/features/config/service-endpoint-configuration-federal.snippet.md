---
id: dotnet-server-sdk/sdk-docs/features/config/service-endpoint-configuration-federal
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Service endpoint configuration example for .NET (server-side).
---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .ServiceEndpoints(Components.ServiceEndpoints()
      .Streaming("https://stream.launchdarkly.us")
      .Polling("https://sdk.launchdarkly.us")
      .Events("https://events.launchdarkly.us"))
    .Build();
```
