---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-data-source-methods-5-x-syntax
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "5.x syntax in section \"Understanding changes to data source methods\""
---

```csharp
// 5.x model: setting custom options for streaming mode
var config = Configuration.Builder("YOUR_SDK_KEY")
    .IsStreamingEnabled(true)
    .ReconnectTime(TimeSpan.FromSeconds(2))
    .Build();

// 5.x model: specifying polling mode and setting custom polling options
var config = Configuration.Builder("YOUR_SDK_KEY")
    .IsStreamingEnabled(false)
    .PollingInterval(TimeSpan.FromSeconds(60))
    .Build();
```
