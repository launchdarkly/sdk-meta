---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-logging-6-0-only-2
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "6.0 only in section \"Understanding changes to logging\""
---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .Logging(LdLog4net.Adapter)
    .Build();
```
