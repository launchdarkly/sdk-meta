---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-custom-uri-configuration-1-x-syntax-2
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "1.x syntax in section \"Understanding changes to custom URI configuration\""
---

```csharp
var config = Configuration.Builder("example-mobile-key")
    .BaseUri(new Uri("https://app.mycompany.launchdarkly.com"))
    .StreamUri(new Uri("https://stream.mycompany.launchdarkly.com"))
    .EventsUri(new Uri("https://events.mycompany.launchdarkly.com"))
    .Build();
```
