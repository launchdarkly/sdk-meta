---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-networking-5-x-syntax
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "5.x syntax in section \"Understanding changes to networking\""
---

```csharp
// 5.x model: setting connection and read timeouts
var config =  Configuration.Builder("YOUR_SDK_KEY")
    .HttpClientTimeout(TimeSpan.FromSeconds(3))
    .ReadTimeout(TimeSpan.FromSeconds(4))
    .Build();
```
