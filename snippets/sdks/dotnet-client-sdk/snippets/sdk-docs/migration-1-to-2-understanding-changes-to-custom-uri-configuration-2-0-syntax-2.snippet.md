---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-custom-uri-configuration-2-0-syntax-2
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 syntax in section \"Understanding changes to custom URI configuration\""
---

```csharp
var config = Configuration.Builder("example-mobile-key")
    .ServiceEndpoints(
        Components.ServiceEndpoints()
            .Polling("https://app.mycompany.launchdarkly.com")
            .Streaming("https://stream.mycompany.launchdarkly.com")
            .Events("https://events.mycompany.launchdarkly.com")
    )
    .Build();
```
