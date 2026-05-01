---
id: dotnet-server-sdk/sdk-docs/transport-layer-security-tls-and-other-networking-issues-net-sdk-6-x-and-7-x-syntax-c
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: ".NET SDK 6.x and 7.x syntax (C#) in section \"Transport Layer Security (TLS) and other networking issues\""
---

```csharp
// Use `stream-tls10.launchdarkly.com` and `events-tls10.launchdarkly.com` TLSv1 endpoints
var config = Configuration.Builder("YOUR_SDK_KEY")
    .ServiceEndpoints(
        Components.ServiceEndpoints()
          .Streaming("https://stream-tls10.launchdarkly.com")
          .Events("https://events-tls10.launchdarkly.com"))
    .Build();
```
