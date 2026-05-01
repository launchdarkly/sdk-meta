---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-data-source-methods-4-0-syntax
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax in section \"Understanding changes to data source methods\""
---

```java
// Setting custom options for streaming mode
LDConfig config = new LDConfig.Builder()
  .dataSource(
    Components.streamingDataSource()
      .backgroundPollIntervalMillis(120000)
  )
  .build();

// Specifying polling mode and setting custom polling options
LDConfig config = new LDConfig.Builder()
  .dataSource(
    Components.pollingDataSource()
      .pollIntervalMillis(60000)
      .backgroundPollIntervalMillis(120000)
  )
  .build();
```
