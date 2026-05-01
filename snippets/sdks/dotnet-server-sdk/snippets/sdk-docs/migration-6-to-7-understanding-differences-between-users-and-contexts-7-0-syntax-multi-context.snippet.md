---
id: dotnet-server-sdk/sdk-docs/migration-6-to-7-understanding-differences-between-users-and-contexts-7-0-syntax-multi-context
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "7.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```csharp
var userContext = Context.New("example-context-key");

var deviceContext = Context.Builder("example-device-key")
    .Kind("device")
    .Build();

var multiContext = Context.NewMulti(userContext, deviceContext);

```
