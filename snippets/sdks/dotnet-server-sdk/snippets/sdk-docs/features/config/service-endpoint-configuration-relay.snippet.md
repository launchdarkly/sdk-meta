---
id: dotnet-server-sdk/sdk-docs/features/config/service-endpoint-configuration-relay
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Service endpoint configuration example for .NET (server-side).
---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .ServiceEndpoints(Components.ServiceEndpoints()
      .RelayProxy("https://your-relay-proxy.com:8030"))
    .Build();
```
