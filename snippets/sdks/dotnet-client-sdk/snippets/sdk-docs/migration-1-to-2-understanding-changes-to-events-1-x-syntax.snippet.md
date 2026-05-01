---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-events-1-x-syntax
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "1.x syntax in section \"Understanding changes to events\""
---

```csharp
// 1.x model: disabling events
var config = Configuration.Builder("example-mobile-key")
    .SendEvents(false)
    .Build();

// 1.x model: customizing event behavior
var config = Configuration.Builder("example-mobile-key")
    .EventCapacity(20000)
    .EventFlushInterval(TimeSpan.FromSeconds(10))
    .PrivateAttribute("email")
    .PrivateAttribute("myCustomAttribute")
    .Build();
```
