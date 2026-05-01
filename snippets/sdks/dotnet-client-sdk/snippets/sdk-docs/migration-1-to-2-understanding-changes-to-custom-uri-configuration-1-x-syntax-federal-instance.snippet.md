---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-custom-uri-configuration-1-x-syntax-federal-instance
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "1.x syntax, federal instance in section \"Understanding changes to custom URI configuration\""
---

```csharp
var config = Configuration.Builder("example-mobile-key")
    .BaseUri(new Uri("https://clientsdk.mycompany.launchdarkly.us"))
    .StreamUri(new Uri("https://stream.mycompany.launchdarkly.us"))
    .EventsUri(new Uri("https://events.mycompany.launchdarkly.us"))
    .Build();
```
