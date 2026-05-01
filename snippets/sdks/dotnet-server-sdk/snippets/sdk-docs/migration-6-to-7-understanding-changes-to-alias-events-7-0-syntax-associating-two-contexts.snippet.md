---
id: dotnet-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-alias-events-7-0-syntax-associating-two-contexts
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "7.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```csharp
var userContext = Context.New("example-context-key");

var deviceContext = Context.Builder("example-device-key")
    .Kind("device")
    .Build();

var multiContext = Context.NewMulti(userContext, deviceContext);

client.Identify(multiContext);
```
