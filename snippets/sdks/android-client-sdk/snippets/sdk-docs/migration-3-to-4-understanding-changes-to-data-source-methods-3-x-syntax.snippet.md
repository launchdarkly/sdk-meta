---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-data-source-methods-3-x-syntax
sdk: android-client-sdk
kind: reference
lang: java
description: "3.x syntax in section \"Understanding changes to data source methods\""
---

```java
// Setting custom options for streaming mode
LDConfig config = new LDConfig.Builder()
  .stream(true)
  .backgroundPollingIntervalMillis(120000)
  .build();

// Specifying polling mode and setting custom polling options
LDConfig config = new LDConfig.Builder()
  .stream(false)
  .pollingIntervalMillis(60000)
  .backgroundPollingIntervalMillis(120000)
  .build();
```
