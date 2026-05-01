---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-networking-2-0-syntax
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 syntax in section \"Understanding changes to networking\""
---

```csharp
// 2.0 model: setting connection and read timeouts
var config = Configuration.Builder("example-mobile-key")
    .Http(
        Components.HttpConfiguration()
            .ConnectTimeout(TimeSpan.FromSeconds(3))
            .SocketTimeout(TimeSpan.FromSeconds(4))
    )
    .Build();

// 2.0 model: specifying an HTTP proxy
var config = Configuration.Builder("example-mobile-key")
    .Http(
        Components.HttpConfiguration()
            .Proxy(new System.Net.WebProxy(new Uri("http://my-proxy:8080"))
    )
    .Build();
```
