---
id: dotnet-server-sdk/sdk-docs/migration-7-to-8-understanding-application-metadata-net-sdk-v8-c
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: ".NET SDK v8 (C#) in section \"Understanding application metadata\""
---

```csharp

var config = Configuration.Builder("YOUR_SDK_KEY")
    .ApplicationInfo(Components.ApplicationInfo()
        .ApplicationID("authentication-service")
        .ApplicationName("Authentication-Service")
        .ApplicationVersion("1.0.0")
        .ApplicationVersionName("v1")
    )
    .Build();

var client = new LdClient(config);

```
