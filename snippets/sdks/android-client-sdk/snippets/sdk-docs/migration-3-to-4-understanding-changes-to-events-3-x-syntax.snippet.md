---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-events-3-x-syntax
sdk: android-client-sdk
kind: reference
lang: java
description: "3.x syntax in section \"Understanding changes to events\""
---

```java
// Customizing event behavior
LDConfig config = new LDConfig.Builder()
  .capacity(20000)
  .flushIntervalMillis(10000)
  .privateAttributes("email", "name", "myCustomAttribute")
  .build();

// Disabling events is not possible in 3.x
```
