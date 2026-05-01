---
id: dotnet-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-data-source-methods-2-0-syntax
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "2.0 syntax in section \"Understanding changes to data source methods\""
---

```csharp
// 2.0 model: setting custom options for streaming mode
var config = Configuration.Builder("example-mobile-key")
    .DataSource(
        Components.StreamingDataSource()
            .InitialReconnectDelay(TimeSpan.FromSeconds(2))
    )
    .Build();

// 2.0 model: specifying polling mode and setting custom polling options
var config = Configuration.Builder("example-mobile-key")
    .DataSource(
        Components.PollingDataSource()
            .PollInterval(TimeSpan.FromSeconds(60))
    )
    .Build();
```
