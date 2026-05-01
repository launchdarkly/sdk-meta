---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-json-net-6-0-only-2
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "6.0 only in section \"Json.NET\""
---

```csharp
    struct ValuesBeingPassedToJavaScriptCode
    {
        [JsonProperty("user")]
        public JRaw User;
    }

    ValuesBeingPassedToJavaScriptCode CreateValuesForJavaScript(User user)
    {
        var userJson = new JRaw(LdJsonSerialization.SerializeObject(user));
        return new ValuesBeingPassedToJavaScriptCode { User = userJson };
    }
```
