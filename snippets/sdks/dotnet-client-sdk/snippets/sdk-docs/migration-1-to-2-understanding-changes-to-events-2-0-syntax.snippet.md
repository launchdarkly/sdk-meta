---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-events-2-0-syntax
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 syntax in section \"Understanding changes to events\""
---

```csharp
// 2.0 model: disabling events
var config = Configuration.Builder("example-mobile-key")
    .Events(Components.NoEvents)
    .Build();

// 2.0 model: customizing event behavior
var config = Configuration.Builder("example-mobile-key")
    .Events(
        Components.SendEvents()
            .Capacity(20000)
            .FlushInterval(TimeSpan.FromSeconds(10))
            .PrivateAttributes(UserAttribute.Email,
                UserAttribute.ForName("myCustomAttribute"))
    )
    .Build();
```
