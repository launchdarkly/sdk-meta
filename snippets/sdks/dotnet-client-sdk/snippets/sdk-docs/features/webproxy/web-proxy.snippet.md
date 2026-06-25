---
id: dotnet-client-sdk/sdk-docs/features/webproxy/web-proxy
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Web proxy configuration for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
var handler = new System.Net.Http.HttpClientHandler();
handler.Proxy = new System.Net.WebProxy("http://my-proxy-host:8080");

var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .Http(Components.HttpConfiguration().MessageHandler(handler))
    .Build();
```
