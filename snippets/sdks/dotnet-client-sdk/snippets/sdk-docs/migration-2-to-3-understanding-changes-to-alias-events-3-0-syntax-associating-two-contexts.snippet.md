---
id: dotnet-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-alias-events-3-0-syntax-associating-two-contexts
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "3.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```csharp
var userContext = Context.New("example-user-key");

var deviceContext = Context.New(ContextKind.Of("device"), "example-device-key");

var multiContext = Context.NewMulti(userContext, deviceContext);

await client.IdentifyAsync(multiContext); // or, use synchronous Identify
```
