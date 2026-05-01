---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-networking-6-0-syntax
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "6.0 syntax in section \"Understanding changes to networking\""
---

```csharp
// 6.0 model: setting connection and read timeouts
var config =  Configuration.Builder("YOUR_SDK_KEY")
    .Http(
        Components.HttpConfiguration()
            .ConnectTimeout(TimeSpan.FromSeconds(3))
            .SocketTimeout(TimeSpan.FromSeconds(4))
    )
    .Build();

// 6.0 model: specifying an HTTP proxy
var config =  Configuration.Builder("YOUR_SDK_KEY")
    .Http(
        Components.HttpConfiguration()
            .Proxy(new System.Net.WebProxy(new Uri("http://my-proxy:8080")))
    )
    .Build();
```
