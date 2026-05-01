---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-events-5-x-syntax
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "5.x syntax in section \"Understanding changes to events\""
---

```csharp
// 5.x model: disabling events
var config = Configuration.Builder("YOUR_SDK_KEY")
    .SendEvents(false)
    .Build();

// 5.x model: customizing event behavior
var config =  Configuration.Builder("YOUR_SDK_KEY")
    .EventCapacity(20000)
    .EventFlushInterval(TimeSpan.FromSeconds(10))
    .PrivateAttribute("email")
    .PrivateAttribute("myCustomAttribute")
    .Build();
```
