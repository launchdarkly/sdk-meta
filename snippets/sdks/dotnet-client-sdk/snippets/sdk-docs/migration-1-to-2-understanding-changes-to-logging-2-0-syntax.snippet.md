---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-logging-2-0-syntax
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 syntax in section \"Understanding changes to logging\""
---

```csharp
using LaunchDarkly.Logging;
using LaunchDarkly.Sdk.Client;

var config = Configuration.Builder("example-mobile-key")
    .Logging(Logs.ToConsole.Level(LogLevel.Warn))
    .Build();
```
