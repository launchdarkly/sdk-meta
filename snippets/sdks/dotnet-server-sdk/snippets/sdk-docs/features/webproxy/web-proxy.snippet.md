---
id: dotnet-server-sdk/sdk-docs/features/webproxy/web-proxy
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Web proxy configuration for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var proxy = new System.Net.WebProxy("http://my-proxy-host:8080");

var config = Configuration.Builder("YOUR_SDK_KEY")
    .Http(Components.HttpConfiguration().Proxy(proxy))
    .Build();
```
