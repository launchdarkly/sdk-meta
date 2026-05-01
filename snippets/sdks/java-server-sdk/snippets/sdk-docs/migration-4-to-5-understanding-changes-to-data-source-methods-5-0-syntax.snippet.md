---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-data-source-methods-5-0-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "5.0 syntax in section \"Understanding changes to data source methods\""
---

```java
// 5.0 model: setting custom options for streaming mode
LDConfig config = new LDConfig.Builder()
  .dataSource(
    Components.streamingDataSource()
      .initialReconnectDelay(Duration.ofSeconds(2))
  )
  .build();

// 5.0 model: specifying polling mode and setting custom polling options
LDConfig config = new LDConfig.Builder()
  .dataSource(
    Components.pollingDataSource()
      .pollInterval(Duration.ofMinutes(1))
  )
  .build();
```
