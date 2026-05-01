---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-custom-uri-configuration-1-x-syntax
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "1.x syntax in section \"Understanding changes to custom URI configuration\""
---

```csharp
var config = Configuration.Builder("example-mobile-key")
    .BaseUri(new Uri("http://my-relay-proxy:8080"))
    .StreamUri(new Uri("http://my-relay-proxy:8080"))
    .EventsUri(new Uri("http://my-relay-proxy:8080"))
    .Build();
```
