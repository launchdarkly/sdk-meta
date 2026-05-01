---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-json-net-2-0-only
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 only in section \"Json.NET\""
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
