---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-custom-uri-configuration-2-0-syntax
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 syntax in section \"Understanding changes to custom URI configuration\""
---

```csharp
var config = Configuration.Builder("example-mobile-key")
    .ServiceEndpoints(
        Components.ServiceEndpoints().RelayProxy("http://my-relay-proxy:8080")
    )
    .Build();
```
