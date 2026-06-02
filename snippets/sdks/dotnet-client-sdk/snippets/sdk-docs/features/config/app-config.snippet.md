---
id: dotnet-client-sdk/sdk-docs/features/config/app-config
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Application metadata configuration example for .NET (client-side).
---

```csharp
var config = Configuration
    .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
    .ApplicationInfo(Components.ApplicationInfo()
        .ApplicationId("authentication-service")
        .ApplicationName("Authentication-Service")
        .ApplicationVersion("1.0.0")
        .ApplicationVersionName("v1")
    )
    .Build();
```
