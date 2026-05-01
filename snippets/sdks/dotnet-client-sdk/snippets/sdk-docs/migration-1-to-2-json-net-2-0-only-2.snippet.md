---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-json-net-2-0-only-2
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 only in section \"Json.NET\""
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
