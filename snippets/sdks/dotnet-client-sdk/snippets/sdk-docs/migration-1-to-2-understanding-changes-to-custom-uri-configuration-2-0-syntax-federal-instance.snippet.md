---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-custom-uri-configuration-2-0-syntax-federal-instance
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 syntax, federal instance in section \"Understanding changes to custom URI configuration\""
---

```csharp
var config = Configuration.Builder("example-mobile-key")
    .ServiceEndpoints(
        Components.ServiceEndpoints()
            .Polling("https://clientsdk.mycompany.launchdarkly.us")
            .Streaming("https://stream.mycompany.launchdarkly.us")
            .Events("https://events.mycompany.launchdarkly.us")
    )
    .Build();
```
