---
id: dotnet-server-sdk/sdk-docs/features/logging/logging-microsoft-extensions
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: .NET-server logging adapter example for Microsoft.Extensions.Logging (C#).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
using LaunchDarkly.Logging;
using LaunchDarkly.Sdk.Server;
using Microsoft.Extensions.Logging;

// Pass an ILoggerFactory, for example from ASP.NET Core dependency injection
var config = Configuration.Builder("YOUR_SDK_KEY")
    .Logging(Logs.CoreLogging(loggerFactory))
    .Build();
```
