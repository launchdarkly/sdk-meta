---
id: dotnet-server-sdk/sdk-docs/features/logging/logging
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: .NET-server logging destination and level configuration example (C#).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
using LaunchDarkly.Logging;
using LaunchDarkly.Sdk.Server;

var config = Configuration.Builder("YOUR_SDK_KEY")
    .Logging(
        Components.Logging(Logs.ToWriter(Console.Out)).Level(LogLevel.Debug)
    )
    .Build();
```
