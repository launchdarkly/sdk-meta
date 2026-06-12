---
id: dotnet-server-sdk/sdk-docs/features/webproxy/web-proxy-v5
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Web proxy configuration for .NET (server-side) via a custom message handler.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var handler = new System.Net.Http.HttpClientHandler();
handler.Proxy = new System.Net.WebProxy("http://my-proxy-host:8080");

var config = Configuration.Builder("YOUR_SDK_KEY")
    .Http(Components.HttpConfiguration().MessageHandler(handler))
    .Build();
```
