---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-events-6-0-syntax
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "6.0 syntax in section \"Understanding changes to events\""
---

```csharp
// 6.0 model: disabling events
var config =  Configuration.Builder("YOUR_SDK_KEY")
    .Events(Components.NoEvents)
    .Build();

// 6.0 model: customizing event behavior
var config =  Configuration.Builder("YOUR_SDK_KEY")
    .Events(
        Components.SendEvents()
            .Capacity(20000)
            .FlushInterval(TimeSpan.FromSeconds(10))
            .PrivateAttributes(UserAttribute.Email,
                UserAttribute.ForName("myCustomAttribute"))
    )
    .Build();
```
