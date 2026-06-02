---
id: dotnet-client-sdk/sdk-docs/initialize-the-client-net-sdk-v3-0-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: ".NET SDK v3.0 (C#) in section \"Initialize the client\""
# TODO(snippet-bug): body uses the v3.0 `LdClient.Init(string, Context, TimeSpan)`
# overload, removed in v4.0 (replaced with a Configuration object).
# The csharp-client-syntax-only scaffold compiles against the latest
# LaunchDarkly.ClientSdk, so this v3-shape call fails. Fix in the
# follow-up snippet-bugs PR: either (a) update the snippet to current
# v4 API and drop the v3-pinned variant, or (b) pin a v3 SDK in a
# parallel scaffold/validator if back-compat docs must stay live.
---

```csharp
// You'll need this context later, but you can ignore it for now.
var context = Context.New("example-context-key");
var timeSpan = TimeSpan.FromSeconds(5);
client = LdClient.Init("example-mobile-key", context, timeSpan);
```
