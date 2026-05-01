---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-logging-6-0-syntax
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "6.0 syntax in section \"Understanding changes to logging\""
---

```csharp
using LaunchDarkly.Logging;
using LaunchDarkly.Sdk.Server;

var config = Configuration.Builder("YOUR_SDK_KEY")
    .Logging(Logs.ToConsole.Level(LogLevel.Warn))
    .Build();
```
