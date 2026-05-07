---
id: dotnet-client-sdk/sdk-docs/initialize-the-client-net-sdk-v3-x-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: ".NET SDK v3.x (C#) in section \"Initialize the client\""
# Bucket C: pinned to a deprecated dotnet-client SDK API surface
# (LdClient.Init(string, …) overload removed in v4.0; v3.x async
# variant; v3.x relay-proxy via ConfigurationBuilder shape that
# changed in v4). The csharp-client-syntax-only scaffold compiles
# against the latest LaunchDarkly.ClientSdk, so these v3-shape calls
# fail overload resolution. See _sdk-docs-port-notes.md.
---

```csharp
// You'll need this context later, but you can ignore it for now.
Context context = Context.New("example-context-key");
client = await LdClient.InitAsync("example-mobile-key", context);
```
