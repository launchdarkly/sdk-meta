---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-json-net-6-0-only
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "6.0 only in section \"Json.NET\""
---

```csharp
    var settings = new JsonSerializerSettings
    {
        Converters = new List<JsonConverter>
        {
            LaunchDarkly.Sdk.Json.LdJsonNet.Converter
        }
    };
```
