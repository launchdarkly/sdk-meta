---
id: dotnet-client-sdk/sdk-docs/features/webproxy/web-proxy-auth
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Web proxy configuration with authentication for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
var handler = new System.Net.Http.HttpClientHandler();
var proxy = new System.Net.WebProxy("http://my-proxy-host:8080");
var credentials = new System.Net.CredentialCache();
credentials.Add(proxy.Address, "Basic",
    new NetworkCredential("username", "password"));
proxy.Credentials = credentials;
handler.Proxy = proxy;

var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .Http(Components.HttpConfiguration().MessageHandler(handler))
    .Build();
```
