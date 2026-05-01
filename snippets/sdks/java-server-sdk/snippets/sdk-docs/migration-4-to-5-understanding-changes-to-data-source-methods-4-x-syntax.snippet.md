---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-data-source-methods-4-x-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "4.x syntax in section \"Understanding changes to data source methods\""
---

```java
// 4.x model: setting custom options for streaming mode
LDConfig config = new LDConfig.Builder()
  .stream(true)
  .reconnectTimeMs(2000)
  .build();

// 4.x model: specifying polling mode and setting custom polling options
LDConfig config = new LDConfig.Builder()
  .stream(false)
  .pollingIntervalMillis(60000)
  .build();
```
