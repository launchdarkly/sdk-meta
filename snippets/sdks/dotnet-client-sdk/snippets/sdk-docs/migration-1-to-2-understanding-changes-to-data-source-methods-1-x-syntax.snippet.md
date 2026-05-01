---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-data-source-methods-1-x-syntax
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "1.x syntax in section \"Understanding changes to data source methods\""
---

```csharp
// 1.x model: setting custom options for streaming mode
var config = Configuration.Builder("example-mobile-key")
    .IsStreamingEnabled(true)
    .ReconnectTime(TimeSpan.FromSeconds(2))
    .Build();

// 1.x model: specifying polling mode and setting custom polling options
var config = Configuration.Builder("example-mobile-key")
    .IsStreamingEnabled(false)
    .PollingInterval(TimeSpan.FromSeconds(60))
    .Build();
```
