---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-networking-1-x-syntax
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "1.x syntax in section \"Understanding changes to networking\""
---

```csharp
// 1.x model: setting connection and read timeouts
var config = Configuration.Builder("example-mobile-key")
    .HttpClientTimeout(TimeSpan.FromSeconds(3))
    .ReadTimeout(TimeSpan.FromSeconds(4))
    .Build();
```
