---
id: dotnet-server-sdk/sdk-docs/features/config/app-config
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Application metadata configuration example for .NET (server-side).
---

```csharp

var config = Configuration.Builder("YOUR_SDK_KEY")
    .ApplicationInfo(Components.ApplicationInfo()
        .ApplicationId("authentication-service")
        .ApplicationName("Authentication-Service")
        .ApplicationVersion("1.0.0")
        .ApplicationVersionName("v1")
    )
    .Build();

var client = new LdClient(config);

```
