---
id: dotnet-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-data-source-methods-6-0-syntax
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "6.0 syntax in section \"Understanding changes to data source methods\""
---

```csharp
// 6.0 model: setting custom options for streaming mode
var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSource(
        Components.StreamingDataSource()
            .InitialReconnectDelay(TimeSpan.FromSeconds(2))
    )
    .Build();

// 6.0 model: specifying polling mode and setting custom polling options
var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSource(
        Components.PollingDataSource()
            .PollInterval(TimeSpan.FromSeconds(60))
    )
    .Build();
```
