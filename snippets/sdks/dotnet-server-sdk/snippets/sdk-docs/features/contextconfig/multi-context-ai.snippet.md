---
id: dotnet-server-sdk/sdk-docs/features/contextconfig/multi-context-ai
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Multi-context example for .NET AI SDK.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
var userContext = Context.New("example-context-key");

var deviceContext = Context.Builder("example-device-key")
    .Kind("device")
    .Build();

var multiContext = Context.NewMulti(userContext, deviceContext);
```
